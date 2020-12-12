package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

var (
	testIP = "127.0.0.1"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	schema.Init()

	params := &RegisterClientParams{ClientID: "default_client_id"}
	_, _, _ = RegisterClient("localhost", params, nil)
}

func teardownTest() {
	schema.Close()

	db.UnsetIsTest()

	utils.UnsetIsTest()
}
