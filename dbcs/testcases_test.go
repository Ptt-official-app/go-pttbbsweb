package dbcs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initTest() {
	if testContentAll0 != nil {
		return
	}

	initTest0()
	initTest1()
	initTest2()
	initTest3()
	initTest4()
	initTest5()
	initTest6()
	initTest7()
	initTest8()
	initTest9()
	initTest10()
	initTest11()
	initTest12()
	initTest13()
	initTest14()
	initTest15()
	initTest16()
	initTest17()
	initTest18()
	initTest19()
	initTest20()
	initTest21()
	initTestUtf80()
	initTestUtf81()
	initTestUtf82()
	initTestUtf83()
	initTestUtf84()
	initTestUtf85()
	initTestUtf86()
	initTestUtf87()
	initTestUtf88()
}

func loadTest(filename string) (contentAll []byte, content []byte, signature []byte, recommend []byte, firstComments []byte, theRestComments []byte) {
	// content-all
	fullFilename := "testcase/" + filename
	contentAll, err := os.ReadFile(fullFilename)
	if err != nil {
		logrus.Errorf("loadTest: unable to open: filename: %v e: %v", filename, err)
		return nil, nil, nil, nil, nil, nil
	}

	// content
	fullFilename = "testcase/" + filename + ".content"
	content, err = os.ReadFile(fullFilename)
	if err != nil {
		logrus.Errorf("loadTest: unable to open content: filename : %v e: %v", filename, err)
		return nil, nil, nil, nil, nil, nil
	}

	if len(content) == 0 {
		content = nil
	}

	// signature
	fullFilename = "testcase/" + filename + ".signature"
	signature, err = os.ReadFile(fullFilename)
	if err != nil {
		logrus.Errorf("loadTest: unable to open signature: filename : %v e: %v", filename, err)
		return nil, nil, nil, nil, nil, nil
	}

	if len(signature) == 0 {
		signature = nil
	}

	// recommend
	fullFilename = "testcase/" + filename + ".recommend"
	recommend, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(recommend) == 0 {
		recommend = nil
	}

	// firstComments
	fullFilename = "testcase/" + filename + ".firstComments"
	firstComments, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(firstComments) == 0 {
		firstComments = nil
	}

	// theRestComments
	fullFilename = "testcase/" + filename + ".theRestComments"
	theRestComments, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return contentAll, content, signature, recommend, firstComments, theRestComments
}
