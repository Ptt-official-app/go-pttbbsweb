package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const CREATE_RANK_R = "/board/:bid/article/:aid/rank"

type CreateRankParams struct {
	Rank int `json:"rank" form:"rank" url:"rank"`
}

type CreateRankPath struct {
	BoardID   bbs.BBoardID  `uri:"bid" binding:"required"`
	ArticleID bbs.ArticleID `uri:"aid" binding:"required"`
}

type CreateRankResult struct {
	Rank int `json:"rank"`
}

func CreateRankWrapper(c *gin.Context) {
	params := &CreateRankParams{}
	path := &CreateRankPath{}
	LoginRequiredPathJSON(CreateRank, params, path, c)
}

func CreateRank(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*CreateRankParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}
	thePath, ok := path.(*CreateRankPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	// check permission
	articleSummary, err := schema.GetArticleSummary(thePath.BoardID, thePath.ArticleID)
	if err != nil {
		return nil, 500, err
	}
	if articleSummary == nil {
		return nil, 400, ErrNoArticle
	}

	if articleSummary.Owner == userID {
		return nil, 403, ErrInvalidUser
	}

	isValid, statusCode, err := isBoardValidUser(thePath.BoardID, c)
	if err != nil {
		return nil, statusCode, err
	}
	if !isValid {
		return nil, 403, ErrInvalidUser
	}

	//update rank
	updateNanoTS := types.NowNanoTS()
	origRank, err := schema.UpdateRank(thePath.BoardID, thePath.ArticleID, userID, theParams.Rank, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	newRank, err := schema.UpdateArticleRank(thePath.BoardID, thePath.ArticleID, theParams.Rank-origRank, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	return &CreateRankResult{Rank: newRank}, 200, nil
}
