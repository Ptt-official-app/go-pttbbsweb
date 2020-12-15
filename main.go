package main

import (
	"flag"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	apiPrefix = "/api"
)

var ()

func withPrefix(path string) string {
	return apiPrefix + path
}

func initGin() (*gin.Engine, error) {
	router := gin.Default()

	//index
	router.GET(api.INDEX_R,
		NewApi(
			api.Index,
			api.NewIndexParams(),
		).Query,
	)

	//register/login
	router.POST(
		withPrefix(api.REGISTER_CLIENT_R),
		NewApi(
			api.RegisterClient,
			api.NewRegisterClientParams(),
		).Json,
	)
	router.POST(
		withPrefix(api.REGISTER_USER_R),
		NewApi(
			api.RegisterUser,
			api.NewRegisterUserParams(),
		).Json,
	)
	router.POST(
		withPrefix(api.LOGIN_R),
		NewApi(
			api.Login,
			api.NewLoginParams(),
		).Json,
	)

	//board
	router.GET(
		withPrefix(api.LOAD_GENERAL_BOARDS_R),
		NewLoginRequiredApi(
			api.LoadGeneralBoards,
			api.NewLoadGeneralBoardsParams(),
		).Query,
	)
	router.GET(
		withPrefix(api.GET_BOARD_DETAIL_R),
		NewLoginRequiredPathApi(
			api.GetBoardDetail,
			&api.GetBoardDetailParams{},
			&api.GetBoardDetailPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.GET_BOARD_SUMMARY_R),
		NewLoginRequiredPathApi(
			api.GetBoardSummary,
			&api.GetBoardSummaryParams{},
			&api.GetBoardSummaryPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_POPULAR_BOARDS_R),
		NewLoginRequiredApi(
			api.LoadPopularBoards,
			&api.LoadPopularBoardsParams{},
		).Query,
	)

	//article
	router.GET(
		withPrefix(api.LOAD_GENERAL_ARTICLES_R),
		NewLoginRequiredPathApi(
			api.LoadGeneralArticles,
			api.NewLoadGeneralArticlesParams(),
			&api.LoadGeneralArticlesPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_BOTTOM_ARTICLES_R),
		NewLoginRequiredPathApi(
			api.LoadBottomArticles,
			&api.LoadBottomArticlesParams{},
			&api.LoadBottomArticlesPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.GET_ARTICLE_R),
		NewLoginRequiredPathApi(
			api.GetArticleDetail,
			&api.GetArticleDetailParams{},
			&api.GetArticleDetailPath{},
		).Query,
	)

	router.GET(
		withPrefix(api.LOAD_POPULAR_ARTICLES_R),
		NewLoginRequiredApi(
			api.LoadPopularArticles,
			&api.LoadPopularArticlesParams{},
		).Query,
	)

	//user
	router.GET(
		withPrefix(api.GET_USER_INFO_R),
		NewLoginRequiredPathApi(
			api.GetUserInfo,
			&api.GetUserInfoParams{},
			&api.GetUserInfoPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_FAVORITE_BOARDS_R),
		NewLoginRequiredPathApi(
			api.LoadFavoriteBoards,
			&api.LoadFavoriteBoardsParams{},
			&api.LoadFavoriteBoardsPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_USER_ARTICLES_R),
		NewLoginRequiredPathApi(
			api.LoadUserArticles,
			&api.LoadUserArticlesParams{},
			&api.LoadUserArticlesPath{},
		).Query,
	)

	//comments
	router.GET(
		withPrefix(api.LOAD_ARTICLE_FIRSTCOMMENTS_R),
		NewLoginRequiredPathApi(
			api.LoadArticleFirstComments,
			&api.LoadArticleFirstCommentsParams{},
			&api.LoadArticleFirstCommentsPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_ARTICLE_COMMENTS_R),
		NewLoginRequiredPathApi(
			api.LoadArticleComments,
			&api.LoadArticleCommentsParams{},
			&api.LoadArticleCommentsPath{},
		).Query,
	)
	router.GET(
		withPrefix(api.LOAD_USER_COMMENTS_R),
		NewLoginRequiredPathApi(
			api.LoadUserComments,
			&api.LoadUserCommentsParams{},
			&api.LoadUserCommentsPath{},
		).Query,
	)

	return router, nil
}

//Params
//      filename: ini filename
//
//Return
//      error: err
func initAllConfig(filename string) error {

	filenameList := strings.Split(filename, ".")
	if len(filenameList) == 1 {
		return ErrInvalidIni
	}

	filenamePrefix := strings.Join(filenameList[:len(filenameList)-1], ".")
	filenamePostfix := filenameList[len(filenameList)-1]
	viper.SetConfigName(filenamePrefix)
	viper.SetConfigType(filenamePostfix)
	viper.AddConfigPath("/etc/go-openbbsmiddleware")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	log.Debugf("viper keys: %v", viper.AllKeys())

	err = types.InitConfig()
	if err != nil {
		return err
	}

	err = backend.InitConfig()
	if err != nil {
		return err
	}

	err = db.InitConfig()
	if err != nil {
		return err
	}

	err = schema.InitConfig()
	if err != nil {
		return err
	}

	return InitConfig()
}

func initMain() error {
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)
	log.SetLevel(log.InfoLevel)

	filename := ""
	flag.StringVar(&filename, "ini", "config.ini", "ini filename")
	flag.Parse()

	err := initAllConfig(filename)
	if err != nil {
		return err
	}

	err = schema.Init()
	if err != nil {
		return err
	}

	return nil
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
