package util

import (
	"math/rand"
	"time"
)

func RandSeq(n int) string {
	time.Sleep(time.Microsecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
