package main

import (
	"fmt"
	"jepass/functions"
)

func main() {
	fmt.Println("--- Welcome to JePass ---")
	fmt.Println("Please select the operation you want to perform. \n1. Password creation\n2. Password update\n3. Check Password\n4. Exit")
	var choice int
	fmt.Scanln(&choice)

main:
	switch choice {
	case 1:
		var passLen int
		fmt.Print("Please enter the length of the password you want to create: ")
		fmt.Scanln(&passLen)

		if passLen < 8 {
			fmt.Println("Password length must be at least 8 characters.")
		} else {
			pass := functions.Create(passLen)
			fmt.Println("Generated password:", pass)
		}

	case 2:
		var oldPass string
		fmt.Print("Please enter the old password you want to update: ")
		fmt.Scanln(&oldPass)

		newPass := functions.Update(oldPass)
		fmt.Println("Updated password:", newPass)

	case 3:
		var password string
		fmt.Print("Please enter the password you want to check: ")
		fmt.Scanln(&password)
		isValid, strength, issues := functions.Check(password)

		if isValid {
			fmt.Println("Password is valid. Strength:", strength)
		} else {
			fmt.Println("Password is invalid. Issues found:")
			for _, issue := range issues {
				fmt.Println("-", issue)
			}
		}
	case 4:
		fmt.Println("Exiting the program.")
		return
	default:
		fmt.Println("Invalid selection. Please choose 1, 2, 3, or 4.")
		fmt.Println("Program terminating...")
		return
	}

	fmt.Println("Do you want to perform another operation? (Yes: 1, No: 2)")
	var again int
	fmt.Scanln(&again)

	if again == 1 {
		fmt.Println("Please select the operation you want to perform. \n1. Password creation\n2. Password update\n3. Check Password\n4. Exit")
		fmt.Scanln(&choice)
		goto main
	} else {
		fmt.Println("Program terminating...")
	}

	fmt.Println("Program terminating...")
}