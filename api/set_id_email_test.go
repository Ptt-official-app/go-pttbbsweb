package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	pttbbsai "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func TestSetIDEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserIDEmail_c.Drop()

	jwt0, _ := pttbbsai.CreateEmailToken("SYSOP", "app", "test@ptt.test", pttbbsai.CONTEXT_SET_ID_EMAIL)

	params0 := &SetIDEmailParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Jwt:          jwt0,
	}

	path0 := &SetIDEmailPath{
		UserID: "SYSOP",
	}

	result0 := &SetIDEmailResult{
		Email: "test@ptt.test",
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
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.

		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     result0,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := SetIDEmail(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetIDEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("SetIDEmail() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("SetIDEmail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
