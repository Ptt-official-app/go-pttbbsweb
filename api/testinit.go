package api

import (
	"os"

	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var (
	IsTest        = false
	origStaticDir = ""
)

func SetIsTest() {
	IsTest = true
	origStaticDir = types.STATIC_DIR
	types.STATIC_DIR = "testcase"
}

func UnsetIsTest() {
	types.STATIC_DIR = origStaticDir
	IsTest = false

	data, _ := os.ReadFile("./testcase/home2/t/testUser2/.fav.orig")

	_ = os.WriteFile("./testcase/home2/t/testUser2/.fav", data, 0o644)
}
