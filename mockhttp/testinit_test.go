package mockhttp

import "github.com/sirupsen/logrus"

func setupTest() {
	logrus.SetLevel(logrus.DebugLevel)
}

func teardownTest() {
}
