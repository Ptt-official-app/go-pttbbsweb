package fav

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func setupTest() {
	types.SetIsTest("fav")
}

func teardownTest() {
	types.UnsetIsTest("fav")
}
