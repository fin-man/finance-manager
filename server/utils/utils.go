package utils

import (
	"crypto/sha1"
	"fmt"
	"log"
)

func Sha1Hash(s string) string {
	h := sha1.New()

	_, err := h.Write([]byte(s))

	if err != nil {
		log.Println(err)
		return s
	}

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
