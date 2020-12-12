package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
)

func TestRegisterClient(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.Client_c.Drop()

	params := &RegisterClientParams{
		ClientID: "test_client_id",
	}

	expected := &RegisterClientResult{Success: true}

	expectedDB := []*schema.Client{
		{
			ClientID:     "test_client_id",
			ClientSecret: "test_client_secret",
			RemoteAddr:   "localhost",
		},
	}

	type args struct {
		remoteAddr string
		params     interface{}
	}
	tests := []struct {
		name       string
		args       args
		expected   interface{}
		expectedDB []*schema.Client
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
			got, _, err := RegisterClient(tt.args.remoteAddr, tt.args.params, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RegisterClient() = %v, want %v", got, tt.expected)
			}

			query := make(map[string]string)
			query["client_id"] = "test_client_id"

			var ret []*schema.Client
			err = schema.Client_c.Find(query, 0, &ret, nil)
			if err != nil {
				t.Errorf("RegisterClient(): unable to find: e: %v", err)
			}
			for _, each := range ret {
				each.UpdateNanoTS = 0
			}
			if !reflect.DeepEqual(ret, expectedDB) {
				t.Errorf("RegisterClient: find: %v expected: %v", ret, expectedDB)
			}
			for idx, each := range ret {
				if idx >= len(expectedDB) {
					t.Errorf("RegisterClient: (%v/%v) %v", idx, len(ret), each)
				}
				utils.TDeepEqual(t, each, expectedDB[idx])
			}
		})
	}
}
