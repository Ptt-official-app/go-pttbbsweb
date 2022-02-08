package dbcs

import (
	"reflect"
	"testing"
)

func Test_splitArticleSignatureCommentsDBCS(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		content []byte
	}
	tests := []struct {
		name                  string
		args                  args
		expectedArticleDBCS   []byte
		expectedSignatureDBCS []byte
		expectedComments      []byte
	}{
		// TODO: Add test cases.
		{
			name:                  "0_" + testFilename0,
			args:                  args{content: testContentAll0},
			expectedArticleDBCS:   testContent0,
			expectedSignatureDBCS: testSignature0,
			expectedComments:      testComment0,
		},
		{
			name:                  "1_" + testFilename1,
			args:                  args{content: testContentAll1},
			expectedArticleDBCS:   testContent1,
			expectedSignatureDBCS: testSignature1,
			expectedComments:      testComment1,
		},
		{
			name:                  "2_" + testFilename2,
			args:                  args{content: testContentAll2},
			expectedArticleDBCS:   testContent2,
			expectedSignatureDBCS: testSignature2,
			expectedComments:      testComment2,
		},
		{
			name:                  "3_" + testFilename3,
			args:                  args{content: testContentAll3},
			expectedArticleDBCS:   testContent3,
			expectedSignatureDBCS: testSignature3,
			expectedComments:      testComment3,
		},
		{
			name:                  "4_" + testFilename4,
			args:                  args{content: testContentAll4},
			expectedArticleDBCS:   testContent4,
			expectedSignatureDBCS: testSignature4,
			expectedComments:      testComment4,
		},
		{
			name:                  "5_" + testFilename5,
			args:                  args{content: testContentAll5},
			expectedArticleDBCS:   testContent5,
			expectedSignatureDBCS: testSignature5,
			expectedComments:      testComment5,
		},
		{
			name:                  "6_" + testFilename6,
			args:                  args{content: testContentAll6},
			expectedArticleDBCS:   testContent6,
			expectedSignatureDBCS: testSignature6,
			expectedComments:      testComment6,
		},
		{
			name:                  "7_" + testFilename7,
			args:                  args{content: testContentAll7},
			expectedArticleDBCS:   testContent7,
			expectedSignatureDBCS: testSignature7,
			expectedComments:      testComment7,
		},
		{
			name:                  "8_" + testFilename8,
			args:                  args{content: testContentAll8},
			expectedArticleDBCS:   testContent8,
			expectedSignatureDBCS: testSignature8,
			expectedComments:      testComment8,
		},
		{
			name:                  "9_" + testFilename9,
			args:                  args{content: testContentAll9},
			expectedArticleDBCS:   testContent9,
			expectedSignatureDBCS: testSignature9,
			expectedComments:      testComment9,
		},
		{
			name:                  "10_" + testFilename10,
			args:                  args{content: testContentAll10},
			expectedArticleDBCS:   testContent10,
			expectedSignatureDBCS: testSignature10,
			expectedComments:      testComment10,
		},
		{
			name:                  "11_" + testFilename11,
			args:                  args{content: testContentAll11},
			expectedArticleDBCS:   testContent11,
			expectedSignatureDBCS: testSignature11,
			expectedComments:      testComment11,
		},
		{
			name:                  "12_" + testFilename12,
			args:                  args{content: testContentAll12},
			expectedArticleDBCS:   testContent12,
			expectedSignatureDBCS: testSignature12,
			expectedComments:      testComment12,
		},
		{
			name:                  "13_" + testFilename13,
			args:                  args{content: testContentAll13},
			expectedArticleDBCS:   testContent13,
			expectedSignatureDBCS: testSignature13,
			expectedComments:      testComment13,
		},
		{
			name:                  "14_" + testFilename14,
			args:                  args{content: testContentAll14},
			expectedArticleDBCS:   testContent14,
			expectedSignatureDBCS: testSignature14,
			expectedComments:      testComment14,
		},
		{
			name:                  "15_" + testFilename15,
			args:                  args{content: testContentAll15},
			expectedArticleDBCS:   testContent15,
			expectedSignatureDBCS: testSignature15,
			expectedComments:      testComment15,
		},
		{
			name:                  "16_" + testFilename16,
			args:                  args{content: testContentAll16},
			expectedArticleDBCS:   testContent16,
			expectedSignatureDBCS: testSignature16,
			expectedComments:      testComment16,
		},
		{
			name:                  "17_" + testFilename17,
			args:                  args{content: testContentAll17},
			expectedArticleDBCS:   testContent17,
			expectedSignatureDBCS: testSignature17,
			expectedComments:      testComment17,
		},
		{
			name:                  "18_" + testFilename18,
			args:                  args{content: testContentAll18},
			expectedArticleDBCS:   testContent18,
			expectedSignatureDBCS: testSignature18,
			expectedComments:      testComment18,
		},
		{
			name:                  "19_" + testFilename19,
			args:                  args{content: testContentAll19},
			expectedArticleDBCS:   testContent19,
			expectedSignatureDBCS: testSignature19,
			expectedComments:      testComment19,
		},
		{
			name:                  "20_" + testFilename20,
			args:                  args{content: testContentAll20},
			expectedArticleDBCS:   testContent20,
			expectedSignatureDBCS: testSignature20,
			expectedComments:      testComment20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArticleDBCS, gotSignatureDBCS, gotComments := splitArticleSignatureCommentsDBCS(tt.args.content)
			if !reflect.DeepEqual(gotArticleDBCS, tt.expectedArticleDBCS) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotArticleDBCS = %v, want %v", gotArticleDBCS, tt.expectedArticleDBCS)
			}
			if !reflect.DeepEqual(gotSignatureDBCS, tt.expectedSignatureDBCS) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotSignatureDBCS = %v, want %v", gotSignatureDBCS, tt.expectedSignatureDBCS)
			}
			if !reflect.DeepEqual(gotComments, tt.expectedComments) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotComments = %v, want %v", gotComments, tt.expectedComments)
			}
		})
	}
}

func Test_tryGetSimpleSignatureIdxes(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		content []byte
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		// TODO: Add test cases.
		{
			name:     "0_" + testFilename0,
			args:     args{content: testContentAll0},
			expected: []int{133},
		},
		{
			name:     "1_" + testFilename1,
			args:     args{content: testContentAll1},
			expected: []int{96},
		},
		{
			name:     "2_" + testFilename2,
			args:     args{content: testContentAll2},
			expected: []int{260},
		},
		{
			name:     "3_" + testFilename3,
			args:     args{content: testContentAll3},
			expected: []int{121},
		},
		{
			name:     "4_" + testFilename4,
			args:     args{content: testContentAll4},
			expected: []int{161},
		},
		{
			name:     "5_" + testFilename5,
			args:     args{content: testContentAll5},
			expected: []int{286},
		},
		{
			name:     "6_" + testFilename6,
			args:     args{content: testContentAll6},
			expected: []int{894, 1439},
		},
		{
			name:     "7_" + testFilename7,
			args:     args{content: testContentAll7},
			expected: []int{186},
		},
		{
			name:     "8_" + testFilename8,
			args:     args{content: testContentAll8},
			expected: []int{},
		},
		{
			name:     "9_" + testFilename9,
			args:     args{content: testContentAll9},
			expected: []int{1660},
		},
		{
			name:     "10_" + testFilename10,
			args:     args{content: testContentAll10},
			expected: []int{635},
		},
		{
			name:     "11_" + testFilename11,
			args:     args{content: testContentAll11},
			expected: []int{1514},
		},
		{
			name:     "12_" + testFilename12,
			args:     args{content: testContentAll12},
			expected: []int{752},
		},
		{
			name:     "13_" + testFilename13,
			args:     args{content: testContentAll13},
			expected: []int{335},
		},
		{
			name:     "14_" + testFilename14,
			args:     args{content: testContentAll14},
			expected: []int{3950},
		},
		{
			name:     "15_" + testFilename15,
			args:     args{content: testContentAll15},
			expected: []int{3228, 8988, 9980},
		},
		{
			name:     "16_" + testFilename16,
			args:     args{content: testContentAll16},
			expected: []int{315},
		},
		{
			name:     "17_" + testFilename17,
			args:     args{content: testContentAll17},
			expected: []int{453},
		},
		{
			name:     "18_" + testFilename18,
			args:     args{content: testContentAll18},
			expected: []int{291},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tryGetSimpleSignatureIdxes(tt.args.content); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("tryGetSimpleSignatureIdxes() = %v, want %v", got, tt.expected)
			}
		})
	}
}
