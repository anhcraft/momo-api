package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Numerals = "0123456789"
const Alphanumeric = Alphabet + Numerals
const Hexadecimal = "0123456789abcde"

func RandString(n int, pool string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = pool[rand.Int63()%int64(len(pool))]
	}
	return string(b)
}

func GenerateDeviceToken() string {
	return strings.ToUpper(RandString(64, Hexadecimal))
}

func GenerateModelId() string {
	return RandString(64, Hexadecimal)
}

func GenerateImei() string {
	return RandString(64, Hexadecimal)
}

func GenerateSessionKeyTracking() string {
	return strings.ToUpper(uuid.New().String())
}

func GenerateRKey() string {
	return RandString(20, Alphanumeric)
}

func GenerateFirebaseToken() string {
	return RandString(22, Alphanumeric) + ":" + RandString(30, Alphanumeric) + "-" +
		RandString(21, Alphanumeric) + "-" + RandString(24, Alphanumeric) + "_" +
		RandString(22, Alphanumeric) + "_" + RandString(11, Alphanumeric) + "_" +
		RandString(22, Alphanumeric) + "-" + RandString(4, Alphanumeric)
}

func GenerateCmdId() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10) + "000000"
}
