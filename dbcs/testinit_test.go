package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
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
