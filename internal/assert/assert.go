package assert

import (
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func AssertTrue(x bool, msg string) {
	if !x {
		exit(msg)
	}
}
