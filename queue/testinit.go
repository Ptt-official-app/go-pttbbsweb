package queue

var (
	IsTest = false
)

func SetIsTest() {
	IsTest = true

	Init()

}

func UnsetIsTest() {
	IsTest = false
}
