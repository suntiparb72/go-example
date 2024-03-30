package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/suntiparb72/go-example/suntiparb"
)

func main() {
	id := uuid.New()
	fmt.Println("Hell World")
	fmt.Printf("UUID: %s", id)
	fmt.Print("\n")
	suntiparb.SayHello()
}
