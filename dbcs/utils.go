package dbcs

import (
	"crypto/md5"
	"encoding/base64"
)

func md5sum(theBytes []byte) string {
	if len(theBytes) == 0 {
		return ""
	}
	hash := md5.Sum(theBytes)
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash[:])
}

func reverseList(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
