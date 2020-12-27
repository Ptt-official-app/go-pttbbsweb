package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func tryUpdateFirstComments(
	firstComments []*schema.Comment,
	firstCommentsMD5 string,
	firstCommentsLastTime types.NanoTS,
	firstCommentsUpdateNanoTS types.NanoTS,
	articleDetailSummary *schema.ArticleDetailSummary) (

	err error) {
	if firstCommentsMD5 == articleDetailSummary.FirstCommentsMD5 {
		return nil
	}

	err = schema.UpdateComments(firstComments, firstCommentsUpdateNanoTS)
	if err != nil {
		return err
	}

	articleFirstComments := &schema.ArticleFirstComments{
		BBoardID:                  articleDetailSummary.BBoardID,
		ArticleID:                 articleDetailSummary.ArticleID,
		FirstCommentsMD5:          firstCommentsMD5,
		FirstCommentsLastTime:     firstCommentsLastTime,
		FirstCommentsUpdateNanoTS: firstCommentsUpdateNanoTS,
	}
	err = schema.UpdateArticleFirstComments(articleFirstComments)
	if err != nil {
		return err
	}

	//remove old first comments.
	err = schema.RemoveOldFirstComments(articleDetailSummary.BBoardID, articleDetailSummary.ArticleID, firstCommentsUpdateNanoTS)

	return nil
}
