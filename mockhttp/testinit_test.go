package mockhttp

import log "github.com/sirupsen/logrus"

func setupTest() {
	log.SetLevel(log.DebugLevel)
}

func teardownTest() {
}
