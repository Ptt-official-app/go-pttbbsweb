package cron

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("api")

	schema.SetIsTest()

	queue.SetIsTest()

	api.SetIsTest()

	params := &api.RegisterClientParams{ClientID: "default_client_id", ClientType: types.CLIENT_TYPE_APP}
	_, _, _ = api.RegisterClient("localhost", bbs.UUserID("SYSOP"), params, nil)
	// logrus.Infof("api.setupTest: after RegisterClient: status: %v e: %v", statusCode, err)
}

func teardownTest() {
	defer utils.UnsetIsTest()

	defer db.UnsetIsTest()

	defer types.UnsetIsTest("api")

	defer schema.UnsetIsTest()

	defer queue.UnsetIsTest()

	defer api.UnsetIsTest()
}
