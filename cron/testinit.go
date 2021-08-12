package cron

var IsTest = false

func SetIsTest() {
	IsTest = true
}

func UnsetIsTest() {
	IsTest = false
}
