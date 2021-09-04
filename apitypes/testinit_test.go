package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

func setupTest() {
	utils.SetIsTest()
	types.SetIsTest("apitypes")
}

func teardownTest() {
	defer utils.UnsetIsTest()
	defer types.UnsetIsTest("apitypes")
}
