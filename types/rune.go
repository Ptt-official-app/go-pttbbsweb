package types

import (
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Rune struct {
	Utf8    string `json:"text" bson:"utf8"` // utf8-string
	Big5    []byte `json:"-" bson:"big5"`    // big5-bytes, stored in db for debugging.
	DBCS    []byte `json:"-" bson:"dbcs"`    // dbcs-bytes, stored in db for concat and debugging.
	DBCSStr string `json:"-" bson:"dbcsstr"` // dbcs-str, stored in db for concat and debugging.
	Color0  Color  `json:"color0" bson:"color0"`
	Color1  Color  `json:"color1" bson:"color1"`
}

var (
	utf8ToBig5 = make(map[string][]byte)
	big5ToUTF8 = make(map[string][]byte)
)

func initBig5() (err error) {
	err = initB2U()
	if err != nil {
		return err
	}

	err = initU2B()
	if err != nil {
		return err
	}

	return nil
}

func initB2U() error {
	if len(big5ToUTF8) > 0 { // already loaded
		return nil
	}

	file, err := os.Open(BIG5_TO_UTF8)
	if err != nil {
		return err
	}
	defer file.Close()

	r := io.Reader(file)
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")
	lines = lines[1:]

	for _, line := range lines {
		lineList := strings.Split(line, " ")
		if len(lineList) != 2 {
			logrus.Warnf("initB2U: unable to split line: line: %v", line)
			continue
		}

		big5, err := initToBig5(lineList[0][2:])
		if err != nil {
			logrus.Warnf("initB2U: unable to initToBig5: lineList: %v", lineList[0])
			continue
		}
		utf8, err := initToUtf8(lineList[1][2:])
		if err != nil {
			logrus.Warnf("initB2U: unable to initToUtf8: lineList: %v", lineList[1])
			continue
		}
		big5ToUTF8[string(big5)] = utf8
	}

	logrus.Infof("initB2U: after map: big5ToUTF8: %v", len(big5ToUTF8))

	return nil
}

func initU2B() error {
	if len(utf8ToBig5) > 0 { // already loaded
		return nil
	}

	file, err := os.Open(UTF8_TO_BIG5)
	if err != nil {
		return err
	}
	defer file.Close()

	r := io.Reader(file)
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")
	lines = lines[1:]

	for _, line := range lines {
		lineList := strings.Split(line, " ")
		if len(lineList) != 2 {
			logrus.Warnf("initU2B: unable to split line: line: %v", line)
			continue
		}

		big5, err := initToBig5(lineList[0][2:])
		if err != nil {
			logrus.Warnf("initU2B: unable to initToBig5: lineList: %v", lineList[0])
			continue
		}
		utf8, err := initToUtf8(lineList[1][2:])
		if err != nil {
			logrus.Warnf("initU2B: unable to initToUtf8: lineList: %v", lineList[1])
			continue
		}
		utf8ToBig5[string(utf8)] = big5
	}

	logrus.Infof("initU2B: after map: utf8ToBig5: %v", len(utf8ToBig5))

	return nil
}

func initToBig5(big5Code string) (theBytes []byte, err error) {
	big5Code = strings.TrimSpace(big5Code)
	theBytes = make([]byte, 2)

	_, err = hex.Decode(theBytes, []byte(big5Code))
	if err != nil {
		return nil, err
	}

	return theBytes, nil
}

func initToUtf8(ucsCode string) (theBytes []byte, err error) {
	ucsCode = strings.TrimSpace(ucsCode)
	ucsBytes := make([]byte, 2)

	_, err = hex.Decode(ucsBytes, []byte(ucsCode))
	if err != nil {
		return nil, err
	}

	ucs2 := int(ucsBytes[0])*256 + int(ucsBytes[1])

	if (ucs2 & (^0x7f)) == 0 {
		theBytes := make([]byte, 1)
		return theBytes, nil
	}

	if (ucs2 & 0xF800) == 0 {
		// (2) 00000yyy yyxxxxxx -> 110yyyyy 10xxxxxx
		theBytes := make([]byte, 2)
		theBytes[0] = byte(0xc0 | (ucs2 >> 6))
		theBytes[1] = byte(0x80 | (ucs2 & 0x3f))
		return theBytes, nil

	} else {
		// (3) zzzzyyyy yyxxxxxx -> 1110zzzz 10yyyyyy 10xxxxxx
		theBytes := make([]byte, 3)

		theBytes[0] = byte(0xE0 | (ucs2 >> 12))
		theBytes[1] = byte(0x80 | ((ucs2 >> 6) & 0x3F))
		theBytes[2] = byte(0x80 | ((ucs2) & 0x3F))
		return theBytes, nil
	}
}

func (r *Rune) Big5ToUtf8() {
	if r.Utf8 != "" {
		return
	}
	r.Utf8 = Big5ToUtf8(r.Big5)
}

func Big5ToUtf8(big5 []byte) (utf8 string) {
	estimatedUtf8Sz := len(big5) * 3 / 2
	utf8Bytes := make([]byte, 0, estimatedUtf8Sz)
	for p_big5 := big5; len(p_big5) > 0; {
		if p_big5[0] < 0x80 {
			utf8Bytes = append(utf8Bytes, p_big5[0])
			p_big5 = p_big5[1:]
		} else {
			if len(p_big5) < 2 {
				logrus.Warnf("Big5ToUtf8: unable to parse big5: big5: %v p_big5: %v", big5, p_big5)
				break
			}
			eachUtf8 := big5ToUTF8[string(p_big5[:2])]
			utf8Bytes = append(utf8Bytes, eachUtf8...)
			p_big5 = p_big5[2:]
		}
	}
	utf8 = string(utf8Bytes)

	return utf8
}

func (r *Rune) Utf8ToBig5() {
	if r.Big5 != nil {
		return
	}
	r.Big5 = Utf8ToBig5(r.Utf8)
}

func (r *Rune) Utf8ToDBCS(origColor Color, isAddCR bool) {
	if len(r.DBCS) != 0 {
		return
	}

	estimatedNBytes := len(r.Utf8)*2/3 + 2*DEFAULT_LEN_COLOR_BYTES
	theBytes := make([]byte, 0, estimatedNBytes)
	colorBytes := r.Color0.BytesWithPreColor(&origColor)
	theBytes = append(theBytes, colorBytes...)

	big5Bytes := Utf8ToBig5(r.Utf8)
	var big5BytesPrefix []byte
	var big5BytesPostfix []byte
	if len(big5Bytes) >= 1 {
		big5BytesPrefix = big5Bytes[:len(big5Bytes)-1]
		big5BytesPostfix = big5Bytes[len(big5Bytes)-1:]
	}
	theBytes = append(theBytes, big5BytesPrefix...)

	colorBytes = r.Color1.BytesWithPreColor(&r.Color0)
	theBytes = append(theBytes, colorBytes...)

	theBytes = append(theBytes, big5BytesPostfix...)
	if isAddCR {
		theBytes = append(theBytes, '\r')
	}

	if len(theBytes) == 0 {
		theBytes = nil
	}

	r.DBCS = theBytes
}

func Utf8ToBig5(utf8 string) (big5 []byte) {
	utf8Bytes := []byte(utf8)
	estimatedBig5Sz := len(utf8Bytes)
	big5 = make([]byte, 0, estimatedBig5Sz)

	for p_utf8 := utf8Bytes; len(p_utf8) > 0; {
		if p_utf8[0] < 0x80 {
			big5 = append(big5, p_utf8[0])
			p_utf8 = p_utf8[1:]
		} else if len(p_utf8) >= 2 && (p_utf8[0]&0xe0) == 0xc0 {
			eachBig5, ok := utf8ToBig5[string(p_utf8[:2])]
			if !ok {
				eachBig5 = []byte{0xff, 0xfd}
			}
			big5 = append(big5, eachBig5...)
			p_utf8 = p_utf8[2:]
		} else if len(p_utf8) >= 3 && (p_utf8[0]&0xf0) == 0xe0 {
			eachBig5, ok := utf8ToBig5[string(p_utf8[:3])]
			if !ok {
				eachBig5 = []byte{0xff, 0xfd}
			}
			big5 = append(big5, eachBig5...)
			p_utf8 = p_utf8[3:]
		}
	}
	if len(big5) == 0 {
		return nil
	}

	return big5
}

// StringsSplitAsRune
//
// sep: single-rune string
func StringsSplitAsRune(str string, sep string) (strs []string) {
	theRune := []rune(str)
	sepRunes := []rune(sep)
	sepRune := sepRunes[0]
	count := 1
	for _, each := range theRune {
		if each == sepRune {
			count++
		}
	}

	strs = make([]string, count)
	startIdx := 0
	strIdx := 0
	for idx, each := range theRune {
		if each == sepRune {
			strs[strIdx] = string(theRune[startIdx:idx])
			strIdx++
			startIdx = idx + 1
		}
	}
	// the last
	strs[strIdx] = string(theRune[startIdx:])

	return strs
}
