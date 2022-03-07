package types

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestColor_BytesWithPreColor(t *testing.T) {
	c0 := &Color{}
	*c0 = DefaultColor
	color0 := &Color{}
	*color0 = DefaultColor
	var expected0 []byte

	c1 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		Highlight:  true,
	}
	color1 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
	}
	expected1 := []byte("\x1b[1m")

	c2 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
	}
	color2 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Highlight:  true,
	}
	expected2 := []byte("\x1b[0;33;44m")

	c3 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		Blink:      true,
	}
	color3 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
	}
	expected3 := []byte("\x1b[5m")

	c4 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
	}
	color4 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Blink:      true,
	}
	expected4 := []byte("\x1b[0;33;44m")

	c5 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		IsReset:    true,
	}
	color5 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
	}
	expected5 := []byte("\x1b[m")

	c6 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
	}
	color6 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		IsReset:    true,
	}
	var expected6 []byte

	c7 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		Highlight:  true,
		Blink:      true,
	}
	color7 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
	}
	expected7 := []byte("\x1b[1;5m")

	c8 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		Highlight:  true,
	}
	color8 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		Blink:      true,
	}
	expected8 := []byte("\x1b[0;1m")

	c9 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Highlight:  true,
	}
	color9 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Blink:      true,
	}
	expected9 := []byte("\x1b[0;1;33;44m")

	c10 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Blink:      true,
	}
	color10 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Highlight:  true,
	}
	expected10 := []byte("\x1b[0;5;33;44m")

	c11 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
	}
	color11 := &Color{
		Foreground: COLOR_FOREGROUND_YELLOW,
		Background: COLOR_BACKGROUND_BLUE,
		Blink:      true,
		Highlight:  true,
	}
	expected11 := []byte("\x1b[0;33;44m")

	c12 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		IsReset:    true,
	}
	color12 := &Color{
		Foreground: COLOR_FOREGROUND_WHITE,
		Background: COLOR_BACKGROUND_BLACK,
		IsReset:    true,
	}
	var expected12 []byte

	type args struct {
		color *Color
	}
	tests := []struct {
		name             string
		c                *Color
		args             args
		expectedTheBytes []byte
	}{
		// TODO: Add test cases.
		{
			name:             "same color",
			c:                c0,
			args:             args{color: color0},
			expectedTheBytes: expected0,
		},
		{
			name:             "with highlight",
			c:                c1,
			args:             args{color: color1},
			expectedTheBytes: expected1,
		},
		{
			name:             "with no highlight",
			c:                c2,
			args:             args{color: color2},
			expectedTheBytes: expected2,
		},
		{
			name:             "with blink",
			c:                c3,
			args:             args{color: color3},
			expectedTheBytes: expected3,
		},
		{
			name:             "with no blink",
			c:                c4,
			args:             args{color: color4},
			expectedTheBytes: expected4,
		},
		{
			name:             "reset",
			c:                c5,
			args:             args{color: color5},
			expectedTheBytes: expected5,
		},
		{
			name:             "pre-reset",
			c:                c6,
			args:             args{color: color6},
			expectedTheBytes: expected6,
		},
		{
			name:             "with highlight/blink",
			c:                c7,
			args:             args{color: color7},
			expectedTheBytes: expected7,
		},
		{
			name:             "with highlight/no blink",
			c:                c8,
			args:             args{color: color8},
			expectedTheBytes: expected8,
		},
		{
			name:             "with highlight/no blink (color)",
			c:                c9,
			args:             args{color: color9},
			expectedTheBytes: expected9,
		},
		{
			name:             "with blink/no highlight (color)",
			c:                c10,
			args:             args{color: color10},
			expectedTheBytes: expected10,
		},
		{
			name:             "with no blink/no highlight (color)",
			c:                c11,
			args:             args{color: color11},
			expectedTheBytes: expected11,
		},
		{
			name:             "with both reset",
			c:                c12,
			args:             args{color: color12},
			expectedTheBytes: expected12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c
			gotTheBytes := c.BytesWithPreColor(tt.args.color)

			testutil.TDeepEqual(t, "got", gotTheBytes, tt.expectedTheBytes)
		})
	}
}
