package dbcs

import (
	"bytes"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

//ParseComments
//
//有可能 reply-edit-info (編輯) 不在 commentsDBCS 裡
//但是會在 allCommentsDBCS 裡 (firstComments)
//只考慮:
//1. appropriately split comments.
//2. 對於每個 comment 裡的 DBCS Parse 成 Utf8.
//3. type / IP / Host / MD5 / TheDate
//
//不考慮:
//1. boardID / articleID / commentID.
//2. createTime / firstCreateTime / InferredCreateTime / AddCreateTime (除了編輯以外)
//
//1. 根據 '\n' 估計 nComments
//2. 找出 pre-comment reply.
//3. 對於每個 comment-leading newline for-loop:
func ParseComments(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	ownerID bbs.UUserID,
	lastTimeNanoTS types.NanoTS,
	commentsDBCS []byte,
	allCommentsDBCS []byte,
	updateNanoTS types.NanoTS,
	isFirstComments bool,
) (

	comments []*schema.Comment,
	newLastTimeNanoTS types.NanoTS) {
	if len(commentsDBCS) == 0 {
		return nil, lastTimeNanoTS
	}

	// estimate nComments
	nEstimatedComments := bytes.Count(commentsDBCS, []byte{'\n'})

	comments = make([]*schema.Comment, 0, nEstimatedComments)

	p_commentsDBCS := commentsDBCS
	p_allCommentsDBCS := allCommentsDBCS

	lastTime := lastTimeNanoTS.ToTime()

	var comment *schema.Comment

	nextCommentIdx := MatchComment(p_commentsDBCS)
	if nextCommentIdx > 0 {
		replyDBCS := p_commentsDBCS[:nextCommentIdx]
		reply := parseReply(replyDBCS, p_allCommentsDBCS)
		if reply != nil {
			comments = append(comments, reply)
		}

		p_allCommentsDBCS = p_allCommentsDBCS[len(replyDBCS):]
		p_commentsDBCS = p_commentsDBCS[len(replyDBCS):]
	}

	for idxNewLine := bytes.Index(p_commentsDBCS, []byte{'\n'}); len(p_commentsDBCS) > 0 && idxNewLine != -1; idxNewLine = bytes.Index(p_commentsDBCS, []byte{'\n'}) {
		commentDBCS := p_commentsDBCS[:idxNewLine]
		comment, lastTime = parseComment(bboardID, articleID, lastTime, commentDBCS, updateNanoTS, isFirstComments)
		comments = append(comments, comment)

		p_commentsDBCS = p_commentsDBCS[idxNewLine:] // with '\n'
		p_allCommentsDBCS = p_allCommentsDBCS[idxNewLine:]

		nextCommentIdx := MatchComment(p_commentsDBCS)

		if nextCommentIdx == -1 { // no more comments
			p_commentsDBCS = p_commentsDBCS[1:] //step forward '\n'
			p_allCommentsDBCS = p_allCommentsDBCS[1:]
			if len(p_commentsDBCS) > 0 {
				replyDBCS := p_commentsDBCS
				reply := parseReply(replyDBCS, p_allCommentsDBCS)
				if reply != nil {
					comments = append(comments, reply)
				}

				p_allCommentsDBCS = p_allCommentsDBCS[len(p_commentsDBCS):]
				p_commentsDBCS = nil
			}
			break
		}

		if nextCommentIdx > 1 { // p_commentsDBCS[0] is '\n', get reply from p_commentsDBCS[1:]
			replyDBCS := p_commentsDBCS[1:nextCommentIdx]

			reply := parseReply(replyDBCS, p_allCommentsDBCS[1:])
			if reply != nil {
				comments = append(comments, reply)
			}
		}

		p_commentsDBCS = p_commentsDBCS[nextCommentIdx:]
		p_allCommentsDBCS = p_allCommentsDBCS[nextCommentIdx:]
	}

	if len(p_commentsDBCS) > 0 { // last comment without '\n'
		comment, lastTime = parseComment(bboardID, articleID, lastTime, p_commentsDBCS, updateNanoTS, isFirstComments)
		comments = append(comments, comment)
	}

	return comments, types.TimeToNanoTS(lastTime)
}

func parseComment(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	lastTime time.Time,
	commentDBCS []byte,
	updateNanoTS types.NanoTS,
	isFirstComments bool) (

	comment *schema.Comment,
	newTime time.Time) {

	theType, p_commentDBCS := parseCommentType(commentDBCS)
	ownerID, p_commentDBCS := parseCommentOwnerID(p_commentDBCS)
	contentDBCS, p_commentDBCS := parseCommentContent(p_commentDBCS)
	contentBig5 := dbcsToBig5(contentDBCS) //the last 11 chars are the dates
	contentUtf8 := big5ToUtf8(contentBig5)
	ip, createTime := parseCommentIPCreateTime(lastTime, p_commentDBCS)
	commentMD5 := md5sum(commentDBCS)

	createNanoTS := types.TimeToNanoTS(createTime)
	commentID := types.ToCommentID(createNanoTS, commentMD5)

	comment = &schema.Comment{
		BBoardID:     bboardID,
		ArticleID:    articleID,
		CommentID:    commentID,
		TheType:      theType,
		CreateTime:   createNanoTS,
		Owner:        ownerID,
		Content:      contentUtf8,
		IP:           ip,
		MD5:          commentMD5,
		UpdateNanoTS: updateNanoTS,
		DBCS:         commentDBCS,
	}
	comment.CleanComment()

	return comment, createTime
}

func parseCommentContent(p_commmentDBCS []byte) (contentDBCS []byte, nextCommentDBCS []byte) {
	if len(p_commmentDBCS) == 0 {
		return nil, nil
	}

	idx := bytes.Index(p_commmentDBCS, []byte{'\x1b'})
	if idx == -1 {
		return p_commmentDBCS[1:], nil
	}

	contentDBCS, nextCommentDBCS = p_commmentDBCS[1:idx], p_commmentDBCS[idx:]
	if len(contentDBCS) > 0 && contentDBCS[0] == ' ' {
		contentDBCS = contentDBCS[1:]
	}
	if len(contentDBCS) == 0 {
		contentDBCS = nil
	}
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = nil
	}
	idx = bytes.Index(nextCommentDBCS, []byte{'m'})
	if idx == -1 {
		nextCommentDBCS = nil
	}
	nextCommentDBCS = nextCommentDBCS[idx+1:]
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = nil
	}

	return contentDBCS, nextCommentDBCS
}

func parseCommentType(p_commmentDBCS []byte) (theType types.CommentType, nextCommentDBCS []byte) {
	if len(p_commmentDBCS) < 15 {
		return types.COMMENT_TYPE_COMMENT, nil
	}
	nextCommentDBCS = p_commmentDBCS[15:]
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = nil
	}

	if bytes.HasPrefix(p_commmentDBCS, MATCH_COMMENT_RECOMMEND_BYTES[1:]) {
		return types.COMMENT_TYPE_RECOMMEND, nextCommentDBCS
	} else if bytes.HasPrefix(p_commmentDBCS, MATCH_COMMENT_BOO_BYTES[1:]) {
		return types.COMMENT_TYPE_BOO, nextCommentDBCS
	} else if bytes.HasPrefix(p_commmentDBCS, MATCH_COMMENT_ARROW_BYTES[1:]) {
		return types.COMMENT_TYPE_COMMENT, nextCommentDBCS
	} else {
		return types.COMMENT_TYPE_COMMENT, nextCommentDBCS
	}
}

func parseCommentOwnerID(p_commmentDBCS []byte) (ownerID bbs.UUserID, nextCommentDBCS []byte) {
	if len(p_commmentDBCS) == 0 {
		return "", nil
	}
	theIdx := bytes.Index(p_commmentDBCS, []byte{'\x1b'})
	if theIdx == -1 {
		return bbs.UUserID(""), nil
	}

	ownerID = bbs.UUserID(string(p_commmentDBCS[:theIdx]))
	if len(p_commmentDBCS) <= theIdx+8 {
		return ownerID, nil
	}
	nextCommentDBCS = p_commmentDBCS[theIdx+8:]

	return ownerID, nextCommentDBCS
}

//parseCommentIPCreateTime
//
//Already separate the data by color.
//There are only ip/create-time information in p_commentDBCS.
func parseCommentIPCreateTime(lastTime time.Time, p_commentDBCS []byte) (ip string, createTime time.Time) {
	if len(p_commentDBCS) == 0 {
		return "", lastTime.Add(COMMENT_STEP_DURATION)
	}
	theIdx := bytes.Index(p_commentDBCS, []byte("\xb1\xc0")) //推
	if theIdx != -1 {                                        //old
		postfix := strings.TrimSpace(types.Big5ToUtf8(p_commentDBCS[theIdx+2:]))
		postfixList := strings.Split(postfix, " ")
		if len(postfixList) != 2 { //unable to parse. return createTime + 10-millisecond
			return "", lastTime.Add(COMMENT_STEP_DURATION)
		}
		ip = postfixList[0]
		dateStr := postfixList[1]
		theTime, err := types.DateStrToTime(dateStr)
		if err != nil {
			return ip, lastTime.Add(COMMENT_STEP_DURATION)
		}

		if theTime.Month() == createTime.Month() && theTime.Day() == createTime.Day() {
			return ip, lastTime.Add(COMMENT_STEP_DURATION)
		} else {
			newTime := types.NewDateTime(lastTime.Year(), theTime.Month(), theTime.Day(), 0, 0, 0)
			if newTime.Before(lastTime) {
				newTime = types.NewDateTime(lastTime.Year()+1, theTime.Month(), theTime.Day(), 0, 0, 0)
			}

			return ip, newTime
		}
	}

	//new: MM/DD HH:mm
	dateTimeStr := strings.TrimSpace(string(p_commentDBCS))
	theTime, err := types.DateTimeStrToTime(dateTimeStr)
	if err != nil {
		return "", lastTime.Add(COMMENT_STEP_DURATION)
	}

	if theTime.Month() == lastTime.Month() && theTime.Day() == lastTime.Day() && theTime.Hour() == lastTime.Hour() && theTime.Minute() == lastTime.Minute() {
		return "", lastTime.Add(COMMENT_STEP_DURATION)
	}

	newDateTime := types.NewDateTime(lastTime.Year(), theTime.Month(), theTime.Day(), theTime.Hour(), theTime.Minute(), 0)
	if newDateTime.Before(lastTime) {
		newDateTime = types.NewDateTime(lastTime.Year()+1, theTime.Month(), theTime.Day(), theTime.Hour(), theTime.Minute(), 0)

	}

	return "", newDateTime
}

//parseReply
//
//只考慮parse
func parseReply(
	replyDBCS []byte,
	editDBCS []byte) (

	reply *schema.Comment) {

	if len(replyDBCS) == 0 {
		return nil
	}

	origReplyDBCS := replyDBCS
	if replyDBCS[len(replyDBCS)-1] == '\n' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil
	}
	if replyDBCS[len(replyDBCS)-1] == '\r' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil
	}

	replyMD5 := md5sum(origReplyDBCS)
	replyBig5 := dbcsToBig5(replyDBCS)
	replyUtf8 := big5ToUtf8(replyBig5)

	editUserID, editNanoTS, editIP, editHost := parseReplyUserIPHost(editDBCS)

	reply = &schema.Comment{
		TheType: types.COMMENT_TYPE_REPLY,
		Owner:   editUserID,
		Content: replyUtf8,
		IP:      editIP,
		Host:    editHost,
		MD5:     replyMD5,

		EditNanoTS: editNanoTS,

		DBCS: origReplyDBCS,
	}

	reply.CleanReply()
	if len(reply.Content) == 0 {
		return nil
	}

	return reply
}

//parseReplyUserIPHost
//※ 編輯: abcd (1.2.3.4 臺灣), 03/21/2021 03:04:47
//
//1. 找到 EDIT_PREFIX
//2. 找到 (, 設定 userID
//3. 找到 ), 設定 ip/host
//4. 設定時間.
func parseReplyUserIPHost(editDBCS []byte) (editUserID bbs.UUserID, editNanoTS types.NanoTS, editIP string, editHost string) {

	//1.  get EDIT_PREFIX
	p_editDBCS := editDBCS
	theIdx := bytes.Index(p_editDBCS, MATCH_COMMENT_EDIT_BYTES[1:])
	if theIdx == -1 {
		return "", 0, "", ""
	}

	//2. get (
	p_editDBCS = p_editDBCS[theIdx+len(MATCH_COMMENT_EDIT_BYTES)-1:]

	theIdx = bytes.Index(p_editDBCS, []byte{'('})
	if theIdx == -1 {
		return "", 0, "", ""
	}
	editUserID = bbs.UUserID(p_editDBCS[:theIdx-1])

	//3. get )
	p_editDBCS = p_editDBCS[theIdx+1:]

	theIdx = bytes.Index(p_editDBCS, []byte{')'})
	if theIdx == -1 {
		return "", 0, "", ""
	}
	ipHost := types.Big5ToUtf8(p_editDBCS[:theIdx])

	ipHostList := strings.Split(ipHost, " ")
	if len(ipHostList) == 1 {
		editIP = ipHostList[0]
	} else {
		editIP = ipHostList[0]
		editHost = ipHostList[1]
	}

	//4. get time.
	p_editDBCS = p_editDBCS[theIdx:]

	theIdx = bytes.Index(p_editDBCS, []byte(", "))
	p_editDBCS = p_editDBCS[theIdx+2:]

	theTime, err := types.DateYearTimeStrToTime(string(p_editDBCS[:19]))
	if err != nil {
		return "", 0, "", ""
	}

	return editUserID, types.TimeToNanoTS(theTime), editIP, editHost
}
