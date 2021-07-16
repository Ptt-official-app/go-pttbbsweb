package dbcs

import (
	"bytes"
)

//splitArticleSignatureCommentsDBCS
//
//given content, try to split as article / bbs-signature / comments
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
func splitArticleSignatureCommentsDBCS(content []byte) (articleDBCS []byte, signatureDBCS []byte, comments []byte) {
	// 1. get idx with signature
	idxes := tryGetSimpleSignatureIdxes(content)

	// start with largest idx
	reverseList(idxes)

	for _, idx := range idxes {
		// 2. validate signature idx and nBytes of signature
		isValid, nBytes := isValidSimpleSignaure(content, idx)
		if !isValid {
			continue
		}

		commentIdx := idx + nBytes
		comments = content[commentIdx:]
		if len(comments) == 0 {
			comments = nil
		}
		return content[:idx], content[idx:commentIdx], comments
	}

	// 3. get content with comments, but no signatures.
	idx := MatchComment(content)
	if idx == -1 {
		return content, nil, nil
	}

	comments = content[idx:]
	if len(comments) == 0 {
		comments = nil
	}

	return content[:idx], nil, comments
}

var MATCH_SIGNATURE_INIT = []byte{ //\n※  發信站: (no \n-- in forward (轉))
	0x0a, 0xa1, 0xb0, 0x20, 0xb5,
	0x6f, 0xab, 0x48, 0xaf, 0xb8, 0x3a, 0x20,
}

func tryGetSimpleSignatureIdxes(content []byte) []int {
	idxes := make([]int, 0, 5) // pre-allocated possible 5 candidate idxes.

	p_content := content
	offsetIdx := 0
	for idx := bytes.Index(p_content, MATCH_SIGNATURE_INIT); len(p_content) > 0 && (idx != -1); {
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

		p_content = p_content[idx+len(MATCH_SIGNATURE_INIT):]
		offsetIdx += idx + len(MATCH_SIGNATURE_INIT)
		idx = bytes.Index(p_content, MATCH_SIGNATURE_INIT)
	}

	return idxes
}

func isValidSimpleSignaure(content []byte, idx int) (isValid bool, nBytes int) {
	p_content := content[idx:]

	return isSimpleSignatureWithFrom(p_content)
}

var (
	MATCH_SIGNATURE_FROM = []byte{ //), 來自:
		0x29, 0x2c, 0x20, 0xa8, 0xd3, 0xa6, 0xdb, 0x3a, 0x20,
	}

	MATCH_SIGNATURE_FROM_OLD = []byte{ //◆  From:
		0xa1, 0xbb, 0x20, 0x46, 0x72, 0x6f, 0x6d, 0x3a, 0x20,
	}

	MATCH_SIGNATURE_FORWARD = []byte{ //※ 轉錄者:
		0xa1, 0xb0, 0x20, 0xc2, 0xe0, 0xbf, 0xfd, 0xaa, 0xcc, 0x3a, 0x20,
	}

	MATCH_SIGNATURE_URL = []byte{ //※ 文章網址:
		0xa1, 0xb0, 0x20, 0xa4, 0xe5, 0xb3, 0xb9, 0xba, 0xf4, 0xa7, 0x7d, 0x3a, 0x20,
	}
)

func isSimpleSignatureWithFrom(p_content []byte) (isValid bool, nBytes int) {
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
	idxNewLine := bytes.Index(p_content, []byte{'\n'})
	line := p_content // unable to find '\n'
	if idxNewLine != -1 {
		line = p_content[:idxNewLine]
	}

	if bytes.Contains(line, MATCH_SIGNATURE_FROM) {
		nBytes += idxNewLine + 1
		nBytes += simpleSignatureWithFromRange(p_content[idxNewLine+1:])
		return true, nBytes
	}

	// check From in the 2nd line.
	nBytes += idxNewLine + 1
	p_content = p_content[(idxNewLine + 1):]

	idxNewLine = bytes.Index(p_content, []byte{'\n'})
	line = p_content // unable to find '\n'
	if idxNewLine != -1 {
		line = p_content[:idxNewLine]
	}

	if bytes.HasPrefix(line, MATCH_SIGNATURE_FROM_OLD) {
		nBytes += idxNewLine + 1
		return true, nBytes
	}

	// check 轉錄者 in the 2nd line.
	if bytes.HasPrefix(line, MATCH_SIGNATURE_FORWARD) {
		nBytes += idxNewLine + 1
		return true, nBytes
	}

	return false, 0
}

//simpleSignatureWithFromRange
//
//※  文章網址: https://www.ptt.cc/bbs
func simpleSignatureWithFromRange(content []byte) int {
	if !bytes.HasPrefix(content, MATCH_SIGNATURE_URL) {
		return 0
	}
	idx := bytes.Index(content, []byte{'\n'})
	return idx + 1
}
