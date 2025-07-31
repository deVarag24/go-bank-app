package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

// Generate random integer between nim to max
func RandomInt(min, max int64) int64{
	return min + rand.Int63n(max - min + 1)
}

// Generate random stringof length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i:=0; i<n; i++ {
		c := alphabet[rand.Intn(k)]

		sb.WriteByte(c)
	}

	return sb.String()
}

// Return random element from enum
func RandomEnum[T any](arr []T) T{
	i := len(arr)

	return arr[rand.Intn(i)]
}

// Generate ramdom uuid
func RandomUuid() uuid.UUID {
	return uuid.New()
}