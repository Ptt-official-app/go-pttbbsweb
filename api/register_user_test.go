package api

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/gin-gonic/gin"
)

func TestRegisterUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.AccessToken_c.Drop()

	params0 := &RegisterUserParams{
		ClientID:        "default_client_id",
		ClientSecret:    "test_client_secret",
		Username:        "testuserid1",
		Password:        "testpasswd",
		PasswordConfirm: "testpasswd",
		Email:           "test@ptt.test",
		TwoFactorToken:  "123123",
	}

	expected0 := &RegisterUserResult{TokenType: "bearer", UserID: "testuserid1", TokenUser: "testuserid1"}
	expectedDB0 := []*schema.AccessToken{{UserID: "testuserid1"}}

	_ = schema.Set2FA("testuserid1", "test@ptt.test", "123123", time.Duration(1)*time.Second)

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     *RegisterUserResult
		expectedStatusCode int
		expectedDB         []*schema.AccessToken
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: "localhost", params: params0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
			expectedDB:         expectedDB0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := RegisterUser(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return

			}

			/*
				query := make(map[string]string)
				query["user_id"] = "testuserid1"

				var ret []*schema.AccessToken
				err = schema.AccessToken_c.Find(query, 0, &ret, nil, nil)
				if err != nil {
					t.Errorf("Login(): unable to find: e: %v", err)
				}
				for _, each := range ret {
					each.UpdateNanoTS = 0
				}
				if len(ret) < 1 {
					t.Errorf("RegisterUser(): unable to find access-token.")
					return
				}
				expected.AccessToken = ret[0].AccessToken
			*/

			result := gotResult.(*RegisterUserResult)
			tt.expectedResult.AccessToken = result.AccessToken

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RegisterUser() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("RegisterUser() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
