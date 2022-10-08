package main

import (
	"fmt"

	"github.com/nirmaldavis/mylogger-go"
)

func main() {
	fmt.Println("Hello Go..!!")

	mylogger.LogInfo("I am using a Go module")
	mylogger.LogWarning("Working .. ?!!")
	mylogger.LogError("Filed to work :-( ")
}
