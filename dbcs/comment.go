package dbcs

import (
	"bytes"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

//ParseFirstComments
//
//Check with origFirstCommentsMD5, if exists, return nil and requires getting firstComments and lastTime from db.
func ParseFirstComments(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	ownerID bbs.UUserID,
	articleCreateTime types.NanoTS,
	commentsDBCS []byte,
	origFirstCommentsMD5 string,
	origFirstCommentsLastTime types.NanoTS,
	updateNanoTS types.NanoTS) (

	firstComments []*schema.Comment,
	firstCommentsMD5 string,
	firstCommentsLastTime types.NanoTS,
	theRestComments []byte) {

	firstCommentsDBCS, theRestComments := splitFirstComments(commentsDBCS)

	//check md5
	firstCommentsMD5 = md5sum(firstCommentsDBCS)
	if firstCommentsMD5 == origFirstCommentsMD5 {
		return nil, origFirstCommentsMD5, origFirstCommentsLastTime, theRestComments
	}

	firstComments, firstCommentsLastTime = ParseComments(bboardID, articleID, ownerID, articleCreateTime, firstCommentsDBCS, commentsDBCS, updateNanoTS, true)

	return firstComments, firstCommentsMD5, firstCommentsLastTime, theRestComments
}

func splitFirstComments(commentsDBCS []byte) (firstCommentsDBCS []byte, theRestComments []byte) {
	p_commentsDBCS := commentsDBCS

	nComments := 0
	nBytes := 0
	for idx := bytes.Index(p_commentsDBCS, []byte{'\n'}); len(p_commentsDBCS) > 0 && idx != -1 && nComments < N_FIRST_COMMENTS; {
		nComments++

		nBytes += idx
		p_commentsDBCS = p_commentsDBCS[idx:] //starting from '\n'

		nextCommentIdx := matchComment(p_commentsDBCS)
		if nextCommentIdx != -1 {

			nBytes += nextCommentIdx
			p_commentsDBCS = p_commentsDBCS[nextCommentIdx:] //starting from beginning of the next comment.

			idx = bytes.Index(p_commentsDBCS, []byte{'\n'})
		} else {
			idx = -1
		}
	}

	if nComments < N_FIRST_COMMENTS { //no more '\n', but not enough comments yet, add the last comment.
		nBytes += len(p_commentsDBCS)
	}

	firstCommentsDBCS, theRestComments = commentsDBCS[:nBytes], commentsDBCS[nBytes:]
	if len(firstCommentsDBCS) == 0 {
		firstCommentsDBCS = nil
	}
	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return firstCommentsDBCS, theRestComments
}

var (
	matchRecommendBytes = []byte{ //\n推
		0x0a, 0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x37, 0x6d,
		0xb1, 0xc0, 0x20,
	}

	matchBooBytes = []byte{ //\n噓
		0x0a, 0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xbc, 0x4e, 0x20,
	}

	matchArrowBytes = []byte{ //\n→
		0x0a, 0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xa1, 0xf7, 0x20,
	}
)

func matchComment(content []byte) int {
	theIdx := -1
	idxRecommend := bytes.Index(content, matchRecommendBytes)

	if idxRecommend != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxRecommend)
	}
	idxBoo := bytes.Index(content, matchBooBytes)
	if idxBoo != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxBoo)
	}
	idxArrow := bytes.Index(content, matchArrowBytes)
	if idxArrow != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxArrow)
	}
	if theIdx == -1 {
		return theIdx
	}

	return theIdx + 1 //the prefix \n belongs to signature or content
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

//ParseComments
//
//It's possible that reply-edit-info is not in commentsDBCS
//but in allCommentsDBCS (firstComments).
//we need allCommentsDBCS to get the edit-time.
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

	nEstimatedComments := bytes.Count(commentsDBCS, []byte{'\n'})

	comments = make([]*schema.Comment, 0, nEstimatedComments)

	p_commentsDBCS := commentsDBCS
	p_allCommentsDBCS := allCommentsDBCS

	lastTime := lastTimeNanoTS.ToTime()

	currentEditNanoTS := types.NanoTS(0)
	currentEditIP := ""
	currentEditHost := ""
	var comment *schema.Comment
	var reply *schema.Comment
	for idx := bytes.Index(p_commentsDBCS, []byte{'\n'}); len(p_commentsDBCS) > 0 && idx != -1; idx = bytes.Index(p_commentsDBCS, []byte{'\n'}) {
		commentDBCS := p_commentsDBCS[:idx]
		comment, lastTime = parseComment(bboardID, articleID, lastTime, commentDBCS, updateNanoTS, isFirstComments)
		comments = append(comments, comment)

		p_commentsDBCS = p_commentsDBCS[idx:] // with '\n'
		p_allCommentsDBCS = p_allCommentsDBCS[idx:]

		nextCommentIdx := matchComment(p_commentsDBCS)

		if nextCommentIdx == -1 { // no more comments
			p_commentsDBCS = p_commentsDBCS[1:] //step forward '\n'
			p_allCommentsDBCS = p_allCommentsDBCS[1:]
			if len(p_commentsDBCS) > 0 {
				replyDBCS := p_commentsDBCS
				reply, currentEditNanoTS, currentEditIP, currentEditHost = parseReply(bboardID, articleID, ownerID, comment.CreateTime, replyDBCS, comment.CommentID, p_allCommentsDBCS, currentEditNanoTS, currentEditIP, currentEditHost, updateNanoTS, isFirstComments)
				if reply != nil {
					comments = append(comments, reply)
				}

				p_allCommentsDBCS = p_allCommentsDBCS[len(p_commentsDBCS):]
				p_commentsDBCS = nil
			}

			//log.Infof("parseFirstcommentsCore: no more comments: p_commentsDBCS: %v to break", p_commentsDBCS)
			break
		}

		if nextCommentIdx > 1 { // p_commentsDBCS[0] is '\n', get reply from p_commentsDBCS[1:]
			replyDBCS := p_commentsDBCS[1:nextCommentIdx]

			reply, currentEditNanoTS, currentEditIP, currentEditHost = parseReply(bboardID, articleID, ownerID, comment.CreateTime, replyDBCS, comment.CommentID, p_allCommentsDBCS[1:], currentEditNanoTS, currentEditIP, currentEditHost, updateNanoTS, isFirstComments)
			//log.Infof("after parseReply: reply: %v", reply)
			if reply != nil {
				comments = append(comments, reply)
			}
		}

		p_commentsDBCS = p_commentsDBCS[nextCommentIdx:]
		p_allCommentsDBCS = p_allCommentsDBCS[nextCommentIdx:]
	}

	if len(p_commentsDBCS) > 0 { // last comment without reply.
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
		BBoardID:        bboardID,
		ArticleID:       articleID,
		CommentID:       commentID,
		TheType:         theType,
		CreateTime:      createNanoTS,
		Owner:           ownerID,
		Content:         contentUtf8,
		IP:              ip,
		MD5:             commentMD5,
		UpdateNanoTS:    updateNanoTS,
		IsFirstComments: isFirstComments,
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

	if bytes.HasPrefix(p_commmentDBCS, matchRecommendBytes[1:]) {
		return types.COMMENT_TYPE_RECOMMEND, nextCommentDBCS
	} else if bytes.HasPrefix(p_commmentDBCS, matchBooBytes[1:]) {
		return types.COMMENT_TYPE_BOO, nextCommentDBCS
	} else if bytes.HasPrefix(p_commmentDBCS, matchArrowBytes[1:]) {
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

func parseReply(
	bboardID bbs.BBoardID,
	articleID bbs.ArticleID,
	ownerID bbs.UUserID,
	commentCreateTime types.NanoTS,
	replyDBCS []byte,
	commentID types.CommentID,
	editDBCS []byte,
	currentEditNanoTS types.NanoTS,
	currentEditIP string,
	currentEditHost string,
	updateNanoTS types.NanoTS,
	isFirstComments bool) (

	reply *schema.Comment,
	editNanoTS types.NanoTS,
	editIP string,
	editHost string) {

	//log.Infof("parseReply: start: replyDBCS: %v", replyDBCS)
	if len(replyDBCS) == 0 {
		return nil, currentEditNanoTS, currentEditIP, currentEditHost
	}
	if replyDBCS[len(replyDBCS)-1] == '\n' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil, currentEditNanoTS, currentEditIP, currentEditHost
	}
	if replyDBCS[len(replyDBCS)-1] == '\r' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil, currentEditNanoTS, currentEditIP, currentEditHost
	}
	//log.Infof("parseReply: after purify: replyDBCS: %v", replyDBCS)

	replyMD5 := md5sum(replyDBCS)

	replyID := types.ToReplyID(commentID)

	replyBig5 := dbcsToBig5(replyDBCS)
	replyUtf8 := big5ToUtf8(replyBig5)

	if currentEditNanoTS > commentCreateTime {
		editNanoTS = currentEditNanoTS
		editIP = currentEditIP
		editHost = currentEditHost
	} else {
		editNanoTS, editIP, editHost = parseReplyIPHost(editDBCS)
		if editNanoTS == 0 {
			editNanoTS = commentCreateTime + REPLY_STEP_NANO_TS
		}
	}

	createNanoTS := commentCreateTime + REPLY_STEP_NANO_TS

	reply = &schema.Comment{
		BBoardID:   bboardID,
		ArticleID:  articleID,
		CommentID:  replyID,
		TheType:    types.COMMENT_TYPE_REPLY,
		RefIDs:     []types.CommentID{commentID},
		CreateTime: createNanoTS,
		Owner:      ownerID,
		Content:    replyUtf8,
		IP:         editIP,
		Host:       editHost,
		MD5:        replyMD5,

		EditNanoTS:      editNanoTS,
		UpdateNanoTS:    updateNanoTS,
		IsFirstComments: isFirstComments,
	}

	reply.CleanReply()
	if len(reply.Content) == 0 {
		return nil, editNanoTS, editIP, editHost
	}

	return reply, editNanoTS, editIP, editHost
}

var (
	editPrefix = []byte("\xa1\xb0 \xbds\xbf\xe8: ")
)

func parseReplyIPHost(editDBCS []byte) (editNanoTS types.NanoTS, editIP string, editHost string) {

	p_editDBCS := editDBCS
	theIdx := bytes.Index(p_editDBCS, editPrefix)
	if theIdx == -1 {
		return 0, "", ""
	}

	p_editDBCS = p_editDBCS[theIdx+len(editPrefix):]

	theIdx = bytes.Index(p_editDBCS, []byte{'('})
	if theIdx == -1 {
		return 0, "", ""
	}
	p_editDBCS = p_editDBCS[theIdx+1:]

	theIdx = bytes.Index(p_editDBCS, []byte{')'})
	if theIdx == -1 {
		return 0, "", ""
	}
	ipHost := types.Big5ToUtf8(p_editDBCS[:theIdx])

	ipHostList := strings.Split(ipHost, " ")
	if len(ipHostList) == 1 {
		editIP = ipHostList[0]
	} else {
		editIP = ipHostList[0]
		editHost = ipHostList[1]
	}

	p_editDBCS = p_editDBCS[theIdx:]

	theIdx = bytes.Index(p_editDBCS, []byte(", "))
	p_editDBCS = p_editDBCS[theIdx+2:]

	theTime, err := types.DateYearTimeStrToTime(string(p_editDBCS[:19]))
	if err != nil {
		return 0, "", ""
	}

	return types.TimeToNanoTS(theTime), editIP, editHost

}
