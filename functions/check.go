package functions

import (
	"strings"
	"github.com/ttacon/chalk"
)

type PasswordStrength int

const (
	Weak PasswordStrength = iota
	Medium
	Strong
	VeryStrong
)

func (ps PasswordStrength) String() string {
	switch ps {
	case Weak:
		return "Weak"
	case Medium:
		return "Medium"
	case Strong:
		return "Strong"
	case VeryStrong:
		return "Very Strong"
	default:
		return "Unknown"
	}
}

func Check(password string) (bool, PasswordStrength, []string) {
	var issues []string
	score := 0

	if len(password) < 8 {
		issues = append(issues, chalk.Red.Color("Password must be at least 8 characters long"))
	} else if len(password) >= 12 {
		score += 2
	} else {
		score += 1
	}

	hasLower := false
	hasUpper := false
	hasNumber := false
	hasSymbol := false

	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasNumber = true
		default:
			hasSymbol = true
		}
	}

	if !hasLower {
		issues = append(issues, chalk.Red.Color("Password should contain lowercase letters"))
	} else {
		score++
	}

	if !hasUpper {
		issues = append(issues, chalk.Red.Color("Password should contain uppercase letters"))
	} else {
		score++
	}

	if !hasNumber {
		issues = append(issues, chalk.Red.Color("Password should contain numbers"))
	} else {
		score++
	}

	if !hasSymbol {
		issues = append(issues, chalk.Red.Color("Password should contain symbols"))
	} else {
		score++
	}

	if strings.Contains(strings.ToLower(password), "password") ||
		strings.Contains(strings.ToLower(password), "123456") {
		issues = append(issues, chalk.Red.Color("Password contains common patterns"))
		score--
	}

	isValid := len(issues) == 0
	
	var strength PasswordStrength
	switch {
	case score <= 2:
		strength = Weak
	case score <= 4:
		strength = Medium
	case score <= 6:
		strength = Strong
	default:
		strength = VeryStrong
	}

	return isValid, strength, issues
}
