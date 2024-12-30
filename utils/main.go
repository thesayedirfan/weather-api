package utils

import "fmt"

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
