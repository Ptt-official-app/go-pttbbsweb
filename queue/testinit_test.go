package queue

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("queue")

	schema.SetIsTest()

	SetIsTest()

	initTest()
}

func teardownTest() {
	UnsetIsTest()

	schema.UnsetIsTest()

	types.UnsetIsTest("queue")

	db.UnsetIsTest()

	utils.UnsetIsTest()
}
