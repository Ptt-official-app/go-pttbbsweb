package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestParseContent(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		content []byte
		origMD5 string
	}
	tests := []struct {
		name                  string
		args                  args
		expectedContent       [][]*types.Rune
		expectedContentMD5    string
		expectedIp            string
		expectedHost          string
		expectedBbs           string
		expectedSignatureMD5  string
		expectedSignatureDBCS []byte
		expectedCommentsDBCS  []byte
	}{
		// TODO: Add test cases.
		{
			name:                  "0_" + testFilename0,
			args:                  args{content: testContentAll0},
			expectedCommentsDBCS:  testComment0,
			expectedContent:       testContent0Utf8,
			expectedContentMD5:    "OH8piMm-mPtCLwSJa3QmsQ",
			expectedIp:            "172.18.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: testSignature0,
			expectedSignatureMD5:  "3Y6qKQU-r-Tl9gf2Gk0r3w",
		},
		{
			name:                  "1_" + testFilename1,
			args:                  args{content: testContentAll1},
			expectedCommentsDBCS:  testComment1,
			expectedContent:       testContent1Utf8,
			expectedContentMD5:    "Gs6Ppca4-OYxHUGYqxj-NA",
			expectedIp:            "172.18.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: testSignature1,
			expectedSignatureMD5:  "3Y6qKQU-r-Tl9gf2Gk0r3w",
		},
		{
			name:                  "2_" + testFilename2,
			args:                  args{content: testContentAll2},
			expectedCommentsDBCS:  testComment2,
			expectedContent:       testContent2Utf8,
			expectedContentMD5:    "j7h-dTQQaxjE9Nhlf0HJLg",
			expectedIp:            "172.19.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: testSignature2,
			expectedSignatureMD5:  "KTh3YC3vYsrqRfCgzKHOSw",
		},
		{
			name:                  "3_" + testFilename3,
			args:                  args{content: testContentAll3},
			expectedCommentsDBCS:  testComment3,
			expectedContent:       testContent3Utf8,
			expectedContentMD5:    "L6QISYJFt-Y5g4Thl-roaw",
			expectedIp:            "172.22.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: testSignature3,
			expectedSignatureMD5:  "tAeApguVSsnkqWWZlrdEmg",
		},
		{
			name:                  "4_" + testFilename4,
			args:                  args{content: testContentAll4},
			expectedCommentsDBCS:  testComment4,
			expectedContent:       testContent4Utf8,
			expectedContentMD5:    "riiRuKCZzG0gAGpQiq4GJA",
			expectedIp:            "172.22.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: testSignature4,
			expectedSignatureMD5:  "tAeApguVSsnkqWWZlrdEmg",
		},
		{
			name:                  "5_" + testFilename5,
			args:                  args{content: testContentAll5},
			expectedCommentsDBCS:  testComment5,
			expectedContent:       testContent5Utf8,
			expectedContentMD5:    "w-zS9i1tkPlPrJA5Xcnb2g",
			expectedIp:            "128.211.200.242",
			expectedBbs:           "批踢踢實業坊(ptt.csie.ntu.edu.tw)",
			expectedSignatureDBCS: testSignature5,
			expectedSignatureMD5:  "3_sDciyphzUMnvY-mj0LFg",
		},
		{
			name:                  "6_" + testFilename6,
			args:                  args{content: testContentAll6},
			expectedCommentsDBCS:  testComment6,
			expectedContent:       testContent6Utf8,
			expectedContentMD5:    "aWhyR589muapoSYPqdtTFg",
			expectedIp:            "110.50.185.154",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature6,
			expectedSignatureMD5:  "LQsvJ759e9HPG_xh_6olNQ",
		},
		{
			name:                  "7_" + testFilename7,
			args:                  args{content: testContentAll7},
			expectedCommentsDBCS:  testComment7,
			expectedContent:       testContent7Utf8,
			expectedContentMD5:    "D4gKa2HsRHIdzaoMGiiYug",
			expectedIp:            "140.112.29.66",
			expectedBbs:           "批踢踢實業坊(ptt.csie.ntu.edu.tw)",
			expectedSignatureDBCS: testSignature7,
			expectedSignatureMD5:  "SjpbD2MHrOFYMFlsjvOlwg",
		},
		{
			name:                  "8_" + testFilename8,
			args:                  args{content: testContentAll8},
			expectedCommentsDBCS:  testComment8,
			expectedContent:       testContent8Utf8,
			expectedContentMD5:    "mOYvLUjoHU8PJLAyex4SkA",
			expectedIp:            "",
			expectedBbs:           "",
			expectedSignatureDBCS: testSignature8,
			expectedSignatureMD5:  "",
		},
		{
			name:                  "9_" + testFilename9,
			args:                  args{content: testContentAll9},
			expectedCommentsDBCS:  testComment9,
			expectedContent:       testContent9Utf8,
			expectedContentMD5:    "AqsdMP43OqLyZo-MTrdWZg",
			expectedIp:            "150.117.235.183",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature9,
			expectedSignatureMD5:  "UI5darPb5Ru_rH_idf5IFg",
		},
		{
			name:                  "10_" + testFilename10,
			args:                  args{content: testContentAll10},
			expectedCommentsDBCS:  testComment10,
			expectedContent:       testContent10Utf8,
			expectedContentMD5:    "rJ2VkdnDw6mjwoeSrS837g",
			expectedIp:            "140.112.217.57",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature10,
			expectedSignatureMD5:  "mXwfm61PZ4-oxKslesLY0Q",
		},
		{
			name:                  "11_" + testFilename11,
			args:                  args{content: testContentAll11},
			expectedCommentsDBCS:  testComment11,
			expectedContent:       testContent11Utf8,
			expectedContentMD5:    "IKCj3KzpwP5pcJxOAPNDNQ",
			expectedIp:            "49.216.65.39",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature11,
			expectedSignatureMD5:  "KaENtKOzIRMWsNrLSF3R_w",
		},
		{
			name:                  "12_" + testFilename12,
			args:                  args{content: testContentAll12},
			expectedCommentsDBCS:  testComment12,
			expectedContent:       testContent12Utf8,
			expectedContentMD5:    "kQ8iQFTH15-9dixhSqAO7g",
			expectedIp:            "111.251.58.111",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature12,
			expectedSignatureMD5:  "32-ergj_D5G3Pv7azq1Fyg",
		},
		{
			name:                  "13_" + testFilename13,
			args:                  args{content: testContentAll13},
			expectedCommentsDBCS:  testComment13,
			expectedContent:       testContent13Utf8,
			expectedContentMD5:    "6sFGDlIC8gctuY5ahR829g",
			expectedIp:            "stego",
			expectedHost:          "",
			expectedBbs:           "批踢踢實業坊(ptt.csie.ntu.edu.tw)",
			expectedSignatureDBCS: testSignature13,
			expectedSignatureMD5:  "AUDBK6ixGaj43AQXqMrhdw",
		},
		{
			name:                  "14_" + testFilename14,
			args:                  args{content: testContentAll14},
			expectedCommentsDBCS:  testComment14,
			expectedContent:       testContent14Utf8,
			expectedContentMD5:    "L8LiexpRIrCfD1m__KuGew",
			expectedIp:            "111.249.44.44",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature14,
			expectedSignatureMD5:  "YZ0L3oByrUV7-OXE9o1OkQ",
		},
		{
			name:                  "15_" + testFilename15,
			args:                  args{content: testContentAll15},
			expectedCommentsDBCS:  testComment15,
			expectedContent:       testContent15Utf8,
			expectedContentMD5:    "ElYeuWMaZ5Vo7fdXk4kPfw",
			expectedIp:            "36.226.170.211",
			expectedHost:          "",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature15,
			expectedSignatureMD5:  "AJBO2K8Q2picE5Euph8UBg",
		},
		{
			name:                  "16_" + testFilename16,
			args:                  args{content: testContentAll16},
			expectedCommentsDBCS:  testComment16,
			expectedContent:       testContent16Utf8,
			expectedContentMD5:    "CG-Fa1c4E4yKvf0e23Pg_g",
			expectedIp:            "58.152.169.221",
			expectedHost:          "香港",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature16,
			expectedSignatureMD5:  "-khF0C89ne2bmKTWIT-A0g",
		},
		{
			name:                  "17_" + testFilename17,
			args:                  args{content: testContentAll17},
			expectedCommentsDBCS:  testComment17,
			expectedContent:       testContent17Utf8,
			expectedContentMD5:    "zXTRiGFwsgepwULPGCUt0w",
			expectedIp:            "123.193.200.197",
			expectedHost:          "臺灣",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: testSignature17,
			expectedSignatureMD5:  "abiQ2PiIZv9bOWdmOsg6_Q",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, gotContentMD5, gotIp, gotHost, gotBbs, gotSignatureMD5, gotSignatureDBCS, gotCommentsDBCS := ParseContent(tt.args.content, tt.args.origMD5)
			testutil.TDeepEqual(t, "got", gotContent, tt.expectedContent)

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
			if gotSignatureMD5 != tt.expectedSignatureMD5 {
				t.Errorf("ParseContent() gotSignatureMD5: %v want: %v", gotSignatureMD5, tt.expectedSignatureMD5)
			}
			if !reflect.DeepEqual(gotSignatureDBCS, tt.expectedSignatureDBCS) {
				t.Errorf("ParseContent() gotSignatureDBCS: %v want: %v", gotSignatureDBCS, tt.expectedSignatureDBCS)
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
