package utils

import (
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
