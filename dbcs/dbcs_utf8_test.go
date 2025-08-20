package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_dbcsToUtf8PerLineIndexM(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		contentDBCS string
		idx         int
	}
	tests := []struct {
		name         string
		args         args
		expectedIdxM int
	}{
		// TODO: Add test cases.
		{
			args:         args{contentDBCS: "\x1b[", idx: 0},
			expectedIdxM: -1,
		},
		{
			args:         args{contentDBCS: "\x1b[m", idx: 0},
			expectedIdxM: 2,
		},
		{
			args:         args{contentDBCS: "abcdefg\x1b[m", idx: 7},
			expectedIdxM: 9,
		},
		{
			args:         args{contentDBCS: "方法Ａ；\x1b[1;33mCrtl+u \x1b[m(*)\x1b[1;33m [\x1b[m(在字母P右邊)\x1b[1;33mm^[[m", idx: len("方法Ａ；")},
			expectedIdxM: len("方法Ａ；") + 6,
		},
		{
			args:         args{contentDBCS: "方法Ａ；\x1b[1;33mCrtl+u \x1b[m(*)\x1b[1;33m [\x1b[m(在字母P右邊)\x1b[1;33mm^[[m", idx: len("方法Ａ；\x1b[1;33mCrtl+u ")},
			expectedIdxM: len("方法Ａ；\x1b[1;33mCrtl+u ") + 2,
		},
		{
			args:         args{contentDBCS: "方法Ａ；\x1b[1;33mCrtl+u \x1b[m(*)\x1b[1;33m [\x1b[m(在字母P右邊)\x1b[1;33mm^[[m", idx: len("方法Ａ；\x1b[1;33mCrtl+u \x1b[m(*)")},
			expectedIdxM: len("方法Ａ；\x1b[1;33mCrtl+u \x1b[m(*)") + 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIdxM := dbcsToUtf8PerLineIndexM(tt.args.contentDBCS, tt.args.idx); gotIdxM != tt.expectedIdxM {
				t.Errorf("dbcsToUtf8PerLineIndexM() = %v, want %v", gotIdxM, tt.expectedIdxM)
			}
		})
	}
}

func Test_dbcsToUtf8PerLineNextUtf8(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		contentDBCS string
	}
	tests := []struct {
		name            string
		args            args
		expectedTheUtf8 string
	}{
		// TODO: Add test cases.
		{
			args: args{contentDBCS: "\x1b[m"},
		},
		{
			args: args{contentDBCS: "\x1b[1;33;145m"},
		},
		{
			args:            args{contentDBCS: "我\x1b[1;33;145m"},
			expectedTheUtf8: "我",
		},
		{
			args:            args{contentDBCS: " \x1b[1;33;145m"},
			expectedTheUtf8: " ",
		},
		{
			args:            args{contentDBCS: ""},
			expectedTheUtf8: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTheUtf8 := dbcsToUtf8PerLineNextUtf8(tt.args.contentDBCS); gotTheUtf8 != tt.expectedTheUtf8 {
				t.Errorf("dbcsToUtf8PerLineNextUtf8() = %v, want %v", gotTheUtf8, tt.expectedTheUtf8)
			}
		})
	}
}

func Test_dbcsToUtf8PerLineParseColor(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		colorDBCS string
		origColor types.Color
	}
	tests := []struct {
		name             string
		args             args
		expectedColor    types.Color
		expectedIsColor1 bool
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name:          "typical color setting",
			args:          args{colorDBCS: "\x1b[1;33;40m", origColor: types.DefaultColor},
			expectedColor: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
		},
		{
			name:          "with \x1b[0m",
			args:          args{colorDBCS: "\x1b[0;31;47m", origColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true, Blink: true}},
			expectedColor: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE},
		},
		{
			name:          "with \x1b[;m",
			args:          args{colorDBCS: "\x1b[;31;47m", origColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true, Blink: true}},
			expectedColor: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE},
		},
		{
			name:          "with \x1b[m",
			args:          args{colorDBCS: "\x1b[m", origColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true, Blink: true}},
			expectedColor: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
		},
		{
			args:             args{colorDBCS: "\x1b[111;132;146m", origColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: false, Blink: false}},
			expectedColor:    types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
			expectedIsColor1: true,
		},
		{
			name:             "110;105;103",
			args:             args{colorDBCS: "\x1b[110;105;103m", origColor: types.DefaultColor},
			expectedColor:    types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_YELLOW},
			expectedIsColor1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotColor, gotIsColor1, err := dbcsToUtf8PerLineParseColor(tt.args.colorDBCS, tt.args.origColor)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbcsToUtf8PerLineParseColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotIsColor1 != tt.expectedIsColor1 {
				t.Errorf("dbcsToUtf8PerLineParseColor: got: %v want: %v", gotIsColor1, tt.expectedIsColor1)
			}

			if !reflect.DeepEqual(gotColor, tt.expectedColor) {
				t.Errorf("dbcsToUtf8PerLineParseColor() = %v, want %v", gotColor, tt.expectedColor)
			}
		})
	}
}

func Test_dbcsToUtf8PerLine(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		contentDBCS string
		color0      types.Color
	}
	tests := []struct {
		name             string
		args             args
		expectedLine     []*types.Rune
		expectedNewColor types.Color
	}{
		// TODO: Add test cases.
		{
			name: "double color",
			args: args{
				color0:      types.DefaultColor,
				contentDBCS: "\x1b[1;30m再來是自以為很高級的\x1b[37m\x1b[130m一\x1b[37m\x1b[130m字\x1b[37m\x1b[130m雙\x1b[37m\x1b[130m色\x1b[m",
			},
			expectedLine: []*types.Rune{
				{
					Utf8:    "再來是自以為很高級的",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					DBCSStr: "\x1b[1;30m再來是自以為很高級的",
				},
				{
					Utf8:    "一",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[130m一",
				},
				{
					Utf8:    "字",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[130m字",
				},
				{
					Utf8:    "雙",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[130m雙",
				},
				{
					Utf8:    "色",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[130m色",
				},
				{
					Utf8:    "",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
					DBCSStr: "\x1b[m",
				},
			},
			expectedNewColor: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
		},

		{
			name: "continuous color-code",
			args: args{
				color0:      types.DefaultColor,
				contentDBCS: "\x1b[1;30m\x1b[31;42m再來是自以為很高級的\x1b[37m\x1b[31;42m\x1b[43m\x1b[130m一\x1b[37m\x1b[34;43m字\x1b[45m",
			},
			expectedLine: []*types.Rune{
				{
					Utf8:    "再來是自以為很高級的",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					DBCSStr: "\x1b[1;30m\x1b[31;42m再來是自以為很高級的",
				},
				{
					Utf8:    "一",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[31;42m\x1b[43m\x1b[130m一",
				},
				{
					Utf8:    "字",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[34;43m字",
				},
				{
					Utf8:    "",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
					DBCSStr: "\x1b[45m",
				},
			},
			expectedNewColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
		},
		{
			name: "continuous color-code",
			args: args{
				color0:      types.DefaultColor,
				contentDBCS: "\x1b[1;30m\x1b[31;42m\x1b[134m再來是自以為很高級的\x1b[37m\x1b[31;42m\x1b[43m\x1b[130m一\x1b[37m\x1b[34;43m字\x1b[45m",
			},
			expectedLine: []*types.Rune{
				{
					Utf8:    "再",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					DBCSStr: "\x1b[1;30m\x1b[31;42m\x1b[134m再",
				},
				{
					Utf8:    "來是自以為很高級的",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
					DBCSStr: "來是自以為很高級的",
				},
				{
					Utf8:    "一",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[31;42m\x1b[43m\x1b[130m一",
				},
				{
					Utf8:    "字",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
					DBCSStr: "\x1b[37m\x1b[34;43m字",
				},
				{
					Utf8:    "",
					Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
					Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
					DBCSStr: "\x1b[45m",
				},
			},
			expectedNewColor: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
		},
		{
			args: args{
				color0:      types.DefaultColor,
				contentDBCS: "※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1",
			},
			expectedLine: []*types.Rune{
				{
					Utf8:    "※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1",
					Color0:  types.DefaultColor,
					Color1:  types.DefaultColor,
					DBCSStr: "※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1",
				},
			},
			expectedNewColor: types.DefaultColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLine, gotNewColor := dbcsToUtf8PerLine(tt.args.contentDBCS, tt.args.color0)
			testutil.TDeepEqual(t, "line", gotLine, tt.expectedLine)
			if !reflect.DeepEqual(gotNewColor, tt.expectedNewColor) {
				t.Errorf("dbcsToUtf8PerLine() gotNewColor = %v, want %v", gotNewColor, tt.expectedNewColor)
			}
		})
	}
}
