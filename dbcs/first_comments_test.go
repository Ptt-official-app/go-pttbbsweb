package dbcs

import (
	"reflect"
	"testing"
)

func Test_splitFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		commentsDBCS []byte
	}
	tests := []struct {
		name                      string
		args                      args
		expectedFirstCommentsDBCS []byte
		expectedTheRestComments   []byte
	}{
		// TODO: Add test cases.
		{
			name:                      "0_" + testFilename0,
			args:                      args{commentsDBCS: testComment0},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS0,
			expectedTheRestComments:   testTheRestCommentsDBCS0,
		},
		{
			name:                      "1_" + testFilename1,
			args:                      args{commentsDBCS: testComment1},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS1,
			expectedTheRestComments:   testTheRestCommentsDBCS1,
		},
		{
			name:                      "2_" + testFilename2,
			args:                      args{commentsDBCS: testComment2},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS2,
			expectedTheRestComments:   testTheRestCommentsDBCS2,
		},
		{
			name:                      "3_" + testFilename3,
			args:                      args{commentsDBCS: testComment3},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS3,
			expectedTheRestComments:   testTheRestCommentsDBCS3,
		},
		{
			name:                      "4_" + testFilename4,
			args:                      args{commentsDBCS: testComment4},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS4,
			expectedTheRestComments:   testTheRestCommentsDBCS4,
		},
		{
			name:                      "5_" + testFilename5,
			args:                      args{commentsDBCS: testComment5},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS5,
			expectedTheRestComments:   testTheRestCommentsDBCS5,
		},
		{
			name:                      "6_" + testFilename6,
			args:                      args{commentsDBCS: testComment6},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS6,
			expectedTheRestComments:   testTheRestCommentsDBCS6,
		},
		{
			name:                      "7_" + testFilename7,
			args:                      args{commentsDBCS: testComment7},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS7,
			expectedTheRestComments:   testTheRestCommentsDBCS7,
		},
		{
			name:                      "8_" + testFilename8,
			args:                      args{commentsDBCS: testComment8},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS8,
			expectedTheRestComments:   testTheRestCommentsDBCS8,
		},
		{
			name:                      "9_" + testFilename9,
			args:                      args{commentsDBCS: testComment9},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS9,
			expectedTheRestComments:   testTheRestCommentsDBCS9,
		},
		{
			name:                      "10_" + testFilename10,
			args:                      args{commentsDBCS: testComment10},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS10,
			expectedTheRestComments:   testTheRestCommentsDBCS10,
		},
		{
			name:                      "11_" + testFilename11,
			args:                      args{commentsDBCS: testComment11},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS11,
			expectedTheRestComments:   testTheRestCommentsDBCS11,
		},
		{
			name:                      "12_" + testFilename12,
			args:                      args{commentsDBCS: testComment12},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS12,
			expectedTheRestComments:   testTheRestCommentsDBCS12,
		},
		{
			name:                      "13_" + testFilename13,
			args:                      args{commentsDBCS: testComment13},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS13,
			expectedTheRestComments:   testTheRestCommentsDBCS13,
		},
		{
			name:                      "14_" + testFilename14,
			args:                      args{commentsDBCS: testComment14},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS14,
			expectedTheRestComments:   testTheRestCommentsDBCS14,
		},
		{
			name:                      "15_" + testFilename15,
			args:                      args{commentsDBCS: testComment15},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS15,
			expectedTheRestComments:   testTheRestCommentsDBCS15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstCommentsDBCS, gotTheRestComments := splitFirstComments(tt.args.commentsDBCS)
			if !reflect.DeepEqual(gotFirstCommentsDBCS, tt.expectedFirstCommentsDBCS) {
				t.Errorf("splitFirstComments() gotFirstCommentsDBCS = %v, want %v", gotFirstCommentsDBCS, tt.expectedFirstCommentsDBCS)
			}
			if !reflect.DeepEqual(gotTheRestComments, tt.expectedTheRestComments) {
				t.Errorf("splitFirstComments() gotTheRestComments = %v, want %v", gotTheRestComments, tt.expectedTheRestComments)
			}
		})
	}
}
