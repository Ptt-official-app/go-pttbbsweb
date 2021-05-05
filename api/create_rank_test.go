package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRank(t *testing.T) {
	setupTest()
	defer teardownTest()

	schema.UpdateArticleSummaries([]*schema.ArticleSummary{{BBoardID: "10_WhoAmI", ArticleID: "test_articleID"}}, 1234567890000000000)

	path0 := &CreateRankPath{
		BoardID:   "10_WhoAmI",
		ArticleID: "test_articleID",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     *CreateRankResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 1},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: &CreateRankParams{Rank: -1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: -1},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP1", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 0},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP2", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 1},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP3", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 2},
			expectedStatusCode: 200,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := CreateRank(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("CreateRank() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateRank() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}

			path0 := &GetArticleDetailPath{BBoardID: "10_WhoAmI", ArticleID: "test_articleID"}
			r0, _, _ := GetArticleDetail(tt.args.remoteAddr, tt.args.userID, nil, path0, nil)
			if r0 != nil {
				ret0 := r0.(*GetArticleDetailResult)
				assert.Equal(t, tt.expectedResult.Rank, ret0.Rank)
			}

		})
		wg.Wait()
	}
}
