package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func TestLoadGeneralArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserReadBoard_c.Drop()

	update0 := &schema.UserReadArticle{UserID: "SYSOP", ArticleID: "1_123124", UpdateNanoTS: types.Time8(1534567891).ToNanoTS()}
	update1 := &schema.UserReadArticle{UserID: "SYSOP", ArticleID: "2_123125", UpdateNanoTS: types.Time8(1234567800).ToNanoTS()}

	_, _ = schema.UserReadArticle_c.Update(update0, update0)
	_, _ = schema.UserReadArticle_c.Update(update1, update1)

	params := &LoadGeneralArticlesParams{}
	path := &LoadGeneralArticlesPath{BBoardID: "1_test1"}
	expectedResult := &LoadGeneralArticlesResult{
		List: []*types.ArticleSummary{
			{
				BBoardID:   bbs.BBoardID("1_test1"),
				ArticleID:  bbs.ArticleID("1_123124"),
				IsDeleted:  false,
				Filename:   "M.1234567890.A.324",
				CreateTime: types.Time8(1234567890),
				MTime:      types.Time8(1234567889),
				Recommend:  8,
				Owner:      "okcool",
				Date:       "12/04",
				Title:      "[問題]然後呢？～",
				Money:      3,
				Filemode:   0,
				URL:        "http://localhost/bbs/test1/M.1234567890.A.324.html",
				Read:       true,
			},
			{
				BBoardID:   bbs.BBoardID("1_test1"),
				ArticleID:  bbs.ArticleID("2_123125"),
				IsDeleted:  false,
				Filename:   "M.1234567890.A.325",
				CreateTime: types.Time8(1234567900),
				MTime:      types.Time8(1234567890),
				Recommend:  3,
				Owner:      "teemo",
				Date:       "12/05",
				Title:      "[問題]再來呢？～",
				Money:      12,
				Filemode:   0,
				URL:        "http://localhost/bbs/test1/M.1234567890.A.325.html",
				Read:       false,
			},
		},
		NextIdx: "textNextIdx",
	}

	c := &gin.Context{}
	c.Request = &http.Request{URL: &url.URL{Path: "/api/board/1_test/articles"}}
	type args struct {
		remoteAddr string
		userID     string
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params, path: path, c: &gin.Context{}},
			expectedResult:     expectedResult,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := LoadGeneralArticles(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotResultVal := gotResult.(*LoadGeneralArticlesResult)
			expectedResultVal := tt.expectedResult.(*LoadGeneralArticlesResult)
			for idx, each := range gotResultVal.List {
				if idx >= len(expectedResultVal.List) {
					t.Errorf("LoadGeneralArticles() (%v/%v): %v", idx, len(gotResultVal.List), each)

				}

				log.Infof("LoadGeneralArticles() (%v/%v): (%v/%v) expected: (%v/%v)", idx, len(gotResultVal.List), each, reflect.TypeOf(each), expectedResultVal.List[idx], reflect.TypeOf(expectedResultVal.List[idx]))

				utils.TDeepEqual(t, each, expectedResultVal.List[idx])
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadGeneralArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
