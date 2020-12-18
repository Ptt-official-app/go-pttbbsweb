package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func Test_dbcsToBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		dbcs []byte
	}
	tests := []struct {
		name         string
		args         args
		expectedBig5 [][]*types.Rune
	}{
		// TODO: Add test cases.
		{
			name:         "0_" + testFilename0 + "_content",
			args:         args{dbcs: testContent0},
			expectedBig5: testContent0Big5,
		},
		{
			name:         "1_" + testFilename1 + "_content",
			args:         args{dbcs: testContent1},
			expectedBig5: testContent1Big5,
		},
		{
			name:         "2_" + testFilename2 + "_content",
			args:         args{dbcs: testContent2},
			expectedBig5: testContent2Big5,
		},
		{
			name:         "3_" + testFilename3 + "_content",
			args:         args{dbcs: testContent3},
			expectedBig5: testContent3Big5,
		},
		{
			name:         "4_" + testFilename4 + "_content",
			args:         args{dbcs: testContent4},
			expectedBig5: testContent4Big5,
		},
		{
			name:         "5_" + testFilename5 + "_content",
			args:         args{dbcs: testContent5},
			expectedBig5: testContent5Big5,
		},
		{
			name:         "6_" + testFilename6 + "_content",
			args:         args{dbcs: testContent6},
			expectedBig5: testContent6Big5,
		},
		{
			name:         "7_" + testFilename7 + "_content",
			args:         args{dbcs: testContent7},
			expectedBig5: testContent7Big5,
		},
		{
			name:         "8_" + testFilename8 + "_content",
			args:         args{dbcs: testContent8},
			expectedBig5: testContent8Big5,
		},
		{
			name:         "9_" + testFilename9 + "_content",
			args:         args{dbcs: testContent9},
			expectedBig5: testContent9Big5,
		},
		{
			name:         "10_" + testFilename10 + "_content",
			args:         args{dbcs: testContent10},
			expectedBig5: testContent10Big5,
		},
		{
			name:         "11_" + testFilename11 + "_content",
			args:         args{dbcs: testContent11},
			expectedBig5: testContent11Big5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBig5 := dbcsToBig5(tt.args.dbcs)

			if len(gotBig5) != len(tt.expectedBig5) {
				t.Errorf("len(Big5): %v expected: %v", len(gotBig5), len(tt.expectedBig5))
			}

			for idx, each := range gotBig5 {
				if idx >= len(tt.expectedBig5) {
					t.Errorf("Big5: (%v/%v) %v", idx, len(gotBig5), each)
					continue
				}

				if len(each) != len(tt.expectedBig5[idx]) {
					t.Errorf("Big5: (%v/%v): len: %v expected: %v", idx, len(gotBig5), len(each), len(tt.expectedBig5[idx]))
				}
				for idx2, each2 := range each {
					if idx2 >= len(tt.expectedBig5[idx]) {
						t.Errorf("Big5: (%v/%v/%v/%v): %v", idx, len(gotBig5), idx2, len(each), each2)
						continue
					}
					if !reflect.DeepEqual(each2, tt.expectedBig5[idx][idx2]) {
						t.Errorf("Big5: (%v/%v/%v/%v) %v expected: %v", idx, len(gotBig5), idx2, len(each), each2, tt.expectedBig5[idx][idx2])
					}
				}
			}
		})
	}
}

func Test_big5ToUtf8(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		a [][]*types.Rune
	}
	tests := []struct {
		name     string
		args     args
		expected [][]*types.Rune
	}{
		// TODO: Add test cases.
		{
			name:     "0_" + testFilename0 + "_content",
			args:     args{a: testContent0Big5},
			expected: testContent0Utf8,
		},
		{
			name:     "1_" + testFilename1 + "_content",
			args:     args{a: testContent1Big5},
			expected: testContent1Utf8,
		},
		{
			name:     "2_" + testFilename2 + "_content",
			args:     args{a: testContent2Big5},
			expected: testContent2Utf8,
		},
		{
			name:     "3_" + testFilename3 + "_content",
			args:     args{a: testContent3Big5},
			expected: testContent3Utf8,
		},
		{
			name:     "4_" + testFilename4 + "_content",
			args:     args{a: testContent4Big5},
			expected: testContent4Utf8,
		},
		{
			name:     "5_" + testFilename5 + "_content",
			args:     args{a: testContent5Big5},
			expected: testContent5Utf8,
		},
		{
			name:     "6_" + testFilename6 + "_content",
			args:     args{a: testContent6Big5},
			expected: testContent6Utf8,
		},
		{
			name:     "7_" + testFilename7 + "_content",
			args:     args{a: testContent7Big5},
			expected: testContent7Utf8,
		},
		{
			name:     "8_" + testFilename8 + "_content",
			args:     args{a: testContent8Big5},
			expected: testContent8Utf8,
		},
		{
			name:     "9_" + testFilename9 + "_content",
			args:     args{a: testContent9Big5},
			expected: testContent9Utf8,
		},
		{
			name:     "10_" + testFilename10 + "_content",
			args:     args{a: testContent10Big5},
			expected: testContent10Utf8,
		},
		{
			name:     "11_" + testFilename11 + "_content",
			args:     args{a: testContent11Big5},
			expected: testContent11Utf8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := big5ToUtf8(tt.args.a)

			if len(got) != len(tt.expected) {
				t.Errorf("len(utf8): %v expected: %v", len(got), len(tt.expected))
			}

			for idx, each := range got {
				if idx >= len(tt.expected) {
					t.Errorf("utf8: (%v/%v) %v", idx, len(got), each)
					continue
				}

				if len(each) != len(tt.expected[idx]) {
					t.Errorf("utf8: (%v/%v): len: %v expected: %v", idx, len(got), len(each), len(tt.expected[idx]))
				}
				for idx2, each2 := range each {
					if idx2 >= len(tt.expected[idx]) {
						t.Errorf("utf8: (%v/%v/%v/%v): %v", idx, len(got), idx2, len(each), each2)
						continue
					}
					if !reflect.DeepEqual(each2, tt.expected[idx][idx2]) {
						t.Errorf("utf8: (%v/%v/%v/%v) %v expected: %v", idx, len(got), idx2, len(each), each2, tt.expected[idx][idx2])
					}
				}
			}

		})
	}
}
