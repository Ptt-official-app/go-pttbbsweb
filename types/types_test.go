package types

import (
	"reflect"
	"testing"
	"time"

	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
)

func TestNanoTS_Base64(t *testing.T) {
	tests := []struct {
		name     string
		tr       NanoTS
		expected string
	}{
		// TODO: Add test cases.
		{
			tr:       NanoTS(0),
			expected: "AAAAAAAAAAA",
		},
		{
			tr:       NanoTS(1234567890000000000),
			expected: "ESIQ9HaNtAA",
		},
		{
			tr:       NanoTS(9223372036854775807),
			expected: "f_________8",
		},
		{
			tr:       NanoTS(-9223372036854775808),
			expected: "gAAAAAAAAAA",
		},
		{
			tr:       NanoTS(-1),
			expected: "__________8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Base64(); got != tt.expected {
				t.Errorf("NanoTS.Base64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToCommentID(t *testing.T) {
	type args struct {
		nanoTS NanoTS
		md5    string
	}
	tests := []struct {
		name     string
		args     args
		expected CommentID
	}{
		// TODO: Add test cases.
		{
			args:     args{nanoTS: NanoTS(0), md5: "testMD5"},
			expected: "AAAAAAAAAAA:testMD5",
		},
		{
			args:     args{nanoTS: NanoTS(1234567890000000000), md5: "testMD5"},
			expected: "ESIQ9HaNtAA:testMD5",
		},
		{
			args:     args{nanoTS: NanoTS(9223372036854775807), md5: "testMD5"},
			expected: "f_________8:testMD5",
		},
		{
			args:     args{nanoTS: NanoTS(-9223372036854775808), md5: "testMD5"},
			expected: "gAAAAAAAAAA:testMD5",
		},
		{
			args:     args{nanoTS: NanoTS(-1), md5: "testMD5"},
			expected: "__________8:testMD5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCommentID(tt.args.nanoTS, tt.args.md5); got != tt.expected {
				t.Errorf("ToCommentID() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToReplyID(t *testing.T) {
	type args struct {
		commentID CommentID
	}
	tests := []struct {
		name     string
		args     args
		expected CommentID
	}{
		// TODO: Add test cases.
		{
			args:     args{commentID: CommentID("gAAAAAAAAAA:testMD5")},
			expected: CommentID("gAAAAAAAAAA:testMD5:R"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToReplyID(tt.args.commentID); got != tt.expected {
				t.Errorf("ToReplyID() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTime4ToNanoTS(t *testing.T) {
	type args struct {
		t pttbbstypes.Time4
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
	}{
		// TODO: Add test cases.
		{
			args:     args{t: pttbbstypes.Time4(1234567890)},
			expected: NanoTS(1234567890000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time4ToNanoTS(tt.args.t); got != tt.expected {
				t.Errorf("Time4ToNanoTS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTime8_ToNanoTS(t *testing.T) {
	tests := []struct {
		name     string
		tr       Time8
		expected NanoTS
	}{
		// TODO: Add test cases.
		{
			tr:       Time8(1234567890),
			expected: NanoTS(1234567890000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToNanoTS(); got != tt.expected {
				t.Errorf("Time8.ToNanoTS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNanoTS_ToTime8(t *testing.T) {
	tests := []struct {
		name     string
		tr       NanoTS
		expected Time8
	}{
		// TODO: Add test cases.
		{
			tr:       NanoTS(1234567890000000000),
			expected: Time8(1234567890),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToTime8(); got != tt.expected {
				t.Errorf("NanoTS.ToTime8() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNanoTS_ToTime4(t *testing.T) {
	tests := []struct {
		name     string
		tr       NanoTS
		expected pttbbstypes.Time4
	}{
		// TODO: Add test cases.
		{
			tr:       NanoTS(1234567890000000000),
			expected: pttbbstypes.Time4(1234567890),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToTime4(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NanoTS.ToTime4() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNanoTS_ToTime(t *testing.T) {
	tests := []struct {
		name     string
		tr       NanoTS
		expected time.Time
	}{
		// TODO: Add test cases.
		{
			tr:       NanoTS(1234567890000000000),
			expected: time.Date(2009, time.Month(2), 14, 7, 31, 30, 0, TIMEZONE),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToTime(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NanoTS.ToTime() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNanoTS_ToNanoTSByMin(t *testing.T) {
	tests := []struct {
		name     string
		tr       NanoTS
		expected NanoTS
	}{
		// TODO: Add test cases.
		{
			tr:       NanoTS(1234567890000000000),
			expected: NanoTS(1234567860000000000),
		},
		{
			tr:       NanoTS(1234567860000000000),
			expected: NanoTS(1234567860000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToNanoTSByMin(); got != tt.expected {
				t.Errorf("NanoTS.ToNanoTSByMin() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBase64ToNanoTS(t *testing.T) {
	type args struct {
		b64 string
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{b64: "AAAAAAAAAAA"},
			expected: NanoTS(0),
		},
		{
			args:     args{b64: "ESIQ9HaNtAA"},
			expected: NanoTS(1234567890000000000),
		},
		{
			args:     args{b64: "f_________8"},
			expected: NanoTS(9223372036854775807),
		},
		{
			args:     args{b64: "gAAAAAAAAAA"},
			expected: NanoTS(-9223372036854775808),
		},
		{
			args:     args{b64: "__________8"},
			expected: NanoTS(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64ToNanoTS(tt.args.b64)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64ToNanoTS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Base64ToNanoTS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimeToNanoTS(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
	}{
		// TODO: Add test cases.
		{
			args:     args{t: time.Date(2009, time.Month(2), 14, 7, 31, 30, 0, TIMEZONE)},
			expected: NanoTS(1234567890000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToNanoTS(tt.args.t); got != tt.expected {
				t.Errorf("TimeToNanoTS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDateStrToTime(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{"2009/02/14"},
			expected: NanoTS(1234540800000000000), //2009-02-14 00:00:00 CST
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateStrToTime(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateStrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(TimeToNanoTS(got), tt.expected) {
				t.Errorf("DateStrToTime() = %v, want %v", TimeToNanoTS(got), tt.expected)
			}
		})
	}
}

func TestDateYearTimeStrToTime(t *testing.T) {
	type args struct {
		dateYearTimeStr string
	}
	tests := []struct {
		name     string
		args     args
		expected time.Time
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{"02/14/2009 07:31:30"},
			expected: time.Date(2009, time.Month(2), 14, 7, 31, 30, 0, TIMEZONE),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateYearTimeStrToTime(tt.args.dateYearTimeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateYearTimeStrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DateYearTimeStrToTime() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDateMinStrToTime(t *testing.T) {
	type args struct {
		dateTimeStr string
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{dateTimeStr: "2009/02/14 07:31"},
			expected: NanoTS(1234567860000000000), //1970-02-14 07:31:00 CST
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateMinStrToTime(tt.args.dateTimeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTimeStrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(TimeToNanoTS(got), tt.expected) {
				t.Errorf("DateTimeStrToTime() = %v, want %v", TimeToNanoTS(got), tt.expected)
			}
		})
	}
}

func TestNewDateTime(t *testing.T) {
	type args struct {
		year   int
		month  time.Month
		day    int
		hr     int
		minute int
		second int
	}
	tests := []struct {
		name     string
		args     args
		expected NanoTS
	}{
		// TODO: Add test cases.
		{
			args:     args{year: 2009, month: time.Month(2), day: 14, hr: 7, minute: 31, second: 30},
			expected: NanoTS(1234567890000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDateTime(tt.args.year, tt.args.month, tt.args.day, tt.args.hr, tt.args.minute, tt.args.second); !reflect.DeepEqual(TimeToNanoTS(got), tt.expected) {
				t.Errorf("NewDateTime() = %v, want %v", got, tt.expected)
			}
		})
	}
}
