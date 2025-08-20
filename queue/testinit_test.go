package queue

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
)

func setupTest() {
	utils.SetIsTest()

	types.SetIsTest("queue")

	db.SetIsTest()

	schema.SetIsTest()

	SetIsTest()

	initTest()
}

func teardownTest() {
	UnsetIsTest()

	schema.UnsetIsTest()

	db.UnsetIsTest()

	types.UnsetIsTest("queue")

	utils.UnsetIsTest()
}
