package main

import (
	"fmt"

	"github.com/ton5169/GoLangLearning/greetings"
)

func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
