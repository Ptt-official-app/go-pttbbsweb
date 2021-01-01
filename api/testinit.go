package api

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

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
}
