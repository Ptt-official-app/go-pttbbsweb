package dbcs

import (
	"strings"
)

//splitArticleSignatureCommentsDBCSStr
//
//given content, try to split as article / bbs-signature / comments
//expecting articleDBCS / signatureDBCS / comments as the sub-string of content.
//
//在 edit 之後, 整個 content 有可能會被擾亂.
//我們只能盡可能的根據得到的資訊來決定是否是屬於 signature 或是 comments.
//
//rules:
//1. 整個 content 只會被分成 3 種可能: article / bbs-signature / comments
//2. bbs-signature 辨識方式: 最後 1 個 ※  發信站:
//3. comments 辨識方式: 符合顏色的推/噓/→
//4. 如果有 bbs-signature, comments 一定是在 bbs-signature 之後.
//5. 一旦出現 comments, 之後的 content 都會是 comments 或 reply.
//
//Implementation:
//1. 嘗試拿到 signature 的 idx.
//2. 如果有 signature idx: 嘗試將 signature 之後的分出 comments, return article / signature, comments
//3. 否則: 嘗試將 content 分出 comments, return article / nil, comments
func splitArticleSignatureCommentsDBCSStr(content string) (articleDBCS string, signatureDBCS string, comments string) {
	// 1. get idx with signature
	idxes := tryGetSimpleSignatureIdxesStr(content)

	// start with largest idx
	reverseList(idxes)

	for _, idx := range idxes {
		// 2. validate signature idx and nBytes of signature
		isValid, nBytes := isValidSimpleSignatureStr(content, idx)
		if !isValid {
			continue
		}

		commentIdx := idx + nBytes
		comments = content[commentIdx:]
		if len(comments) == 0 {
			comments = ""
		}
		return content[:idx], content[idx:commentIdx], comments
	}

	// 3. get content with comments, but no signatures.
	idx := MatchCommentStr(content)
	if idx == -1 {
		return content, "", ""
	}

	comments = content[idx:]
	if len(comments) == 0 {
		comments = ""
	}

	return content[:idx], "", comments
}

const MATCH_SIGNATURE_INIT_STR = "\n※ 發信站:" //\n※  發信站: (no \n-- in forward (轉))

func tryGetSimpleSignatureIdxesStr(content string) []int {
	idxes := make([]int, 0, 5) // pre-allocated possible 5 candidate idxes.

	p_content := content
	offsetIdx := 0
	for idx := strings.Index(p_content, MATCH_SIGNATURE_INIT_STR); len(p_content) > 0 && (idx != -1); idx = strings.Index(p_content, MATCH_SIGNATURE_INIT_STR) {
		theIdx := offsetIdx + idx
		if idx >= 5 && p_content[idx-5] == '\r' && p_content[idx-4] == '\n' && p_content[idx-3] == '-' && p_content[idx-2] == '-' && p_content[idx-1] == '\r' {
			theIdx -= 5
		} else if idx >= 4 && p_content[idx-4] == '\n' && p_content[idx-3] == '-' && p_content[idx-2] == '-' && p_content[idx-1] == '\r' {
			theIdx -= 4
		} else if idx >= 3 && p_content[idx-3] == '\n' && p_content[idx-2] == '-' && p_content[idx-1] == '-' {
			theIdx -= 3
		} else if idx >= 1 && p_content[idx-1] == '\r' {
			theIdx -= 1 // nolint // consistent with programming pattern
		}

		idxes = append(idxes, theIdx)

		p_content = p_content[idx+len(MATCH_SIGNATURE_INIT_STR):]
		offsetIdx += idx + len(MATCH_SIGNATURE_INIT_STR)
	}

	return idxes
}

func isValidSimpleSignatureStr(content string, idx int) (isValid bool, nBytes int) {
	p_content := content[idx:]

	return isSimpleSignatureWithFromStr(p_content)
}

var (
	MATCH_SIGNATURE_FROM_STR = "), 來自: " //), 來自:

	MATCH_SIGNATURE_FROM_OLD_STR = "◆ From: " //◆ From:

	MATCH_SIGNATURE_FORWARD_STR = "※ 轉錄者: "

	MATCH_SIGNATURE_URL_STR = "※ 文章網址: "
)

func isSimpleSignatureWithFromStr(p_content string) (isValid bool, nBytes int) {
	nBytes = 0
	if len(p_content) >= 6 && p_content[0] == '\r' && p_content[1] == '\n' && p_content[2] == '-' && p_content[3] == '-' && p_content[4] == '\r' && p_content[5] == '\n' { //\r\n--\r\n
		p_content = p_content[6:]
		nBytes += 6
	} else if len(p_content) >= 5 && p_content[0] == '\n' && p_content[1] == '-' && p_content[2] == '-' && p_content[3] == '\r' && p_content[4] == '\n' { //\n--\r\n
		p_content = p_content[5:]
		nBytes += 5
	} else if len(p_content) >= 4 && p_content[0] == '\n' && p_content[1] == '-' && p_content[2] == '-' && p_content[3] == '\n' { //\n -- \n
		p_content = p_content[4:]
		nBytes += 4
	} else if len(p_content) >= 2 && p_content[0] == '\r' && p_content[1] == '\n' { //\r\n
		p_content = p_content[2:]
		nBytes += 2
	} else if len(p_content) >= 1 && p_content[0] == '\n' { //\n
		p_content = p_content[1:]
		nBytes += 1 // nolint // consistent with programming pattern
	}

	// check 來自 in the 1st line.
	idxNewLine := strings.Index(p_content, "\n")
	isNewLine := true
	if idxNewLine == -1 {
		isNewLine = false
		idxNewLine = len(p_content)
	}
	line := p_content[:idxNewLine]

	if strings.Contains(line, MATCH_SIGNATURE_FROM_STR) {
		nBytes += idxNewLine
		if isNewLine {
			nBytes++
			nBytes += simpleSignatureWithFromRangeStr(p_content[idxNewLine+1:])
		}
		return true, nBytes
	}

	if !isNewLine {
		return false, 0
	}

	// check From in the 2nd line.
	nBytes += idxNewLine + 1
	p_content = p_content[(idxNewLine + 1):]

	idxNewLine = strings.Index(p_content, "\n")
	line = p_content // unable to find '\n'
	if idxNewLine != -1 {
		line = p_content[:idxNewLine]
	}

	if strings.HasPrefix(line, MATCH_SIGNATURE_FROM_OLD_STR) {
		nBytes += idxNewLine + 1
		return true, nBytes
	}

	// check 轉錄者 in the 2nd line.
	if strings.HasPrefix(line, MATCH_SIGNATURE_FORWARD_STR) {
		nBytes += idxNewLine + 1
		return true, nBytes
	}

	return false, 0
}

//simpleSignatureWithFromRange
//
//※  文章網址: https://www.ptt.cc/bbs
func simpleSignatureWithFromRangeStr(content string) int {
	if !strings.HasPrefix(content, MATCH_SIGNATURE_URL_STR) {
		return 0
	}
	idx := strings.Index(content, "\n")
	if idx == -1 {
		return len(content)
	}
	return idx + 1
}
