package utils

import (
	base "encoding/base64"
	"time"
)

var (
	TimeLayout string = "2006-01-02"
)

func ConvertTimeToUnixMillis(timeStamp string) (int64, error) {
	now, err := time.Parse(TimeLayout, timeStamp)

	if err != nil {
		return 0, err
	}
	return (now.UnixNano() / 1000000), nil
}

func EncodeToBase64(s string) string {
	return base.StdEncoding.EncodeToString([]byte(s))
}

func DecodeBase64(s string) (string, error) {
	b, err := base.StdEncoding.DecodeString(s)

	return string(b), err
}
