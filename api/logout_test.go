package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLogout(t *testing.T) {
	setupTest()
	defer teardownTest()

	theContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	jsonBytes := []byte{}

	req, _ := http.NewRequest("POST", LOGOUT_R, bytes.NewBuffer(jsonBytes))

	theContext.Request = req

	logrus.Infof("theContext: %v", theContext)

	type args struct {
		remoteAddr string
		params     interface{}
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
			args: args{
				remoteAddr: "localhost",
				c:          theContext,
			},

			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			gotResult, gotStatusCode, err := Logout(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("Logout() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("Logout() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
