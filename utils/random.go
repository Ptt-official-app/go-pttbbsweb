package utils

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func GenRandomString() string {
	random := uuid.Must(uuid.NewRandom())
	bin, _ := random.MarshalBinary()
	str := base64.URLEncoding.EncodeToString(bin)
	return str[:22]
}
