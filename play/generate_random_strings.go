package main

import (
	"bytes"
	"math/rand"
	"time"
)

func generateRandomStringFromChars(charset string, length int) string {
	rand.Seed(time.Now().Unix())
	var b bytes.Buffer

	for i := 0; i < length; i++ {
		b.WriteString(string(charset[rand.Intn(len(charset))]))
	}

	return b.String()
}

func generateLdapId(merchant_name string) string {
	ldapId := generateRandomStringFromChars("ABCDEFGHJKMNPQRSTUVWXYZ123456789", 10)

	return ldapId
}
