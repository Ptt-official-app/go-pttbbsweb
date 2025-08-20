package dbcs

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
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
