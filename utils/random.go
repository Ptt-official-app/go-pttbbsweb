package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// avoid '-' and '_' in leading char.
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

// GenRandomInt64
// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32350135
func GenRandomInt64(theMax int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(theMax))
	if err != nil {
		logrus.Warnf("GenRandomInt64: unable to rand: e: %v", err)
	}
	return nBig.Int64()
}
