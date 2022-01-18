package boardd

func setupTest() {
	SetIsTest()

	initVars()
}

func teardownTest() {
	defer UnsetIsTest()

	defer freeTestVars()
}
