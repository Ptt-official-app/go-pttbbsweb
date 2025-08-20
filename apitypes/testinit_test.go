package apitypes

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
)

func setupTest() {
	utils.SetIsTest()
	types.SetIsTest("apitypes")
}

func teardownTest() {
	defer utils.UnsetIsTest()
	defer types.UnsetIsTest("apitypes")
}
