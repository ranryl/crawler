package utils

import "fmt"

// PanicError panic error
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

// LogError print error
func LogError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
