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

var (
	getFavoriteBoard_r   = withPrefix("/Board/favorite/:username")
	getUserPostList_r    = withPrefix("/User/article/:username")
	getUserCommentList_r = withPrefix("/User/Comment/:username")
	getUserInfo_r        = withPrefix("/User/Users/:username")
)

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

	//user
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
			api.NewGetBoardDetailParams(),
			&api.GetBoardDetailPath{},
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
