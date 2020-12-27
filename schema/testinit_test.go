package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("schema")

	SetIsTest()

}

func teardownTest() {
	UnsetIsTest()

	types.UnsetIsTest("schema")

	db.UnsetIsTest()

	utils.UnsetIsTest()

}
