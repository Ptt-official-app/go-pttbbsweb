package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func setupTest() {
	initTest()
	types.SetIsTest("dbcs")
	db.SetIsTest()
	schema.SetIsTest()
}

func teardownTest() {
	schema.UnsetIsTest()
	db.UnsetIsTest()
	types.UnsetIsTest("dbcs")
}
