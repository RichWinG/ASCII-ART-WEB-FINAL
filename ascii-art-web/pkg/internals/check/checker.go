package check

import (
	"fmt"
)

func Valid(s []string) bool {
	for _, el := range s {
		if len(el) > 0 {
			return true
		}
	}
	return false
}

func Ascii(input string) bool {
	for _, el := range input {
		if (el < ' ' || el > '~') && el != '\n' {
			fmt.Println("Your input should be in ascii")
			return false
		}
	}
	return true
}
