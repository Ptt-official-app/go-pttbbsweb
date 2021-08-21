package schema

import (
	"reflect"
	"testing"
)

func Test_articleTitleToTitleRegexCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	title0 := []rune("abcd")
	nGram0 := 3
	expected0 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd"}

	title1 := []rune("abcd")
	nGram1 := 4
	expected1 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title2 := []rune("abcd")
	nGram2 := 5
	expected2 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title3 := []rune("再來呢？～")
	nGram3 := 5
	expected3 := []string{"再", "來", "呢", "？", "～", "再來", "來呢", "呢？", "？～", "再來呢", "來呢？", "呢？～", "再來呢？", "來呢？～", "再來呢？～"}

	type args struct {
		title      []rune
		titleRegex []string
		nGram      int
	}
	tests := []struct {
		name                  string
		args                  args
		expectedNewTitleRegex []string
	}{
		// TODO: Add test cases.
		{
			name:                  "abcd with 3",
			args:                  args{title: title0, nGram: nGram0},
			expectedNewTitleRegex: expected0,
		},
		{
			name:                  "abcd with 4",
			args:                  args{title: title1, nGram: nGram1},
			expectedNewTitleRegex: expected1,
		},
		{
			name:                  "abcd with 5",
			args:                  args{title: title2, nGram: nGram2},
			expectedNewTitleRegex: expected2,
		},
		{
			name:                  "再來呢？～ with 5",
			args:                  args{title: title3, nGram: nGram3},
			expectedNewTitleRegex: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewTitleRegex := articleTitleToTitleRegexCore(tt.args.title, tt.args.titleRegex, tt.args.nGram); !reflect.DeepEqual(gotNewTitleRegex, tt.expectedNewTitleRegex) {
				t.Errorf("articleTitleToTitleRegexCore() = %v, want %v", gotNewTitleRegex, tt.expectedNewTitleRegex)
			}
		})
	}
}

func Test_articleTitleToTitleRegex(t *testing.T) {
	setupTest()
	defer teardownTest()

	title0 := "abcd"
	expected0 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title1 := "abcde"
	expected1 := []string{"a", "b", "c", "d", "e", "ab", "bc", "cd", "de", "abc", "bcd", "cde", "abcd", "bcde", "abcde"}

	title2 := "abcdef"
	expected2 := []string{"a", "b", "c", "d", "e", "f", "ab", "bc", "cd", "de", "ef", "abc", "bcd", "cde", "def", "abcd", "bcde", "cdef", "abcde", "bcdef"}

	title3 := "abcdef"
	expected3 := []string{"a", "b", "c", "d", "e", "f", "ab", "bc", "cd", "de", "ef", "abc", "bcd", "cde", "def", "abcd", "bcde", "cdef", "abcde", "bcdef"}

	type args struct {
		title string
	}
	tests := []struct {
		name               string
		args               args
		expectedTitleRegex []string
	}{
		// TODO: Add test cases.
		{
			name:               "title only with abcd",
			args:               args{title: title0},
			expectedTitleRegex: expected0,
		},
		{
			name:               "title only with abcde",
			args:               args{title: title1},
			expectedTitleRegex: expected1,
		},
		{
			name:               "title only with abcdef",
			args:               args{title: title2},
			expectedTitleRegex: expected2,
		},
		{
			name:               "title with class",
			args:               args{title: title3},
			expectedTitleRegex: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTitleRegex := articleTitleToTitleRegex(tt.args.title); !reflect.DeepEqual(gotTitleRegex, tt.expectedTitleRegex) {
				t.Errorf("articleTitleToTitleRegex() = %v, want %v", gotTitleRegex, tt.expectedTitleRegex)
			}
		})
	}
}
