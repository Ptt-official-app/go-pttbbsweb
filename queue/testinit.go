package queue

var IsTest = false

func SetIsTest() {
	IsTest = true

	Start()
}

func UnsetIsTest() {
	IsTest = false

	Close()
}
