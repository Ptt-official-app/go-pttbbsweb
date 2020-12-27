package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func TestParseContent(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		content []byte
		origMD5 string
	}
	tests := []struct {
		name                 string
		args                 args
		expectedContent_r    [][]*types.Rune
		expectedContentMD5   string
		expectedIp           string
		expectedHost         string
		expectedBbs          string
		expectedCommentsDBCS []byte
	}{
		// TODO: Add test cases.
		{
			name:                 "0_" + testFilename0,
			args:                 args{content: testContentAll0},
			expectedCommentsDBCS: testComment0,
			expectedContent_r:    testContent0Utf8,
			expectedContentMD5:   "uiyjjLmG88ABAFUWyM4Llg",
			expectedIp:           "172.18.0.1",
			expectedBbs:          "批踢踢 docker(pttdocker.test)",
		},
		{
			name:                 "1_" + testFilename1,
			args:                 args{content: testContentAll1},
			expectedCommentsDBCS: testComment1,
			expectedContent_r:    testContent1Utf8,
			expectedContentMD5:   "aRc6U80bRtOY62yXKkpAqw",
			expectedIp:           "172.18.0.1",
			expectedBbs:          "批踢踢 docker(pttdocker.test)",
		},
		{
			name:                 "2_" + testFilename2,
			args:                 args{content: testContentAll2},
			expectedCommentsDBCS: testComment2,
			expectedContent_r:    testContent2Utf8,
			expectedContentMD5:   "rgwnxxLXbL8W4UdbQVN0pQ",
			expectedIp:           "172.19.0.1",
			expectedBbs:          "批踢踢 docker(pttdocker.test)",
		},
		{
			name:                 "3_" + testFilename3,
			args:                 args{content: testContentAll3},
			expectedCommentsDBCS: testComment3,
			expectedContent_r:    testContent3Utf8,
			expectedContentMD5:   "2hfH0_ofkV5BgHJMr1H2tg",
			expectedIp:           "172.22.0.1",
			expectedBbs:          "批踢踢 docker(pttdocker.test)",
		},
		{
			name:                 "4_" + testFilename4,
			args:                 args{content: testContentAll4},
			expectedCommentsDBCS: testComment4,
			expectedContent_r:    testContent4Utf8,
			expectedContentMD5:   "TD1vkp4KtB5bqVEFubzuOw",
			expectedIp:           "172.22.0.1",
			expectedBbs:          "批踢踢 docker(pttdocker.test)",
		},
		{
			name:                 "5_" + testFilename5,
			args:                 args{content: testContentAll5},
			expectedCommentsDBCS: testComment5,
			expectedContent_r:    testContent5Utf8,
			expectedContentMD5:   "pJfBYcAQfVEy3qo4dFTK6Q",
			expectedIp:           "128.211.200.242",
			expectedBbs:          "批踢踢實業坊(ptt.csie.ntu.edu.tw)",
		},
		{
			name:                 "6_" + testFilename6,
			args:                 args{content: testContentAll6},
			expectedCommentsDBCS: testComment6,
			expectedContent_r:    testContent6Utf8,
			expectedContentMD5:   "U-iJYSa7oipPsxWA3uc0JA",
			expectedIp:           "110.50.185.154",
			expectedHost:         "臺灣",
			expectedBbs:          "批踢踢實業坊(ptt.cc)",
		},
		{
			name:                 "7_" + testFilename7,
			args:                 args{content: testContentAll7},
			expectedCommentsDBCS: testComment7,
			expectedContent_r:    testContent7Utf8,
			expectedContentMD5:   "gAWCYGqeUfVSVJThJ7Piyg",
			expectedIp:           "140.112.29.66",
			expectedBbs:          "批踢踢實業坊(ptt.csie.ntu.edu.tw)",
		},
		{
			name:                 "8_" + testFilename8,
			args:                 args{content: testContentAll8},
			expectedCommentsDBCS: testComment8,
			expectedContent_r:    testContent8Utf8,
			expectedContentMD5:   "_5lx9kHXbXzf-nY0th0Qug",
			expectedIp:           "",
			expectedBbs:          "",
		},
		{
			name:                 "9_" + testFilename9,
			args:                 args{content: testContentAll9},
			expectedCommentsDBCS: testComment9,
			expectedContent_r:    testContent9Utf8,
			expectedContentMD5:   "_6Pa5fD4PQtUqqL2LQflAw",
			expectedIp:           "150.117.235.183",
			expectedHost:         "臺灣",
			expectedBbs:          "批踢踢實業坊(ptt.cc)",
		},
		{
			name:                 "10_" + testFilename10,
			args:                 args{content: testContentAll10},
			expectedCommentsDBCS: testComment10,
			expectedContent_r:    testContent10Utf8,
			expectedContentMD5:   "8pICfQNUA9TppiHyDHOWdA",
			expectedIp:           "140.112.217.57",
			expectedHost:         "臺灣",
			expectedBbs:          "批踢踢實業坊(ptt.cc)",
		},
		{
			name:                 "11_" + testFilename11,
			args:                 args{content: testContentAll11},
			expectedCommentsDBCS: testComment11,
			expectedContent_r:    testContent11Utf8,
			expectedContentMD5:   "d8jUdTahmDo4R_zkO6yr5g",
			expectedIp:           "49.216.65.39",
			expectedHost:         "臺灣",
			expectedBbs:          "批踢踢實業坊(ptt.cc)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent_r, gotContentMD5, gotIp, gotHost, gotBbs, gotCommentsDBCS := ParseContent(tt.args.content, tt.args.origMD5)
			if !reflect.DeepEqual(gotContent_r, tt.expectedContent_r) {
				t.Errorf("ParseContent() gotContent_r = %v, want %v", gotContent_r, tt.expectedContent_r)
			}
			if gotContentMD5 != tt.expectedContentMD5 {
				t.Errorf("ParseContent() gotContentMD5 = %v, want %v", gotContentMD5, tt.expectedContentMD5)
			}
			if gotIp != tt.expectedIp {
				t.Errorf("ParseContent() gotIp = %v, want %v", gotIp, tt.expectedIp)
			}
			if gotHost != tt.expectedHost {
				t.Errorf("ParseContent() gotHost = %v, want %v", gotHost, tt.expectedHost)
			}
			if gotBbs != tt.expectedBbs {
				t.Errorf("ParseContent() gotBbs = %v, want %v", gotBbs, tt.expectedBbs)
			}
			if !reflect.DeepEqual(gotCommentsDBCS, tt.expectedCommentsDBCS) {
				t.Errorf("ParseContent() gotCommentsDBCS = %v, want %v", gotCommentsDBCS, tt.expectedCommentsDBCS)
			}
		})
	}
}

func Test_parseIPHostBBS0(t *testing.T) {
	setupTest()
	defer teardownTest()
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		args         args
		expectedIp   string
		expectedHost string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp, gotHost := parseIPHostBBS0(tt.args.str)
			if gotIp != tt.expectedIp {
				t.Errorf("parseIPHostBBS0() gotIp = %v, want %v", gotIp, tt.expectedIp)
			}
			if gotHost != tt.expectedHost {
				t.Errorf("parseIPHostBBS0() gotHost = %v, want %v", gotHost, tt.expectedHost)
			}
		})
	}
}
