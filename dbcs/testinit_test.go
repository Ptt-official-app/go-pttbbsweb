package dbcs

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

func setupTest() {
	initTest()
	types.SetIsTest("dbcs")
}

func teardownTest() {
	types.UnsetIsTest("dbcs")
}
