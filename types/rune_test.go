package types

import (
	"bytes"
	"testing"
)

func Test_initBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name                 string
		expectedBig5         []byte
		expectedUtf8         string
		expectedBig5FromUtf8 []byte
		wantErr              bool
	}{
		// TODO: Add test cases.
		{
			expectedBig5:         []byte{0x89, 0x71},
			expectedUtf8:         "辑",
			expectedBig5FromUtf8: []byte{0x89, 0x71},
		},
		{
			expectedBig5:         []byte{0xa4, 0xad},
			expectedUtf8:         "五",
			expectedBig5FromUtf8: []byte{0xa4, 0xad},
		},
		{
			expectedBig5:         []byte{0xa3, 0xb8},
			expectedUtf8:         "ㄧ",
			expectedBig5FromUtf8: []byte{0xa3, 0xb8},
		},
		{
			expectedBig5:         []byte{0xa4, 0x40},
			expectedUtf8:         "一",
			expectedBig5FromUtf8: []byte{0xa4, 0x40},
		},
		{
			expectedBig5:         []byte{0xb3, 0xfc},
			expectedUtf8:         "壹",
			expectedBig5FromUtf8: []byte{0xb3, 0xfc},
		},
		{ //not in golang
			expectedBig5:         []byte{0x85, 0x55},
			expectedUtf8:         "词",
			expectedBig5FromUtf8: []byte{0x85, 0x55},
		},
		{
			expectedBig5:         []byte{0x82, 0xab},
			expectedUtf8:         "直",
			expectedBig5FromUtf8: []byte{0xaa, 0xbd},
		},
		{
			expectedBig5:         []byte{0x9d, 0xde},
			expectedUtf8:         "♥",
			expectedBig5FromUtf8: []byte{0x9d, 0xde},
		},
		{ //not in golang
			expectedBig5:         []byte{0x99, 0xfc},
			expectedUtf8:         "㈲",
			expectedBig5FromUtf8: []byte{0x99, 0xfc},
		},
		{ //not in golang
			expectedBig5:         []byte{0x9d, 0x7b},
			expectedUtf8:         "弻",
			expectedBig5FromUtf8: []byte{0x9d, 0x7b},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initBig5(); (err != nil) != tt.wantErr {
				t.Errorf("initBig5() error = %v, wantErr %v", err, tt.wantErr)
			}

			utf8FromBig5 := big5ToUTF8[string(tt.expectedBig5)]
			if string(utf8FromBig5) != tt.expectedUtf8 {
				t.Errorf("initBig5 (fromBig5): utf8: %v expected: %v", string(utf8FromBig5), tt.expectedUtf8)
			}

			big5FromUtf8 := utf8ToBig5[tt.expectedUtf8]
			if !bytes.Equal(big5FromUtf8, tt.expectedBig5FromUtf8) {
				t.Errorf("initBig5 (fromUtf8): big5: %v expected: %v", big5FromUtf8, tt.expectedBig5FromUtf8)
			}
		})
	}
}

func TestRune_Utf8ToBig5(t *testing.T) {
	r1 := &Rune{
		Utf8: "今晚我想來點﹖然後呢？～1﹢1=？test所以勒～",
	}
	expected1 := []byte{0xa4, 0xb5, 0xb1, 0xdf, 0xa7, 0xda, 0xb7, 0x51, 0xa8, 0xd3, 0xc2, 0x49, 0xa1, 0x53, 0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3, 0x31, 0xa1, 0xde, 0x31, 0x3d, 0xa1, 0x48, 0x74, 0x65, 0x73, 0x74, 0xa9, 0xd2, 0xa5, 0x48, 0xb0, 0xc7, 0xa1, 0xe3}

	r2 := &Rune{
		Utf8: "這是ช ช้าง~聽起來很 cool～It says from google 翻譯 that it's an elephant.",
	}

	expected2 := []byte{
		0xb3, 0x6f, 0xac, 0x4f, 0xff, 0xfd, 0x20, 0xff, 0xfd, //這是ช ช้
		0xff, 0xfd, 0xff, 0xfd, 0xff, 0xfd, 0x7e, //าง~
		0xc5, 0xa5, 0xb0, 0x5f, 0xa8, 0xd3, 0xab, 0xdc, //聽起來很
		0x20, 0x63, 0x6f, 0x6f, 0x6c, 0xa1, 0xe3, 0x49, 0x74, // cool～It
		0x20, 0x73, 0x61, 0x79, 0x73, 0x20, 0x66, 0x72, 0x6f, // says fro
		0x6d, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, //m google
		0xc2, 0xbd, 0xc4, 0xb6, 0x20, 0x74, 0x68, 0x61, 0x74, //翻譯 that
		0x20, 0x69, 0x74, 0x27, 0x73, 0x20, 0x61, 0x6e, 0x20, //it's an
		0x65, 0x6c, 0x65, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x2e, //elephant.
	}

	r3 := &Rune{
		Utf8: "䶜",
	}

	expected3 := []byte{0xff, 0xfd}

	type fields struct {
		Utf8   string
		Big5   []byte
		Color0 Color
		Color1 Color
	}
	tests := []struct {
		name         string
		r            *Rune
		expectedBig5 []byte
	}{
		// TODO: Add test cases.
		{
			r:            r1,
			expectedBig5: expected1,
		},
		{
			r:            r2,
			expectedBig5: expected2,
		},
		{
			r:            r3,
			expectedBig5: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.Utf8ToBig5()

			if !bytes.Equal(r.Big5, tt.expectedBig5) {
				t.Errorf("Rune.Utf8ToBig5: Big5: %x expected: %x", r.Big5, tt.expectedBig5)
			}
		})
	}
}

func TestRune_Big5ToUtf8(t *testing.T) {

	r1 := &Rune{
		Big5: []byte{
			0xa4, 0xb5, 0xb1, 0xdf, 0xa7, 0xda, 0xb7, 0x51, //今晚我想
			0xa8, 0xd3, 0xc2, 0x49, 0xa1, 0x53, 0xb5, 0x4d, //來點﹖然
			0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3, //後呢？～
			0x31, 0xa1, 0xde, 0x31, 0x3d, 0xa1, 0x48, 0x74, //1﹢1=？t
			0x65, 0x73, 0x74, 0xa9, 0xd2, 0xa5, 0x48, 0xb0, //est所以勒
			0xc7, 0xa1, 0xe3, //勒～
		},
	}
	expected1 := "今晚我想來點﹖然後呢？～1﹢1=？test所以勒～"

	r2 := &Rune{
		Big5: []byte{
			0xb3, 0x6f, 0xac, 0x4f, 0xff, 0xfd, 0x20, 0xff, 0xfd, //這是ช ช้
			0xff, 0xfd, 0xff, 0xfd, 0xff, 0xfd, 0x7e, //าง~
			0xc5, 0xa5, 0xb0, 0x5f, 0xa8, 0xd3, 0xab, 0xdc, //聽起來很
			0x20, 0x63, 0x6f, 0x6f, 0x6c, 0xa1, 0xe3, 0x49, 0x74, // cool～It
			0x20, 0x73, 0x61, 0x79, 0x73, 0x20, 0x66, 0x72, 0x6f, // says fro
			0x6d, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, //m google
			0xc2, 0xbd, 0xc4, 0xb6, 0x20, 0x74, 0x68, 0x61, 0x74, //翻譯 that
			0x20, 0x69, 0x74, 0x27, 0x73, 0x20, 0x61, 0x6e, 0x20, //it's an
			0x65, 0x6c, 0x65, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x2e, //elephant.
		},
	}
	expected2 := "這是 ~聽起來很 cool～It says from google 翻譯 that it's an elephant."

	type fields struct {
		Utf8   string
		Big5   []byte
		Color0 Color
		Color1 Color
	}
	tests := []struct {
		name         string
		r            *Rune
		expectedUtf8 string
	}{
		// TODO: Add test cases.
		{
			r:            r1,
			expectedUtf8: expected1,
		},
		{
			r:            r2,
			expectedUtf8: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.Big5ToUtf8()

			if r.Utf8 != tt.expectedUtf8 {
				t.Errorf("Rune.Big5ToUtf8: Utf8: %v expected: %v", r.Utf8, expected1)
			}
		})
	}
}
