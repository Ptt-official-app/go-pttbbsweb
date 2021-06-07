package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func updateArticleContentInfo(
	boardID bbs.BBoardID,
	articleID bbs.ArticleID,
	content [][]*types.Rune,
	contentMD5 string,
	ip string,
	host string,
	bbs string,
	signatureMD5 string,
	signatureDBCS []byte,
	updateNanoTS types.NanoTS) (err error) {

	if contentMD5 == "" {
		return nil
	}

	contentInfo := &schema.ArticleContentInfo{
		ContentMD5: contentMD5,

		Content: content,
		IP:      ip,
		Host:    host,
		BBS:     bbs,

		SignatureDBCS: signatureDBCS,
		SignatureMD5:  signatureMD5,

		ContentUpdateNanoTS: updateNanoTS,
	}

	err = schema.UpdateArticleContentInfo(boardID, articleID, contentInfo)

	return err
}

func deserializeArticlesAndUpdateDB(userID bbs.UUserID, bboardID bbs.BBoardID, articleSummaries_b []*bbs.ArticleSummary, updateNanoTS types.NanoTS) (articleSummaries []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]bool, err error) {
	if len(articleSummaries_b) == 0 {
		return nil, nil, nil
	}

	articleSummaries = make([]*schema.ArticleSummary, len(articleSummaries_b))
	for idx, each_b := range articleSummaries_b {
		articleSummaries[idx] = schema.NewArticleSummary(each_b, updateNanoTS)
	}

	err = schema.UpdateArticleSummaries(articleSummaries, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadArticles := make([]*schema.UserReadArticle, 0, len(articleSummaries_b))
	userReadArticleMap = make(map[bbs.ArticleID]bool)
	for _, each_b := range articleSummaries_b {
		if each_b.Read {
			each_db := &schema.UserReadArticle{
				UserID:       userID,
				ArticleID:    each_b.ArticleID,
				UpdateNanoTS: updateNanoTS,
			}
			userReadArticles = append(userReadArticles, each_db)

			userReadArticleMap[each_db.ArticleID] = true
		}
	}

	err = schema.UpdateUserReadArticles(userReadArticles, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	//get n-comments
	updateArticleNComments(bboardID, articleSummaries)

	return articleSummaries, userReadArticleMap, err
}

func updateArticleNComments(bboardID bbs.BBoardID, articleSummaries []*schema.ArticleSummary) {
	if len(articleSummaries) == 0 {
		return
	}

	articleIDs := make([]bbs.ArticleID, len(articleSummaries))
	for idx, each := range articleSummaries {
		articleIDs[idx] = each.ArticleID
	}

	articleNComments, err := schema.GetArticleNCommentsByArticleIDs(bboardID, articleIDs)
	if err != nil {
		return
	}

	nCommentsByArticleIDMap := make(map[bbs.ArticleID]*schema.ArticleNComments)
	for _, each := range articleNComments {
		nCommentsByArticleIDMap[each.ArticleID] = each
	}

	for _, each := range articleSummaries {
		eachArticleNComments := nCommentsByArticleIDMap[each.ArticleID]
		if eachArticleNComments == nil {
			continue
		}

		each.NComments = eachArticleNComments.NComments
		each.Rank = eachArticleNComments.Rank
	}
}
