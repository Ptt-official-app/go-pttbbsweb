package utils

import (
	"encoding/base64"

	"github.com/google/uuid"
)

//avoid '-' and '_' in leading char.
func GenRandomString() string {
	for {
		random := uuid.Must(uuid.NewRandom())
		bin, _ := random.MarshalBinary()
		str := base64.URLEncoding.EncodeToString(bin)
		if str[0] == '-' || str[0] == '_' {
			continue
		}
		return str[:22]
	}
}
