package cron

import (
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/queue"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
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

	user := &api.UserInfo{UserID: "SYSOP", IsOver18: true}
	_, _, _ = api.RegisterClient("localhost", user, params, nil)
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
