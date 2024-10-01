package queue

import (
	"os"
)

func initTest() {
	if testContentAll11 != nil {
		return
	}

	initTest11()
	initTestUtf87()
}

func loadTest(filename string) (contentAll []byte, content []byte, signature []byte, recommend []byte, firstComments []byte, theRestComments []byte) {
	// content-all
	fullFilename := "testcase/" + filename
	contentAll, err := os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	// content
	fullFilename = "testcase/" + filename + ".content"
	content, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(content) == 0 {
		content = nil
	}

	// signature
	fullFilename = "testcase/" + filename + ".signature"
	signature, err = os.ReadFile(fullFilename)
	if err != nil {
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
