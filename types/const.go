package types

const (
	//TODO: have go-pttbbs able to be modulized

	//////////
	//pttstruct.h
	//
	//SHMSIZE is computed in cache/shm.go NewSHM
	//////////
	IDLEN   = 12 /* Length of bid/uid */
	IPV4LEN = 15 /* a.b.c.d form */

	PASS_INPUT_LEN = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	PASSLEN = 14 /* Length of encrypted passwd field */
	REGLEN  = 38 /* Length of registration data */

	REALNAMESZ = 20 /* Size of real-name field */
	NICKNAMESZ = 24 /* SIze of nick-name field */
	EMAILSZ    = 50 /* Size of email field */
	ADDRESSSZ  = 50 /* Size of address field */
	CAREERSZ   = 40 /* Size of career field */
	PHONESZ    = 20 /* Size of phone field */

	USERNAMESZ = 24 /* Size of Username in MailQueue */
	RCPTSZ     = 50 /* Size of RCPT in MailQueue */

	PASSWD_VERSION = 4194

	BTLEN = 48 /* Length of board title */

	TTLEN = 64 /* Length of title */
	FNLEN = 28 /* Length of filename */
)
