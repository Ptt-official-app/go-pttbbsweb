package dbcs

import (
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	log "github.com/sirupsen/logrus"
)

//ParseContent
//
//Assume:
//1. the content is with chars >= 32 and '\x1b', '\r', \n'
//2. the timestamp of the 1st-comments (around 10 comments, including the last-same-min comments) are within 1-year of the createTime.
//3. the timestamp of the rest of the comments are able to reverse-inferred from mtime. compared as stored as nano-ts.
//4. assuming no more than 60000 comments  (60 x 1000) in 1 minute.
func ParseContent(contentBytes []byte, origContentMD5 string) (content [][]*types.Rune, contentMD5 string, ip string, host string, bbs string, signatureMD5 string, signatureDBCS []byte, commentsDBCS []byte) {

	contentDBCS, signatureDBCS, commentsDBCS := splitArticleSignatureCommentsDBCS(contentBytes)

	contentMD5 = md5sum(contentDBCS)

	if contentMD5 == origContentMD5 {
		return nil, "", "", "", "", "", nil, commentsDBCS
	}

	contentBig5 := dbcsToBig5(contentDBCS)
	content = big5ToUtf8(contentBig5)

	signatureBig5 := dbcsToBig5(signatureDBCS)
	signatureUtf8 := big5ToUtf8(signatureBig5)
	ip, host, bbs = parseSignatureIPHostBBS(signatureUtf8)
	signatureMD5 = md5sum(signatureDBCS)

	return content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, commentsDBCS
}

func parseSignatureIPHostBBS(signatureUtf8 [][]*types.Rune) (ip string, host string, bbs string) {
	if len(signatureUtf8) == 0 { //unable to parse signature.
		return "", "", ""
	}

	//check \n
	zerothLine := signatureUtf8[0]
	if len(zerothLine) == 0 {
		signatureUtf8 = signatureUtf8[1:]
	} else if len(zerothLine) == 1 && zerothLine[0].Utf8 == "" {
		signatureUtf8 = signatureUtf8[1:]
	}
	if len(signatureUtf8) == 0 {
		return "", "", ""
	}

	//check --
	zerothLine = signatureUtf8[0]
	if zerothLine[0].Utf8 == "--" {
		signatureUtf8 = signatureUtf8[1:]
	}

	if len(signatureUtf8) == 0 {
		return "", "", ""
	}

	//check the 0th line
	zerothLine = signatureUtf8[0]
	if len(zerothLine) == 0 { //unable to get rune
		return "", "", ""
	}
	//expecting only 1 rune
	if len(zerothLine) > 1 {
		log.Warning("dbcs.parseIPHostBBS: multiple colors")
	}
	zerothRune := zerothLine[0].Utf8
	prefix := "※ 發信站: "
	if !strings.HasPrefix(zerothRune, prefix) { //unable to get BBS
		return "", "", ""
	}

	zerothRune = zerothRune[len(prefix):]
	theIdx := strings.Index(zerothRune, ")")
	if theIdx == -1 { //unable to get BBS
		return "", "", zerothRune
	}

	//got bbs, set next zerothRune
	bbs, zerothRune = zerothRune[:theIdx+1], zerothRune[(theIdx+1):]
	bbs = strings.TrimSpace(bbs)

	from := "來自: "
	theIdx = strings.Index(zerothRune, from)
	if theIdx != -1 { // found origin.
		ip, host = parseIPHostBBS0(zerothRune[theIdx+len(from):])
		return ip, host, bbs
	}

	//2. check the 1st line
	if len(signatureUtf8) < 2 { //unable to find more lines
		return "", "", bbs
	}

	firstLine := signatureUtf8[1]
	//2.1. expecting only 1 rune
	if len(firstLine) == 0 { //unable to get rune
		return "", "", bbs
	}
	if len(firstLine) > 1 {
		log.Warning("dbcs.parseIPHostBBS: multiple colors")
	}
	firstRune := firstLine[0].Utf8
	//2.1.1. forward
	forward := "※ 轉錄者: "
	if strings.HasPrefix(firstRune, forward) { //found the forward
		forwardStr := firstRune[len(forward):]
		theIdx := strings.Index(forwardStr, "(")
		if theIdx == -1 {
			return "", "", bbs
		}
		theIdx2 := strings.Index(forwardStr, ")")
		if theIdx2 == -1 {
			return "", "", bbs
		}
		ipList := strings.Split(forwardStr[(theIdx+1):theIdx2], " ")
		if len(ipList) == 1 {
			return ipList[0], "", bbs
		}
		return ipList[0], ipList[1], bbs
	}

	//2.1.2. old-form
	from = "◆ From: "

	if !strings.HasPrefix(firstRune, from) { //unable to find the from
		return "", "", bbs
	}
	return firstRune[len(from):], "", bbs
}

func parseIPHostBBS0(str string) (ip string, host string) {
	theIdx := strings.Index(str, " (")
	if theIdx == -1 { //unable to find host
		return strings.TrimSpace(str), ""
	}

	ip = str[:theIdx]

	theIdx2 := strings.Index(str, ")")
	if theIdx2 == -1 {
		host = str[(theIdx + 2):]
	} else {
		host = str[(theIdx + 2):theIdx2]
	}

	return strings.TrimSpace(ip), strings.TrimSpace(host)
}
