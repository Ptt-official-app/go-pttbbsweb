package dbcs

import (
	"bytes"
)

var (
	MATCH_COMMENT_RECOMMEND_BYTES = []byte{ //推
		//\x1b[1;37m\xb1\xc0 \x1b[33mabcd   \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:07
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x37, 0x6d,
		0xb1, 0xc0, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	MATCH_COMMENT_BOO_BYTES = []byte{ //噓
		//\x1b[1;31m\xbcN \x1b[33mabcd  \x1b[m\x1b[33m: \xb0\xa3\xabD\xb9\xef\xa4\xe2\xa4]....          \x1b[m 03/18 18:13
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xbc, 0x4e, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	MATCH_COMMENT_ARROW_BYTES = []byte{ //→
		//\x1b[1;31m\xa1\xf7 \x1b[33mabcd     \x1b[m\x1b[33m: 7777777777                                       \x1b[m 03/18 12:07
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xa1, 0xf7, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	//※ 編輯: abcd (1.2.3.4 臺灣), 03/21/2021 03:04:47
	//\xa1\xb0 \xbds\xbf\xe8: abcd (1.2.3.4 \xbbO\xc6W), 03/18/2021 12:07:22
	MATCH_COMMENT_EDIT_BYTES = []byte("\xa1\xb0 \xbds\xbf\xe8: ")

	//※ abcde:轉錄至看板 SYSOP
	//\xa1\xb0 \x1b[1;32mabcd\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Mavericks\x1b[m                               03/18 12:07
	MATCH_COMMENT_FORWARD_BYTES = []byte{ //\x1b[0;32m:\xc2
		0x1b, 0x5b, 0x30, 0x3b, 0x33, 0x32, 0x6d, 0x3a, 0xc2, 0xe0, 0xbf, 0xfd, 0xa6, 0xdc, 0xac, 0xdd, 0xaa, 0x4f, 0x20,
	}
)

//MatchComment
//
//TODO: record the idxes of each condition, rematch only the condition with the smallest idx.
func MatchComment(content []byte) int {
	theIdx := -1

	idxRecommend := bytes.Index(content, MATCH_COMMENT_RECOMMEND_BYTES)
	if idxRecommend != -1 && (idxRecommend == 0 || content[idxRecommend-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxRecommend)
	}

	idxBoo := bytes.Index(content, MATCH_COMMENT_BOO_BYTES)
	if idxBoo != -1 && (idxBoo == 0 || content[idxBoo-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxBoo)
	}
	idxArrow := bytes.Index(content, MATCH_COMMENT_ARROW_BYTES)
	if idxArrow != -1 && (idxArrow == 0 || content[idxArrow-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxArrow)
	}
	idxEdit := bytes.Index(content, MATCH_COMMENT_EDIT_BYTES)
	if idxEdit != -1 && (idxEdit == 0 || content[idxEdit-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxEdit)
	}
	idxForward := matchCommentForward(content)
	if idxForward != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxForward)
	}

	if theIdx == -1 {
		return theIdx
	}

	return theIdx //do not include the leading \n
}

func matchCommentIntegratedIdx(theIdx int, idx int) int {
	if theIdx == -1 {
		return idx
	}
	if theIdx < idx {
		return theIdx
	}
	return idx
}

func matchCommentForward(content []byte) (idx int) {
	for len(content) > 0 {
		idx = bytes.Index(content, MATCH_COMMENT_FORWARD_BYTES)
		if idx == -1 {
			return -1
		}

		skipIdx := idx
		if idx < 3 {
			content = content[skipIdx+len(MATCH_COMMENT_FORWARD_BYTES):]
			continue
		}

		idx--
		for ; idx >= 2 && content[idx] != ' '; idx-- {
		}
		if idx == 2 && content[idx-2] == '\xa1' && content[idx-1] == '\xb0' && content[idx] == ' ' {
			return idx - 2
		}
		if content[idx-3] == '\n' && content[idx-2] == '\xa1' && content[idx-1] == '\xb0' && content[idx] == ' ' {
			return idx - 2
		}

		content = content[skipIdx+len(MATCH_COMMENT_FORWARD_BYTES):]
	}

	return -1
}
