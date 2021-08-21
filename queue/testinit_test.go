package queue

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

var IsTest = false

func SetIsTest() {
	IsTest = true

	Start()
}

func UnsetIsTest() {
	IsTest = false

	Close()
}

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
