package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomStr(n int) string {
	var s strings.Builder
	for i := 1; i <= n; i++ {
		s.WriteByte(byte('a' + rand.Intn(26)))
	}
	return s.String()
}

func RandomUsr() string {
	return RandomStr(10)
}

func RandomFullName() string {
	return RandomStr(20)
}

func RandomPass() string {
	return RandomStr(10)
}

func RandomUserType() int32 {
	return RandomInt(1, 3)
}

func RandomShortname() string {
	return RandomStr(10)
}

func RandomProblemname() string {
	return RandomStr(10)
}

func RandomContent() string {
	return RandomStr(20)
}

func RandomSubtasks() int32 {
	return RandomInt(1, 6)
}

func RandomAnswers(n int) []string {
	val := make([]string, n)
	for i := 0; i < n; i++ {
		val[i] = RandomStr(3)
	}
	return val
}
