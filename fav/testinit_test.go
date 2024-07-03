package fav

import (
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func setupTest() {
	types.SetIsTest("fav")
}

func teardownTest() {
	types.UnsetIsTest("fav")
}
