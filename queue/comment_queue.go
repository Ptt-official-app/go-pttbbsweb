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
	BBoardID          bbs.BBoardID
	ArticleID         bbs.ArticleID
	OwnerID           bbs.UUserID
	CommentDBCS       []byte
	ArticleCreateTime types.NanoTS
	ArticleMTime      types.NanoTS
	UpdateNanoTS      types.NanoTS
}

func QueueCommentDBCS(bboardID bbs.BBoardID, articleID bbs.ArticleID, ownerID bbs.UUserID, commentDBCS []byte, articleCreateTime types.NanoTS, articleMTime types.NanoTS, updateNanoTS types.NanoTS) (err error) {
	commentQueue := &CommentQueue{
		BBoardID:          bboardID,
		ArticleID:         articleID,
		OwnerID:           ownerID,
		CommentDBCS:       commentDBCS,
		ArticleCreateTime: articleCreateTime,
		ArticleMTime:      articleMTime,
		UpdateNanoTS:      updateNanoTS,
	}

	// log.Infof("QueueCommentDBCS: to send to queue: bboardID: %v articleID: %v", bboardID, articleID)
	select {
	case theQueue <- commentQueue:
	case <-time.After(defaultTimeout):
		err = ErrTimeout{
			timeout: defaultTimeout,
		}
	}

	// log.Infof("QueueCommentDBCS: after send to queue: bboardID: %v articleID: %v e: %v", bboardID, articleID, err)

	return err
}

func ProcessCommentQueueLoop(idx int, quit chan struct{}) {
	for {
		select {
		case commentQueue, ok := <-theQueue:
			if !ok {
				return
			}

			_ = ProcessCommentQueue(commentQueue)
		case <-quit:
			return
		}
	}
}

//ProcessCommentQueue
//
//We use LastTime as the reference time to obtain stable timestamp.
//(lastTime is from firstComments, assuming not change a lot.)
//(mtime changes frequently and may result in unstable timestamp.)
func ProcessCommentQueue(q *CommentQueue) (err error) {
	// 1. parse comments.
	comments := dbcs.ParseComments(q.OwnerID, q.CommentDBCS, q.CommentDBCS)
	// log.Infof("processCommentQueue: after parseComments: comments: %v", len(comments))
	if len(comments) == 0 {
		return nil
	}

	// 2. integrate comments.
	toAddComments, toDeleteComments, err := dbcs.IntegrateComments(q.BBoardID, q.ArticleID, comments, q.ArticleCreateTime, q.ArticleMTime, false, true)
	if err != nil {
		return err
	}

	// 3. remove comment-ids first.
	toRemoveCommentIDs := make([]types.CommentID, len(toDeleteComments))
	for idx, each := range toDeleteComments {
		toRemoveCommentIDs[idx] = each.CommentID
	}

	err = schema.RemoveCommentIDs(q.BBoardID, q.ArticleID, toRemoveCommentIDs, q.UpdateNanoTS, "not-in-file")
	if err != nil {
		return err
	}

	// 4. update comments.
	err = schema.UpdateComments(toAddComments, q.UpdateNanoTS)
	if err != nil {
		return err
	}

	// 5. update article comments.
	return schema.UpdateArticleCommentsByArticleID(q.BBoardID, q.ArticleID, q.UpdateNanoTS)
}
