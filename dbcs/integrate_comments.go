package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func IntegrateComments(boardID bbs.BBoardID, articleID bbs.ArticleID, comments []*schema.Comment, articleCreateTime types.NanoTS, articleMTime types.NanoTS, isForwardOnly bool, isLastAlignEndNanoTS bool) (newComments []*schema.Comment, toDeleteComments []*schema.CommentMD5, err error) {
	origComments, err := schema.GetAllCommentMD5s(boardID, articleID)
	if err != nil {
		return nil, nil, err
	}
	edBlocks, err := CalcEDBlocks(comments, origComments, articleCreateTime, articleMTime)
	if err != nil {
		return nil, nil, err
	}

	nBlock := InferTimestamp(edBlocks, isForwardOnly, isLastAlignEndNanoTS)
	//count new comments
	nNewComments := 0
	for idx, each := range edBlocks {
		if idx == nBlock {
			break
		}
		nNewComments += len(each.NewComments)
	}

	if nNewComments != 0 {
		newComments = make([]*schema.Comment, 0, nNewComments)
		for idx, each := range edBlocks {
			if idx == nBlock {
				break
			}
			for _, each2 := range each.NewComments {
				newComments = append(newComments, each2.NewComment)
			}
		}
	}

	// board-id and article-id
	for _, each := range newComments {
		each.BBoardID = boardID
		each.ArticleID = articleID
	}

	//early-return for isForawrdOnly
	if isForwardOnly {
		return newComments, nil, nil
	}

	//to delete comments
	nToDeleteComments := 0
	for idx, each := range edBlocks {
		if idx == nBlock {
			break
		}
		nToDeleteComments += len(each.OrigComments)
	}

	if nToDeleteComments != 0 {
		toDeleteComments = make([]*schema.CommentMD5, 0, nToDeleteComments)
		for idx, each := range edBlocks {
			if idx == nBlock {
				break
			}
			for _, each2 := range each.OrigComments {
				toDeleteComments = append(toDeleteComments, each2.OrigComment)
			}
		}
	}

	return newComments, toDeleteComments, nil
}
