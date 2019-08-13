package library

import (
	"math/rand"
	"encoding/hex"
	"crypto/md5"
)

func GenerateUUID()chan int{
	ch := make(chan int,10)
	go func() {
		ch <-rand.Int()
	}()
	return ch
}

func MakeMD5(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}