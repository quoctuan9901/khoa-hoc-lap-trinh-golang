package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPostiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap so nguyen duong.")
	}
}

func GetPostiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value >= 0 {
			return value
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap so thuc duong.")
	}
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap it nhat mot ky tu.")
	}
}

func GetOptionalString(prompt string, oldValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}

	return input
}

func GetOptionalPostiveFloat(prompt string, oldValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil && value < 0 {
		fmt.Println("❌ Gia tri khong hop le! Giu nguyen gia tri cu")
		return oldValue
	}
	
	return value
}

func ClearScreen() {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error clearing screen: ", err)
	}
}

type HasId interface {
	GetId() int
}

func IsIdUnique[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}

	return true
}