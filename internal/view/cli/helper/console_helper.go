package helper

import (
	"fmt"
	"log"
)

func IsValid(err error) {
	if err != nil {
		log.Fatalf("Error: %v")
	}
}

func ReadKey() {
	_, err := fmt.Scan()
	if err != nil {}
}