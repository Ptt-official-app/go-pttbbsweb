package main

import (
	"path/filepath"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func withPrefix(path string) string {
	return types.API_PREFIX + path
}

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	//options
	router.OPTIONS("/*path", api.OptionsWrapper)

	//index
	router.GET(withPrefix(api.INDEX_R), api.IndexWrapper)

	//register/login
	router.POST(withPrefix(api.REGISTER_CLIENT_R), api.RegisterClientWrapper)
	router.POST(withPrefix(api.REGISTER_USER_R), api.RegisterUserWrapper)
	router.POST(withPrefix(api.LOGIN_R), api.LoginWrapper)
	router.POST(withPrefix(api.ATTEMPT_REGISTER_USER_R), api.AttemptRegisterUserWrapper)

	//board
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_R), api.LoadGeneralBoardsWrapper)
	router.GET(withPrefix(api.GET_BOARD_DETAIL_R), api.GetBoardDetailWrapper)
	router.GET(withPrefix(api.GET_BOARD_SUMMARY_R), api.GetBoardSummaryWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_BOARDS_R), api.LoadPopularBoardsWrapper)
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_BY_CLASS_R), api.LoadGeneralBoardsByClassWrapper)
	router.GET(withPrefix(api.LOAD_AUTO_COMPLETE_BOARDS_R), api.LoadAutoCompleteBoardsWrapper)
	router.POST(withPrefix(api.CREATE_BOARD_R), api.CreateBoardWrapper)

	//article
	router.GET(withPrefix(api.LOAD_GENERAL_ARTICLES_R), api.LoadGeneralArticlesWrapper)
	router.GET(withPrefix(api.LOAD_BOTTOM_ARTICLES_R), api.LoadBottomArticlesWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_R), api.GetArticleDetailWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_ARTICLES_R), api.LoadPopularArticlesWrapper)
	router.POST(withPrefix(api.CREATE_ARTICLE_R), api.CreateArticleWrapper)

	//user
	router.GET(withPrefix(api.GET_USER_INFO_R), api.GetUserInfoWrapper)
	router.GET(withPrefix(api.LOAD_FAVORITE_BOARDS_R), api.LoadFavoriteBoardsWrapper)
	router.GET(withPrefix(api.LOAD_USER_ARTICLES_R), api.LoadUserArticlesWrapper)
	router.POST(withPrefix(api.CHANGE_PASSWD_R), api.ChangePasswdWrapper)
	router.POST(withPrefix(api.ATTEMPT_CHANGE_EMAIL_R), api.AttemptChangeEmailWrapper)
	router.POST(withPrefix(api.CHANGE_EMAIL_R), api.ChangeEmailWrapper)
	router.POST(withPrefix(api.ATTEMPT_SET_ID_EMAIL_R), api.AttemptSetIDEmailWrapper)
	router.POST(withPrefix(api.SET_ID_EMAIL_R), api.SetIDEmailWrapper)

	//comments
	router.GET(withPrefix(api.LOAD_ARTICLE_COMMENTS_R), api.LoadArticleCommentsWrapper)
	router.GET(withPrefix(api.LOAD_USER_COMMENTS_R), api.LoadUserCommentsWrapper)

	//html
	router.GET(api.ROOT_HTML_R, api.IndexHtmlWrapper)
	router.GET(api.INDEX_HTML_R, api.IndexHtmlWrapper)

	router.GET(api.REGISTER_HTML_R, api.RegisterHtmlWrapper)
	router.GET(api.LOGIN_HTML_R, api.LoginHtmlWrapper)

	router.GET(api.USER_HTML_R, api.UserHtmlWrapper)
	router.GET(api.USER_CHANGE_PASSWD_HTML_R, api.UserChangePasswdHtmlWrapper)
	router.GET(api.USER_ATTEMPT_CHANGE_EMAIL_HTML_R, api.UserAttemptChangeEmailHtmlWrapper)
	router.GET(api.USER_CHANGE_EMAIL_HTML_R, api.UserChangeEmailHtmlWrapper)
	router.GET(api.USER_ATTEMPT_SET_ID_EMAIL_HTML_R, api.UserAttemptSetIDEmailHtmlWrapper)
	router.GET(api.USER_SET_ID_EMAIL_HTML_R, api.UserSetIDEmailHtmlWrapper)

	router.GET(api.BOARDS_FAVORITES_HTML_R, api.AllHtmlWrapper)
	router.GET(api.BOARDS_POPULAR_HTML_R, api.AllHtmlWrapper)
	router.GET(api.BOARDS_HTML_R, api.AllHtmlWrapper)

	router.GET(api.ARTICLES_HTML_R, api.AllHtmlWrapper)
	router.GET(api.ARTICLE_HTML_R, api.AllHtmlWrapper)

	router.Static("/static", filepath.Join(types.STATIC_DIR, "static"))

	staticFiles := []string{
		"asset-manifest.json",
		"favicon.ico",
		"logo192.png",
		"logo512.png",
		"manifest.json",
		"robots.txt",
	}

	for _, each := range staticFiles {
		router.StaticFile("/"+each, filepath.Join(types.STATIC_DIR, each))
	}

	return router, nil
}

func main() {
	err := initMain()
	if err != nil {
		log.Errorf("unable to initMain: e: %v", err)
		return
	}
	router, err := initGin()
	if err != nil {
		return
	}

	_ = router.Run(types.HTTP_HOST)
}
