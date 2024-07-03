package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestCalculateAllUserVisitCounts(t *testing.T) {
	setupTest()
	defer teardownTest()
	defer UserVisit_c.Drop()

	userVisit := &UserVisit{UserID: "", Action: "test", UpdateNanoTS: types.NowNanoTS()}
	_ = UpdateUserVisit(userVisit)

	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test get 1 user visit count",
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateAllUserVisitCounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAllUserVisitCounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateAllUserVisitCounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
