package random

import (
	"fmt"
	"math/rand"
	"time"
)


func VerifyCode(length int) string {
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%v", rand.Intn(9999999))
	return code[0:length]
}
