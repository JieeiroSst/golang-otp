package main

import (
	"fmt"
	"github.com/uaraven/gotp"
	"time"
)

func main() {
	totp := gotp.NewDefaultTOTP([]byte("secret key"))
	count := 10
	now := time.Now()
	timestamp := now.Add(time.Duration(-count) * time.Second)
	code := totp.At(timestamp)

	fmt.Println(code)

	ok := totp.VerifyAt(code, timestamp)
	fmt.Println(ok)
}