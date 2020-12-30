package api

func isValidRemoteAddr(remoteAddr string) bool {
	if IsTest {
		return true
	}
	switch remoteAddr {
	case "":
		return false
	default:
		return true
	}
}
