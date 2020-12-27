package schema

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	IsTest = false
	lock   sync.Mutex
)

func SetIsTest() {
	lock.Lock()
	IsTest = true

	logrus.Info("schema: SetIsTest")

	_ = Init()
}

func UnsetIsTest() {
	Close()

	logrus.Info("schema: UnsetIsTest")

	IsTest = false
	lock.Unlock()
}
