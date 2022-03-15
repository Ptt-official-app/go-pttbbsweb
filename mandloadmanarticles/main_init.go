package main

import (
	"flag"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

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

	logrus.Debugf("viper keys: %v", viper.AllKeys())

	err = types.InitConfig()
	if err != nil {
		return err
	}

	err = mand.InitConfig()
	if err != nil {
		return err
	}

	return nil
}

func initMain() (brdname string, path string, err error) {
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)
	logrus.SetLevel(logrus.InfoLevel)

	iniFilename := ""
	flag.StringVar(&iniFilename, "ini", "config.ini", "ini filename")

	flag.StringVar(&brdname, "brdname", "", "brdname")

	flag.StringVar(&path, "path", "", "path")

	flag.Parse()

	err = initAllConfig(iniFilename)
	if err != nil {
		return "", "", err
	}

	err = mand.Init(false)
	if err != nil {
		return "", "", err
	}

	return brdname, path, nil
}
