package dbcs

import (
	"bytes"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/sirupsen/logrus"
)

//ParseFirstComments
//
//Check with origFirstCommentsMD5, if exists, return nil and requires getting firstComments and lastTime from db.
func ParseFirstComments(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	ownerID bbs.UUserID,
	articleCreateTime types.NanoTS,
	commentsDBCS []byte,
	origFirstCommentsMD5 string,
	origFirstCommentsLastTime types.NanoTS,
	updateNanoTS types.NanoTS) (

	firstComments []*schema.Comment,
	firstCommentsMD5 string,
	firstCommentsLastTime types.NanoTS,
	theRestCommentsDBCS []byte) {

	firstCommentsDBCS, theRestCommentsDBCS := splitFirstComments(commentsDBCS)

	//check md5
	firstCommentsMD5 = md5sum(firstCommentsDBCS)
	if firstCommentsMD5 == origFirstCommentsMD5 {
		return nil, origFirstCommentsMD5, origFirstCommentsLastTime, theRestCommentsDBCS
	}

	firstComments, firstCommentsLastTime = ParseComments(bboardID, articleID, ownerID, articleCreateTime, firstCommentsDBCS, commentsDBCS, updateNanoTS, true)

	return firstComments, firstCommentsMD5, firstCommentsLastTime, theRestCommentsDBCS
}

//splitFirstComments
//
//match the first N_FIRST_COMMENTS comments
func splitFirstComments(commentsDBCS []byte) (firstCommentsDBCS []byte, theRestComments []byte) {
	p_commentsDBCS := commentsDBCS

	nComments := 0
	nBytes := 0
	for idxNewLine := bytes.Index(p_commentsDBCS, []byte{'\n'}); len(p_commentsDBCS) > 0 && idxNewLine != -1 && nComments < N_FIRST_COMMENTS; {
		nComments++

		nBytes += idxNewLine
		p_commentsDBCS = p_commentsDBCS[idxNewLine:] //starting from '\n'

		nextCommentIdx := MatchComment(p_commentsDBCS)
		logrus.Infof("splitFirstComments: nComments: %v nextCommentIdx: %v", nComments, nextCommentIdx)
		if nextCommentIdx == -1 {
			break
		}

		nBytes += nextCommentIdx
		p_commentsDBCS = p_commentsDBCS[nextCommentIdx:] //starting from beginning of the next comment.

		idxNewLine = bytes.Index(p_commentsDBCS, []byte{'\n'})
	}

	if nComments < N_FIRST_COMMENTS { //no more '\n', but not enough comments yet, add the last comment.
		nBytes += len(p_commentsDBCS)
	}

	//defensive programming for '\n'
	if nBytes < len(commentsDBCS) && commentsDBCS[nBytes] == '\n' {
		nBytes++
	}
	firstCommentsDBCS, theRestComments = commentsDBCS[:nBytes], commentsDBCS[nBytes:]
	if len(firstCommentsDBCS) == 0 {
		firstCommentsDBCS = nil
	}
	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return firstCommentsDBCS, theRestComments
}
