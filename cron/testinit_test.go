package cron

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/api"
	"github.com/Ptt-official-app/go-pttbbsweb/boardd"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/queue"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/Ptt-official-app/go-pttbbsweb/utils"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("api")

	schema.SetIsTest()

	queue.SetIsTest()

	boardd.SetIsTest()

	api.SetIsTest()

	SetIsTest()

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

	defer boardd.UnsetIsTest()

	defer api.UnsetIsTest()

	defer UnsetIsTest()
}
