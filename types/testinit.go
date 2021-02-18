package types

import "sync"

var (
	origBig5ToUtf8                  = ""
	origUtf8ToBig5                  = ""
	origEmailTokenTemplate          = ""
	origIDEmailTokenTemplate        = ""
	origAttemptRegisterUserTemplate = ""

	lock sync.Mutex
)

func setupTest() {
	SetIsTest("types")
}

func teardownTest() {
	UnsetIsTest("types")
}

func SetIsTest(pkgName string) {
	lock.Lock()
	if pkgName != "main" {
		origBig5ToUtf8 = BIG5_TO_UTF8
		origUtf8ToBig5 = UTF8_TO_BIG5

		origEmailTokenTemplate = EMAILTOKEN_TEMPLATE
		origIDEmailTokenTemplate = IDEMAILTOKEN_TEMPLATE
		origAttemptRegisterUserTemplate = ATTEMPT_REGISTER_USER_TEMPLATE

		BIG5_TO_UTF8 = "../types/uao250-b2u.big5.txt"
		UTF8_TO_BIG5 = "../types/uao250-u2b.big5.txt"

		EMAILTOKEN_TEMPLATE = "../etc/emailtoken.template"
		IDEMAILTOKEN_TEMPLATE = "../etc/idemailtoken.template"
		ATTEMPT_REGISTER_USER_TEMPLATE = "../etc/attemptregister.template"

		initBig5()
	}
}

func UnsetIsTest(pkgName string) {
	defer lock.Unlock()

	if pkgName != "main" {
		BIG5_TO_UTF8 = origBig5ToUtf8
		UTF8_TO_BIG5 = origUtf8ToBig5

		ATTEMPT_REGISTER_USER_TEMPLATE = origAttemptRegisterUserTemplate
		IDEMAILTOKEN_TEMPLATE = origIDEmailTokenTemplate
		EMAILTOKEN_TEMPLATE = origEmailTokenTemplate
	}
}
