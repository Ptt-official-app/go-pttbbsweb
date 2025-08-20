package schema

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
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
