package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Printf("Hello, %s", uuid.NewString())
}
