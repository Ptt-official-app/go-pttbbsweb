package fav

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func setupTest() {
	types.SetIsTest("fav")
}

func teardownTest() {
	types.UnsetIsTest("fav")
}
