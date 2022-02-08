package dbcs

import (
	"io"
	"io/ioutil"
	"os"
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
}

func loadTest(filename string) (contentAll []byte, content []byte, signature []byte, recommend []byte, firstComments []byte, theRestComments []byte) {
	// content-all
	fullFilename := "testcase/" + filename
	file0, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file0.Close()

	r := io.Reader(file0)
	contentAll, _ = ioutil.ReadAll(r)

	// content
	fullFilename = "testcase/" + filename + ".content"
	file1, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file1.Close()

	r = io.Reader(file1)
	content, _ = ioutil.ReadAll(r)

	if len(content) == 0 {
		content = nil
	}

	// signature
	fullFilename = "testcase/" + filename + ".signature"
	file2, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file2.Close()

	r = io.Reader(file2)
	signature, _ = ioutil.ReadAll(r)

	if len(signature) == 0 {
		signature = nil
	}

	// recommend
	fullFilename = "testcase/" + filename + ".recommend"
	file3, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file3.Close()

	r = io.Reader(file3)
	recommend, _ = ioutil.ReadAll(r)

	if len(recommend) == 0 {
		recommend = nil
	}

	// firstComments
	fullFilename = "testcase/" + filename + ".firstComments"
	file4, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file4.Close()

	r = io.Reader(file4)
	firstComments, _ = ioutil.ReadAll(r)

	if len(firstComments) == 0 {
		firstComments = nil
	}

	// theRestComments
	fullFilename = "testcase/" + filename + ".theRestComments"
	file5, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file5.Close()

	r = io.Reader(file5)
	theRestComments, _ = ioutil.ReadAll(r)

	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return contentAll, content, signature, recommend, firstComments, theRestComments
}
