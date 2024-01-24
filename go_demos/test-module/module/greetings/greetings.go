package greetings

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("hi, %v. welcome",name)
}