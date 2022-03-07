package dbcs

import (
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

//ParseFirstComments
//
//Check with origFirstCommentsMD5, if exists, return nil and requires getting firstComments and lastTime from db.
func ParseFirstCommentsStr(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	ownerID bbs.UUserID,
	articleCreateTime types.NanoTS,
	articleMTime types.NanoTS,
	commentsDBCS string,
	origFirstCommentsMD5 string) (

	firstComments []*schema.Comment,
	firstCommentsMD5 string,
	theRestCommentsDBCS string,
	err error) {

	firstCommentsDBCS, theRestCommentsDBCS := splitFirstCommentsStr(commentsDBCS)

	// check md5
	firstCommentsMD5 = Md5sum([]byte(firstCommentsDBCS))
	if firstCommentsMD5 == origFirstCommentsMD5 {
		return nil, origFirstCommentsMD5, theRestCommentsDBCS, nil
	}

	comments := ParseCommentsStr(ownerID, firstCommentsDBCS, commentsDBCS)

	isLastAlignEndNanoTS := len(theRestCommentsDBCS) == 0

	firstComments, _, err = IntegrateComments(bboardID, articleID, comments, articleCreateTime, articleMTime, true, isLastAlignEndNanoTS)

	return firstComments, firstCommentsMD5, theRestCommentsDBCS, err
}

//splitFirstCommentsStr
//
//match the first N_FIRST_COMMENTS comments
func splitFirstCommentsStr(commentsDBCS string) (firstCommentsDBCS string, theRestComments string) {
	p_commentsDBCS := commentsDBCS

	nComments := 0
	nBytes := 0
	for idxNewLine := strings.Index(p_commentsDBCS, "\n"); len(p_commentsDBCS) > 0 && idxNewLine != -1 && nComments < N_FIRST_COMMENTS; {
		nComments++

		nBytes += idxNewLine
		p_commentsDBCS = p_commentsDBCS[idxNewLine:] // starting from '\n'

		nextCommentIdx := MatchCommentStr(p_commentsDBCS)
		if nextCommentIdx == -1 {
			break
		}

		nBytes += nextCommentIdx
		p_commentsDBCS = p_commentsDBCS[nextCommentIdx:] // starting from beginning of the next comment.

		idxNewLine = strings.Index(p_commentsDBCS, "\n")
	}

	if nComments < N_FIRST_COMMENTS { // no more '\n', but not enough comments yet, add the last comment.
		nBytes += len(p_commentsDBCS)
	}

	// defensive programming for '\n'
	if nBytes < len(commentsDBCS) && commentsDBCS[nBytes] == '\n' {
		nBytes++
	}
	firstCommentsDBCS, theRestComments = commentsDBCS[:nBytes], commentsDBCS[nBytes:]

	return firstCommentsDBCS, theRestComments
}
