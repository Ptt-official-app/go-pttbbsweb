package dbcs

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestParseContentStr(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		contentStr     string
		origContentMD5 string
	}
	tests := []struct {
		name                  string
		args                  args
		expectedContent       [][]*types.Rune
		expectedContentPrefix [][]*types.Rune
		expectedContentMD5    string
		expectedIp            string
		expectedHost          string
		expectedBbs           string
		expectedSignatureMD5  string
		expectedSignatureDBCS string
		expectedCommentsDBCS  string
	}{
		// TODO: Add test cases.
		{
			name:                  "0_" + testUtf8Filename0,
			args:                  args{contentStr: string(testUtf8ContentAll0)},
			expectedCommentsDBCS:  string(testUtf8Comment0),
			expectedContent:       testUtf8Content0Utf8[4:],
			expectedContentPrefix: testUtf8Content0Utf8[:4],
			expectedContentMD5:    "Y-TVLUfQ7X3ZmkNWBAk9yA",
			expectedIp:            "172.18.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: string(testUtf8Signature0),
			expectedSignatureMD5:  "2PKyrdISYcwwi0zBqNC-ww",
		},
		{
			name:                  "1_" + testUtf8Filename1,
			args:                  args{contentStr: string(testUtf8ContentAll1)},
			expectedCommentsDBCS:  string(testUtf8Comment1),
			expectedContent:       testUtf8Content1Utf8[4:],
			expectedContentPrefix: testUtf8Content1Utf8[:4],
			expectedContentMD5:    "drA4l_UBKWaEMxEBg2qZ8w",
			expectedHost:          "臺灣",
			expectedIp:            "101.12.88.158",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: string(testUtf8Signature1),
			expectedSignatureMD5:  "fedy_wwrsI0WaIYtxcaWUQ",
		},
		{
			name:                  "2_" + testUtf8Filename2,
			args:                  args{contentStr: string(testUtf8ContentAll2)},
			expectedCommentsDBCS:  string(testUtf8Comment2),
			expectedContent:       testUtf8Content2Utf8[4:],
			expectedContentPrefix: testUtf8Content2Utf8[:4],
			expectedContentMD5:    "9EdTkvqJiqF5W-4bkVK9Uw",
			expectedHost:          "",
			expectedIp:            "172.18.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: string(testUtf8Signature2),
			expectedSignatureMD5:  "2PKyrdISYcwwi0zBqNC-ww",
		},
		{
			name:                  "3_" + testUtf8Filename3,
			args:                  args{contentStr: string(testUtf8ContentAll3)},
			expectedCommentsDBCS:  string(testUtf8Comment3),
			expectedContent:       testUtf8Content3Utf8[4:],
			expectedContentPrefix: testUtf8Content3Utf8[:4],
			expectedContentMD5:    "-M_XfQpEfjaOCd1HYznyVQ",
			expectedHost:          "",
			expectedIp:            "172.19.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: string(testUtf8Signature3),
			expectedSignatureMD5:  "JxM9jx0VoUiceI-TjTbanQ",
		},
		{
			name:                  "4_" + testUtf8Filename4,
			args:                  args{contentStr: string(testUtf8ContentAll4)},
			expectedCommentsDBCS:  string(testUtf8Comment4),
			expectedContent:       testUtf8Content4Utf8[4:],
			expectedContentPrefix: testUtf8Content4Utf8[:4],
			expectedContentMD5:    "MpnMgVtdAhQnf1f9wIH9EA",
			expectedHost:          "",
			expectedIp:            "172.22.0.1",
			expectedBbs:           "批踢踢 docker(pttdocker.test)",
			expectedSignatureDBCS: string(testUtf8Signature4),
			expectedSignatureMD5:  "2By-wvJx0g5Zb3YLET75Xg",
		},
		{
			name:                  "5_" + testUtf8Filename5,
			args:                  args{contentStr: string(testUtf8ContentAll5)},
			expectedCommentsDBCS:  string(testUtf8Comment5),
			expectedContent:       testUtf8Content5Utf8[4:],
			expectedContentPrefix: testUtf8Content5Utf8[:4],
			expectedContentMD5:    "58yrYdHg3mX-I4WG3T4Ciw",
			expectedHost:          "臺灣",
			expectedIp:            "123.193.200.197",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: string(testUtf8Signature5),
			expectedSignatureMD5:  "1lG3T4CLtCMV_XDTOt6rqg",
		},
		{
			name:                  "6_" + testUtf8Filename6,
			args:                  args{contentStr: string(testUtf8ContentAll6)},
			expectedCommentsDBCS:  string(testUtf8Comment6),
			expectedContent:       testUtf8Content6Utf8[4:],
			expectedContentPrefix: testUtf8Content6Utf8[:4],
			expectedContentMD5:    "6sUPTY66DSnJCxeiJPUzuw",
			expectedHost:          "臺灣",
			expectedIp:            "223.138.218.62",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: string(testUtf8Signature6),
			expectedSignatureMD5:  "0mwvGdNlr06vfEfF8K8s7w",
		},
		{
			name:                  "7_" + testUtf8Filename7,
			args:                  args{contentStr: string(testUtf8ContentAll7)},
			expectedCommentsDBCS:  string(testUtf8Comment7),
			expectedContent:       testUtf8Content7Utf8[4:],
			expectedContentPrefix: testUtf8Content7Utf8[:4],
			expectedContentMD5:    "uy8-lWMYCSbrGcRgaj4Gqw",
			expectedHost:          "",
			expectedIp:            "115.43.83.245",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: string(testUtf8Signature7),
			expectedSignatureMD5:  "LoYyhk9-5ed7MUbfkcpVWQ",
		},
		{
			name:                  "8_" + testUtf8Filename8,
			args:                  args{contentStr: string(testUtf8ContentAll8)},
			expectedCommentsDBCS:  string(testUtf8Comment8),
			expectedContent:       testUtf8Content8Utf8[4:],
			expectedContentPrefix: testUtf8Content8Utf8[:4],
			expectedContentMD5:    "g9HPCGQWB4ZV0IOCmL2HUA",
			expectedHost:          "",
			expectedIp:            "163.27.69.176",
			expectedBbs:           "批踢踢實業坊(ptt.cc)",
			expectedSignatureDBCS: string(testUtf8Signature8),
			expectedSignatureMD5:  "-PyaAi9zz5BqP8iqQyUOPw",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotContent, gotContentPrefix, gotContentMD5, gotIp, gotHost, gotBbs, gotSignatureMD5, gotSignatureDBCS, gotCommentsDBCS := ParseContentStr(tt.args.contentStr, tt.args.origContentMD5, true)
			testutil.TDeepEqual(t, "content", gotContent, tt.expectedContent)
			testutil.TDeepEqual(t, "prefix", gotContentPrefix, tt.expectedContentPrefix)
			if gotContentMD5 != tt.expectedContentMD5 {
				t.Errorf("ParseContentStr() gotContentMD5 = %v, want %v", gotContentMD5, tt.expectedContentMD5)
			}
			if gotIp != tt.expectedIp {
				t.Errorf("ParseContentStr() gotIp = %v, want %v", gotIp, tt.expectedIp)
			}
			if gotHost != tt.expectedHost {
				t.Errorf("ParseContentStr() gotHost = %v, want %v", gotHost, tt.expectedHost)
			}
			if gotBbs != tt.expectedBbs {
				t.Errorf("ParseContentStr() gotBbs = %v, want %v", gotBbs, tt.expectedBbs)
			}
			if gotSignatureMD5 != tt.expectedSignatureMD5 {
				t.Errorf("ParseContentStr() gotSignatureMD5 = %v, want %v", gotSignatureMD5, tt.expectedSignatureMD5)
			}
			testutil.TDeepEqual(t, "signatureDBCS", gotSignatureDBCS, tt.expectedSignatureDBCS)
			testutil.TDeepEqual(t, "commentsDBCS", gotCommentsDBCS, tt.expectedCommentsDBCS)
		})
		wg.Wait()
	}
}
