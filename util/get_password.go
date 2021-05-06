package util

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
	"syscall"
)

func GetPassword() (string, error) {
	fmt.Println("Enter Password (Press Enter if you dont have any): ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}
