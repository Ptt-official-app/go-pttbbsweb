package dbcs

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

// ParseContentStr
//
// Assume:
// 1. the content is with chars >= 32 and '\x1b', '\r', \n'
// 2. the timestamp of the 1st-comments (around 10 comments, including the last-same-min comments) are within 1-year of the createTime.
// 3. the timestamp of the rest of the comments are able to reverse-inferred from mtime. compared as stored as nano-ts.
// 4. assuming no more than 60000 comments  (60 x 1000) in 1 minute.
func ParseContentStr(contentStr string, origContentMD5 string, isSplit bool) (content [][]*types.Rune, contentPrefix [][]*types.Rune, contentMD5 string, ip string, host string, bbs string, signatureMD5 string, signatureDBCS string, commentsDBCS string) {
	// remove \r
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")

	var contentDBCS string
	if isSplit {
		contentDBCS, signatureDBCS, commentsDBCS = splitArticleSignatureCommentsDBCSStr(contentStr)
	} else {
		contentDBCS = contentStr
	}

	contentMD5 = Md5sum([]byte(contentDBCS))

	if contentMD5 == origContentMD5 {
		return nil, nil, "", "", "", "", "", "", commentsDBCS
	}

	content = dbcsToUtf8(contentDBCS)

	if isSplit {
		signatureUtf8 := dbcsToUtf8(signatureDBCS)
		ip, host, bbs = parseSignatureIPHostBBS(signatureUtf8)
		signatureMD5 = Md5sum([]byte(signatureDBCS))

		content, contentPrefix = parseContentPrefix(content)
	}

	return content, contentPrefix, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, commentsDBCS
}
