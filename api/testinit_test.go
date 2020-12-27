package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/sirupsen/logrus"
)

var (
	testIP = "127.0.0.1"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("api")

	schema.SetIsTest()

	queue.SetIsTest()

	params := &RegisterClientParams{ClientID: "default_client_id"}
	_, statusCode, err := RegisterClient("localhost", bbs.UUserID("SYSOP"), params, nil)
	logrus.Infof("api.setupTest: after RegisterClient: status: %v e: %v", statusCode, err)
}

func teardownTest() {
	queue.UnsetIsTest()

	schema.UnsetIsTest()

	types.UnsetIsTest("api")

	db.UnsetIsTest()

	utils.UnsetIsTest()
}
