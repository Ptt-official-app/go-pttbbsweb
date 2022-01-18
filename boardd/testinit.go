package boardd

import "sync"

var (
	IsTest    = false
	TestMutex sync.Mutex
)

func SetIsTest() {
	TestMutex.Lock()

	IsTest = true

	Init(false)
}

func UnsetIsTest() {
	defer TestMutex.Unlock()

	defer func() {
		IsTest = false
	}()

	defer Finalize(false)
}
