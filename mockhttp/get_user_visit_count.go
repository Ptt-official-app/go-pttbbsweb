package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func GetUserVisitCount() (ret *api.GetUserVisitCountResult) {
	return &api.GetUserVisitCountResult{Total: 0}
}
