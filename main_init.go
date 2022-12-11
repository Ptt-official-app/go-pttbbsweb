package main

import (
	"flag"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/boardd"
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/spf13/viper"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
)

// Params
//
//	filename: ini filename
//
// Return
//
//	error: err
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

	err = queue.InitConfig()
	if err != nil {
		return err
	}

	err = boardd.InitConfig()
	if err != nil {
		return err
	}

	err = mand.InitConfig()
	if err != nil {
		return err
	}

	err = dbcs.InitConfig()
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

	err = boardd.Init(false)
	if err != nil {
		return err
	}

	err = mand.Init(false)
	if err != nil {
		return err
	}

	logrus.Infof("initMain: done: MAX_USERS: %v MAX_ACTIVE: %v MAX_BOARD: %v HOTBOARDCACHE: %v", ptttype.MAX_USERS, ptttype.MAX_ACTIVE, ptttype.MAX_BOARD, ptttype.HOTBOARDCACHE)

	return nil
}
