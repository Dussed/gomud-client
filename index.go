package main

import (
	"fmt"

	// "github.com/dussed/gomud-client/drawing"
	"github.com/dussed/gomud-client/level"
)

func main() {
	fmt.Println("testing")

	err := level.Load("debug")

	if err != nil {
		fmt.Println(err.Error())
	}
	// drawing.DrawDebug()
}
