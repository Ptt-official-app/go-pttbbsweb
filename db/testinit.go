package db

var (
	isTest = false
)

func setupTest() {
	SetIsTest()
}

func teardownTest() {
	UnsetIsTest()
}

func SetIsTest() {
	isTest = true
}

func UnsetIsTest() {
	isTest = false
}
