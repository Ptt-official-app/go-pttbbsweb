package main

import (
	"context"
	"fmt"

	"github.com/Ptt-official-app/go-pttbbsweb/mand"
	"github.com/sirupsen/logrus"
)

func main() {
	brdname, path, err := initMain()
	if err != nil {
		logrus.Fatalf("unable to initMain: e: %v", err)
		return
	}

	ctx := context.Background()

	req := &mand.ListRequest{
		BoardName: brdname,
		Path:      path,
	}

	resp, err := mand.Cli.List(ctx, req)
	if err != nil {
		logrus.Errorf("unable to get content: e: %v", err)
		return
	}

	logrus.Infof("brdname: %v path: %v", brdname, path)

	fmt.Printf("=====begin=====\n")
	for idx, each := range resp.Entries {
		fmt.Printf("(%v/%v) %v", idx, len(resp.Entries), each)
	}
	fmt.Printf("====end=====\n")
}
