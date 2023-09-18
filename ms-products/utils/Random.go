package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var random *rand.Rand

const DefaultAlphabet = "abcdefghijklmnopqrstuvxwyz"

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(length int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[random.Intn(k)]
		err := sb.WriteByte(c)
		if err != nil {
			panic(1)
		}
	}

	return sb.String()
}

func RandomProductName() string {
	return fmt.Sprintf("%s %s", RandomString(6, DefaultAlphabet), RandomString(6, DefaultAlphabet))
}

func RandomProductDescription() string {
	return RandomString(20, DefaultAlphabet)
}

func RandomProductPrice() string {
	nums := "123456789"
	numsZero := "0123456789"
	return fmt.Sprintf("%s.%s", RandomString(3, nums), RandomString(2, numsZero))
}

func RandomProductID() int32 {
	return int32(random.Intn(1000))
}
