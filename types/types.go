package types

import (
	"encoding/base64"
	"encoding/binary"
	"time"

	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
)

//CommentID
//
//XXX currently it's very hard to maintain the comment-id.
//if we do comment-id only based on MD5:
//  got duplicated md5-id if the owner posts the same comments
//  within 1 min.
//
//if we add the inferred CreateTime into the comment-id:
//  the CreateTime may be changed if the author deletes
//  some other comments within same minute.
type CommentID string

func ToCommentID(nanoTS NanoTS, md5 string) CommentID {
	return CommentID(nanoTS.Base64() + ":" + md5)
}
func ToReplyID(commentID CommentID) CommentID {
	return commentID + ":R"
}

const (
	TS_TO_NANO_TS  = NanoTS(1000000000) //10^9
	MIN_TO_NANO_TS = NanoTS(60) * TS_TO_NANO_TS
	YEAR_NANO_TS   = NanoTS(365*86400) * TS_TO_NANO_TS
)

type NanoTS int64

func Time4ToNanoTS(t pttbbstypes.Time4) NanoTS {
	return NanoTS(t) * TS_TO_NANO_TS
}

type Time8 int64

func (t Time8) ToNanoTS() NanoTS {
	return NanoTS(t) * TS_TO_NANO_TS
}

func (t NanoTS) ToTime8() Time8 {
	return Time8(t / TS_TO_NANO_TS)
}

func (t NanoTS) ToTime4() pttbbstypes.Time4 {
	return pttbbstypes.Time4(t / TS_TO_NANO_TS)
}

func (t NanoTS) ToTime() time.Time {
	return time.Unix(int64(t/TS_TO_NANO_TS), int64(t%TS_TO_NANO_TS)).In(TIMEZONE)
}

//ToNanoTSByMin
//
//nano-ts by minutes.
//used in comment-time-slot.
func (t NanoTS) ToNanoTSByMin() NanoTS {
	return (t / MIN_TO_NANO_TS) * MIN_TO_NANO_TS
}

func (t NanoTS) Base64() string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t))

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(buf)
}

func Base64ToNanoTS(b64 string) (NanoTS, error) {
	buf, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(b64)
	if err != nil {
		return 0, err
	}
	nanoTS_u := binary.BigEndian.Uint64(buf)

	return NanoTS(nanoTS_u), nil
}

func NowNanoTS() NanoTS {
	return TimeToNanoTS(time.Now())
}

func TimeToNanoTS(t time.Time) NanoTS {
	return NanoTS(t.UnixNano())
}

//DateStrToTime
//
//(YYYY/)MM/DD
//
//used in old-comment.
//
//golang requires year to parse TIMEZONE with CST.
//https://github.com/golang/go/issues/34101
func DateStrToTime(dateStr string) (time.Time, error) {
	layout := "2006/01/02"
	t, err := time.ParseInLocation(layout, dateStr, TIMEZONE)
	if err != nil {
		return t, err
	}
	return t.In(TIMEZONE), nil
}

//DateYearTimeStrToTime
//
//MM/DD/YYYY hh:mm:ss
//
//used in edit (編輯)
func DateYearTimeStrToTime(dateYearTimeStr string) (time.Time, error) {
	layout := "01/02/2006 15:04:05"
	t, err := time.ParseInLocation(layout, dateYearTimeStr, TIMEZONE)
	if err != nil {
		return t, err
	}
	return t.In(TIMEZONE), nil
}

//DateMinStrToTime
//
//(YYYY/)MM/DD hh:mm
//
//used in comment.
//
//golang requires year to parse TIMEZONE with CST.
//https://github.com/golang/go/issues/34101
func DateMinStrToTime(dateTimeStr string) (time.Time, error) {
	layout := "2006/01/02 15:04"
	t, err := time.ParseInLocation(layout, dateTimeStr, TIMEZONE)
	if err != nil {
		return t, err
	}
	return t.In(TIMEZONE), nil
}

func NewDateTime(year int, month time.Month, day int, hr int, minute int, second int) time.Time {
	return time.Date(year, month, day, hr, minute, second, 0, TIMEZONE).In(TIMEZONE)
}
