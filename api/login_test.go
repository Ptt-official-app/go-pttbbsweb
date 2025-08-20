package api

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLogin(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.AccessToken_c.Drop()

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Username:     "testuserid1",
		Password:     "testpasswd",
	}

	expected0 := &LoginResult{TokenType: "bearer", UserID: "testuserid1", TokenUser: "testuserid1"}
	expectedDB0 := []*schema.AccessToken{{UserID: "testuserid1"}}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name       string
		args       args
		expected   *LoginResult
		expectedDB []*schema.AccessToken
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			args:       args{remoteAddr: "localhost", params: params0},
			expected:   expected0,
			expectedDB: expectedDB0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := Login(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			/*
				query := make(map[string]interface{})
				query[schema.ACCESS_TOKEN_USER_ID_b] = "testuserid1"

				var ret []*schema.AccessToken
				err = schema.AccessToken_c.Find(query, 0, &ret, nil, nil)
				logrus.Infof("api.TestLogin: after Find: query: %v ret: %v e: %v", query, ret, err)
				if err != nil {
					t.Errorf("Login(): unable to find: e: %v", err)
				}
				for _, each := range ret {
					each.UpdateNanoTS = 0
				}
				if len(ret) < 1 {
					t.Errorf("Login(): unable to find access-token")
					return
				}
				expected.AccessToken = ret[0].AccessToken
			*/
			result := got.(*LoginResult)
			tt.expected.AccessToken = result.AccessToken

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Login() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLoginWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Username:     "SYSOP",
		Password:     "123123",
	}
	type args struct {
		params *LoginParams
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{params: params0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c, r := testSetRequest(
				LOGIN_R,
				LOGIN_R,
				tt.args.params,
				"",
				"",
				nil,
				"POST",
				LoginWrapper,
			)

			r.ServeHTTP(w, c.Request)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			setCookie := w.Header().Get("Set-Cookie")
			logrus.Infof("setCookie: %v", setCookie)
		})
	}
}
