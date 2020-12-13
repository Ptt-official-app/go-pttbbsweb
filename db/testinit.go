package db

import "sync"

var (
	isTest = false
	lock   sync.Mutex
)

func setupTest() {
	SetIsTest()
}

func teardownTest() {
	UnsetIsTest()
}

func SetIsTest() {
	lock.Lock()
	isTest = true
}

func UnsetIsTest() {
	isTest = false
	lock.Unlock()
}
