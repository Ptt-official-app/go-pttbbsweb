package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.AccessToken_c.Drop()

	params := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Username:     "testuserid1",
		Password:     "testpasswd",
	}

	expected := &LoginResult{TokenType: "bearer"}
	expectedDB := []*schema.AccessToken{{UserID: "testuserid1"}}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name       string
		args       args
		expected   interface{}
		expectedDB []*schema.AccessToken
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			args:       args{remoteAddr: "localhost", params: params},
			expected:   expected,
			expectedDB: expectedDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := Login(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			query := make(map[string]string)
			query["user_id"] = "testuserid1"

			var ret []*schema.AccessToken
			err = schema.AccessToken_c.Find(query, 0, &ret, nil)
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

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Login() = %v, want %v", got, tt.expected)
			}
		})
	}
}
