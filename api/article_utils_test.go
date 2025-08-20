package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_simplifyEachRunes(t *testing.T) {
	setupTest()
	defer teardownTest()

	runes0 := []*types.Rune{
		{Utf8: "這是測試0", Color0: types.DefaultColor, Color1: types.DefaultColor},
		{Utf8: "這是測試1", Color0: types.DefaultColor, Color1: types.DefaultColor},
		{Utf8: "這是測試2", Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK}, Color1: types.DefaultColor},
	}

	ret0 := []*types.Rune{
		{Utf8: "這是測試0這是測試1", Color0: types.DefaultColor, Color1: types.DefaultColor},
		{Utf8: "這是測試2", Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK}, Color1: types.DefaultColor},
	}

	type args struct {
		runes []*types.Rune
	}
	tests := []struct {
		name    string
		args    args
		wantRet []*types.Rune
	}{
		// TODO: Add test cases.
		{
			args:    args{runes: runes0},
			wantRet: ret0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := simplifyEachRunes(tt.args.runes); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("simplifyEachRunes() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
