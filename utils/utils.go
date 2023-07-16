package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(int64(time.Now().UnixMicro()))

}

//function to generate random numbers

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//function to generate random String

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// random owner

func RandomOwner() string {
	return RandomString(6)
}

//function to generate random money

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
