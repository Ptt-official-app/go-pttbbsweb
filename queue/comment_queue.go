package queue

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	theQueue chan *CommentQueue
	theQuit  []chan struct{}
)

type CommentQueue struct {
	BBoardID     bbs.BBoardID
	ArticleID    bbs.ArticleID
	OwnerID      bbs.UUserID
	CommentDBCS  []byte
	LastTime     types.NanoTS
	UpdateNanoTS types.NanoTS
}

func QueueCommentDBCS(bboardID bbs.BBoardID, articleID bbs.ArticleID, ownerID bbs.UUserID, commentDBCS []byte, firstCommentsLastTime types.NanoTS, updateNanoTS types.NanoTS) (err error) {

	commentQueue := &CommentQueue{
		BBoardID:     bboardID,
		ArticleID:    articleID,
		OwnerID:      ownerID,
		LastTime:     firstCommentsLastTime,
		CommentDBCS:  commentDBCS,
		UpdateNanoTS: updateNanoTS,
	}

	timeout := time.After(100 * time.Millisecond)

	//log.Infof("QueueCommentDBCS: to send to queue: bboardID: %v articleID: %v", bboardID, articleID)
	select {
	case theQueue <- commentQueue:
	case <-timeout:
		err = ErrTimeout
	}

	//log.Infof("QueueCommentDBCS: after send to queue: bboardID: %v articleID: %v e: %v", bboardID, articleID, err)

	return err
}

func ProcessCommentQueue(idx int, quit chan struct{}) {
	//log.Infof("ProcessCommentQueue: (%v) start", idx)
	isQuit := false
	for !isQuit {
		select {
		case commentQueue, ok := <-theQueue:
			//log.Infof("ProcessCommentQueue: (%v): received: commentQueue: (%v/%v) ok: %v", idx, commentQueue.BBoardID, commentQueue.ArticleID, ok)
			if ok {
				processCommentQueue(commentQueue)
			}
		case <-quit:
			isQuit = true
		}
	}
	//log.Infof("ProcessCommentQueue: (%v) done", idx)
}

//processCommentQueue
//
//We use LastTime as the reference time to obtain stable timestamp.
//(lastTime is from firstComments, assuming not change a lot.)
//(mtime changes frequently and may result in unstable timestamp.)
func processCommentQueue(q *CommentQueue) {
	comments, _ := dbcs.ParseComments(q.BBoardID, q.ArticleID, q.OwnerID, q.LastTime, q.CommentDBCS, q.CommentDBCS, q.UpdateNanoTS, false)
	//log.Infof("processCommentQueue: after parseComments: comments: %v", len(comments))
	if len(comments) == 0 {
		return
	}

	toAddComments, toUpdateCommentSummaries, toRemoveCommentIDs, err := diffComments(q.BBoardID, q.ArticleID, comments, q.UpdateNanoTS)
	//log.Infof("processCommentQueue: after diff: toAddComments: %v toUpdateCommentSummaries: %v toRemoveCommentIDs: %v e: %v", len(toAddComments), len(toUpdateCommentSummaries), len(toRemoveCommentIDs), err)
	if err != nil {
		return
	}

	err = schema.UpdateComments(toAddComments, q.UpdateNanoTS)
	//log.Infof("processCommentQueue: after UpdateComments: e: %v", err)
	if err != nil {
		return
	}

	err = schema.UpdateCommentSummaries(q.BBoardID, q.ArticleID, toUpdateCommentSummaries, q.UpdateNanoTS)
	if err != nil {
		return
	}

	err = schema.RemoveCommentIDs(q.BBoardID, q.ArticleID, toRemoveCommentIDs, q.UpdateNanoTS, "invalid in comment-queue")
	//log.Infof("processCommentQueue: after RemoveCommentIDs: e: %v", err)
	if err != nil {
		return
	}

	schema.UpdateArticleCommentsByArticleID(q.BBoardID, q.ArticleID, q.UpdateNanoTS)
}

func diffComments(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	comments []*schema.Comment,
	updateNanoTS types.NanoTS) (

	toAddComments []*schema.Comment,
	toUpdateCommentSummaries []*schema.CommentSummary,
	toRemoveCommentIDs []types.CommentID,
	err error) {

	startNanoTS := comments[0].CreateTime
	startNanoTSByMin := startNanoTS.ToNanoTSByMin()
	endNanoTS := comments[len(comments)-1].CreateTime
	endNanoTSByMin := endNanoTS.ToNanoTSByMin() + types.MIN_TO_NANO_TS

	commentSummaries, err := schema.GetCommentSummaries(bboardID, articleID, startNanoTSByMin, endNanoTSByMin)
	//log.Infof("diffComments: after get commment-summaries: (%v/%v) startNanoTS: %v startNanoTSByMin: %v endNanoTS: %v endNanoTSByMin: %v commentSummaries: %v e: %v", bboardID, articleID, startNanoTS, startNanoTSByMin, endNanoTS, endNanoTSByMin, len(commentSummaries), err)
	if err != nil {
		return nil, nil, nil, err
	}

	toAddComments, toUpdateCommentSummaries, toRemoveCommentIDs = diffCommentsCore(comments, commentSummaries, updateNanoTS)

	return toAddComments, toUpdateCommentSummaries, toRemoveCommentIDs, nil
}

//diffCommentsCore
//
//CommentID includes CreateTime and MD5,
//should be sufficient to determine same comment.
//
//We just need to check:
//1. if the comment is not in the comment-summaries: to-add-comments.
//2. if the comment is in the comment-summaries, but we have newer update-nano-ts: to-update-comment-summary
//3. if the comment-summary is with newer updateNanoTS: do nothing.
//4. if the comment-summary is not in the comments: to-remove-comments.
func diffCommentsCore(
	comments []*schema.Comment,
	commentSummaries_db []*schema.CommentSummary,
	updateNanoTS types.NanoTS) (

	toAddComments []*schema.Comment,
	toUpdateCommentSummaries []*schema.CommentSummary,
	toRemoveCommentIDs []types.CommentID) {

	commentMap := make(map[types.CommentID]struct{})
	for _, each := range comments {
		commentMap[each.CommentID] = struct{}{}
	}

	commentSummaryMap := make(map[types.CommentID]*schema.CommentSummary)
	for _, each := range commentSummaries_db {
		commentSummaryMap[each.CommentID] = each
	}

	toAddComments = make([]*schema.Comment, 0, len(comments))
	toUpdateCommentSummaries = make([]*schema.CommentSummary, 0, len(comments))
	for _, each := range comments {
		eachSummary, ok := commentSummaryMap[each.CommentID]
		//1. the comment is not in the comment-summaries: to-add-comments.
		if !ok {
			toAddComments = append(toAddComments, each)
			continue
		}

		//2. requires update updateNanoTS: to-update-comment-summaries
		if eachSummary.IsDeleted || eachSummary.UpdateNanoTS < each.UpdateNanoTS {
			eachSummary.UpdateNanoTS = each.UpdateNanoTS
			toUpdateCommentSummaries = append(toUpdateCommentSummaries, eachSummary)
		}
	}

	//the comment-id may be re-organized because the poster may delete the comments.
	toRemoveCommentIDs = make([]types.CommentID, 0, len(commentSummaries_db))
	for _, each := range commentSummaries_db {
		//3. the comment-summary is newer: do nothing.
		if each.IsDeleted || each.UpdateNanoTS >= updateNanoTS {
			continue
		}
		//4. the comment-summary is not in the comments: to remove.
		if _, ok := commentMap[each.CommentID]; !ok {
			toRemoveCommentIDs = append(toRemoveCommentIDs, each.CommentID)
		}
	}

	return toAddComments, toUpdateCommentSummaries, toRemoveCommentIDs
}
