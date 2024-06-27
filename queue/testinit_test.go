package queue

import (
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/Ptt-official-app/go-pttbbsweb/utils"
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
