package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	apiPrefix = "/api"
)

func withPrefix(path string) string {
	return apiPrefix + path
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

	//board
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_R), api.LoadGeneralBoardsWrapper)
	router.GET(withPrefix(api.GET_BOARD_DETAIL_R), api.GetBoardDetailWrapper)
	router.GET(withPrefix(api.GET_BOARD_SUMMARY_R), api.GetBoardSummaryWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_BOARDS_R), api.LoadPopularBoardsWrapper)

	//article
	router.GET(withPrefix(api.LOAD_GENERAL_ARTICLES_R), api.LoadGeneralArticlesWrapper)
	router.GET(withPrefix(api.LOAD_BOTTOM_ARTICLES_R), api.LoadBottomArticlesWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_R), api.GetArticleDetailWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_ARTICLES_R), api.LoadPopularArticlesWrapper)

	//user
	router.GET(withPrefix(api.GET_USER_INFO_R), api.GetUserInfoWrapper)
	router.GET(withPrefix(api.LOAD_FAVORITE_BOARDS_R), api.LoadFavoriteBoardsWrapper)
	router.GET(withPrefix(api.LOAD_USER_ARTICLES_R), api.LoadUserArticlesWrapper)

	//comments
	router.GET(withPrefix(api.LOAD_ARTICLE_FIRSTCOMMENTS_R), api.LoadArticleFirstCommentsWrapper)
	router.GET(withPrefix(api.LOAD_ARTICLE_COMMENTS_R), api.LoadArticleCommentsWrapper)
	router.GET(withPrefix(api.LOAD_USER_COMMENTS_R), api.LoadUserCommentsWrapper)

	router.GET("/static/js/*path", api.JSWrapper)
	router.GET("/", api.IndexHtmlWrapper)
	router.GET("/index.html", api.IndexHtmlWrapper)

	router.Static("/static/css", filepath.Join(types.STATIC_DIR, "static", "css"))
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

	err = db.InitConfig()
	if err != nil {
		return err
	}

	err = schema.InitConfig()
	if err != nil {
		return err
	}

	err = pttbbsapi.InitConfig()
	if err != nil {
		return err
	}

	err = pttbbstypes.InitConfig()
	if err != nil {
		return err
	}

	return nil
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

	queue.Init()

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
