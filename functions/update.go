package functions

import (
	"math/rand"
)

func Update(oldPass string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	upperAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	password := oldPass
	
	password += string(alphabet[rand.Intn(len(alphabet))])
	password += string(upperAlphabet[rand.Intn(len(upperAlphabet))])
	password += string(numbers[rand.Intn(len(numbers))])
	password += string(symbols[rand.Intn(len(symbols))])

	return password
}
