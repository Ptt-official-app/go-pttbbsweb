package schema

import (
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/Ptt-official-app/go-pttbbsweb/utils"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("schema")

	SetIsTest()

	initTest11()
}

func teardownTest() {
	UnsetIsTest()

	types.UnsetIsTest("schema")

	db.UnsetIsTest()

	utils.UnsetIsTest()
}
