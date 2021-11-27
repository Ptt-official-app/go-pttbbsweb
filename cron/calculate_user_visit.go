package cron

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"

	"github.com/sirupsen/logrus"
)

func RetryCalculateUserVisit() {
	for {
		logrus.Infof("RetryCalculateUserVisit: to calculate user visit")
		CalculateUserVisit()
		logrus.Infof("RetryCalculateUserVisit: to sleep 10 mins")
		time.Sleep(10 * time.Minute)
	}
}

func CalculateUserVisit() {
	count, err := schema.CalculateAllUserVisitCounts()
	if err != nil {
		logrus.Printf("get error in calculate user visit count %v", err)
	}
	logrus.Infof("user visit %d", count)
	// set to redis
	err = schema.SetUserVisitCount(count)
	if err != nil {
		logrus.Printf("set to redis  %v", err)
	}
}
