package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/boardd"
	"github.com/sirupsen/logrus"
)

func main() {
	filename, err := initMain()
	if err != nil {
		logrus.Fatalf("unable to initMain: e: %v", err)
		return
	}

	filenameList := strings.Split(filename, "/")

	ctx := context.Background()

	brdname := &boardd.BoardRef_Name{Name: filenameList[0]}
	brdref := &boardd.BoardRef{Ref: brdname}
	req := &boardd.ContentRequest{
		BoardRef: brdref,
		Filename: filename,
		PartialOptions: &boardd.PartialOptions{
			SelectType: boardd.PartialOptions_SELECT_FULL,
		},
	}

	resp, err := boardd.Cli.Content(ctx, req)
	if err != nil {
		logrus.Errorf("unable to get content: e: %v", err)
		return
	}

	logrus.Infof("filename: %v consistency-token: %v offset: %v length: %v total-length: %v", filename, resp.Content.ConsistencyToken, resp.Content.Offset, resp.Content.Length, resp.Content.TotalLength)

	fmt.Printf("=====begin=====\n")
	fmt.Printf("%v", string(resp.Content.Content))
	fmt.Printf("====end=====\n")
}
