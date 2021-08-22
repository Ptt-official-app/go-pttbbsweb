package apitypes

import (
	"fmt"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"strings"
)

func ToURL(fbboardID FBoardID, farticleID FArticleID) string {
	// ALLPOST need to translate to original board
	if fbboardID == "ALLPOST" {
		articleID := farticleID.ToArticleID()
		articleSummaries, err := schema.GetArticleSummariesByArticleIDs([]bbs.ArticleID{articleID})
		if err != nil {
			//TODO error handling
			fmt.Println(err)
			return ""
		}
		//TODO if not find?
		for _, v := range articleSummaries {
			b := strings.Join(strings.Split(string(v.BBoardID), "_")[1:], "")
			if b == "ALLPOST" {
				continue
			}
			return types.FRONTEND_PREFIX + types.API_PREFIX + "/" + "board" + "/" + b + "/" + "article" + "/" + string(farticleID)
		}
		return ""
	}
	return types.FRONTEND_PREFIX + types.API_PREFIX + "/" + "board" + "/" + string(fbboardID) + "/" + "article" + "/" + string(farticleID)
}
