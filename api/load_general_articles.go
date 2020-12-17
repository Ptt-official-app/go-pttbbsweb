package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const LOAD_GENERAL_ARTICLES_R = "/board/:bid/articles"

type LoadGeneralArticlesParams struct {
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

func NewLoadGeneralArticlesParams() *LoadGeneralArticlesParams {
	return &LoadGeneralArticlesParams{
		Max: DEFAULT_MAX_LIST,
	}
}

type LoadGeneralArticlesPath struct {
	BBoardID bbs.BBoardID `uri:"bid"`
}

type LoadGeneralArticlesResult struct {
	List    []*types.ArticleSummary `json:"list"`
	NextIdx string                  `json:"next_idx"`
}

func LoadGeneralArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadGeneralArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	//backend load-general-articles
	theParams_b := &backend.LoadGeneralArticlesParams{
		StartIdx:  theParams.StartIdx,
		NArticles: theParams.Max,
	}
	var result_b *backend.LoadGeneralArticlesResult

	urlMap := make(map[string]string)
	urlMap["bid"] = string(thePath.BBoardID)
	url := backend.WithPrefix(utils.MergeURL(urlMap, backend.LOAD_GENERAL_ARTICLES_R))
	logrus.Infof("to backend: url: %v theParams_b: %v", url, theParams_b)
	statusCode, err = utils.HttpGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	r := &LoadGeneralArticlesResult{}
	r.Deserialize(result_b)

	//check isRead
	err = checkReadArticles(userID, r.List)
	if err != nil {
		return nil, 500, err
	}

	//update user_read_board
	if result_b.IsNewest { //nice to have, but not necessary.
		err = updateUserReadBoard(userID, thePath.BBoardID)
		if err != nil {
			return nil, 500, err
		}
	}

	return r, 200, nil
}

func checkReadArticles(userID bbs.UUserID, theList []*types.ArticleSummary) error {
	checkArticleIDMap := make(map[bbs.ArticleID]int)
	queryArticleIDs := make([]bbs.ArticleID, 0, len(theList))
	for idx, each := range theList {
		if each.Read {
			continue
		}

		//check with read-time
		checkArticleIDMap[each.ArticleID] = idx
		queryArticleIDs = append(queryArticleIDs, each.ArticleID)
	}

	//query
	query := make(map[string]interface{})
	query[schema.USER_READ_ARTICLE_USER_ID_b] = userID
	queryArticles := make(map[string]interface{})
	queryArticles["$in"] = queryArticleIDs
	query[schema.USER_READ_ARTICLE_ARTICLE_ID_b] = queryArticles

	var dbResults []*schema.UserReadArticle
	err := schema.UserReadArticle_c.Find(query, 0, &dbResults, nil)
	if err != nil {
		return err
	}

	for _, each := range dbResults {
		eachArticleID := each.ArticleID
		eachReadNanoTS := each.UpdateNanoTS

		listIdx := checkArticleIDMap[eachArticleID]
		eachInTheList := theList[listIdx]

		eachPostNanoTS := eachInTheList.CreateTime.ToNanoTS()
		eachInTheList.Read = eachReadNanoTS > eachPostNanoTS
	}

	return nil
}

func updateUserReadBoard(userID bbs.UUserID, BBoardID bbs.BBoardID) error {
	nowNanoTS := utils.GetNowNanoTS()

	query := &schema.UserReadBoard{
		UserID:       userID,
		BBoardID:     BBoardID,
		UpdateNanoTS: nowNanoTS,
	}

	_, err := schema.UserReadBoard_c.Update(query, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *LoadGeneralArticlesResult) Deserialize(r_b *backend.LoadGeneralArticlesResult) {

	r.List = make([]*types.ArticleSummary, len(r_b.Articles))
	for i := 0; i < len(r.List); i++ {
		each := &types.ArticleSummary{}
		each.Deserialize(r_b.Articles[i])
		r.List[i] = each
	}

	r.NextIdx = r_b.NextIdx
}
