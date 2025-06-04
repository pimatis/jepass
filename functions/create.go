package functions

import (
	"math/rand"
)

func Create(length int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	upperAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	password := ""

	for i := 0; i < length; i++ {
		password += string(alphabet[rand.Intn(len(alphabet))])
		password += string(upperAlphabet[rand.Intn(len(upperAlphabet))])
		password += string(numbers[rand.Intn(len(numbers))])
		password += string(symbols[rand.Intn(len(symbols))])
	}

	return password
}
