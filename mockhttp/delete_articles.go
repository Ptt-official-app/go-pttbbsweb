package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func DeleteArticles(params *api.DeleteArticlesParams) (ret *api.DeleteArticlesResult) {
	ret = &api.DeleteArticlesResult{
		ArticleIDs: params.ArticleIDs,
	}
	return ret
}