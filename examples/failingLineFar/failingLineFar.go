package main

import (
	"errors"
	"fmt"

	"github.com/gregwebs/errlog"
)

var (
	debug = errlog.NewLogger(&errlog.Config{
		PrintFunc:          func(format string, args ...any) { fmt.Printf(format, args...) },
		LinesBefore:        6,
		LinesAfter:         3,
		PrintError:         true,
		PrintSource:        true,
		PrintStack:         true,
		ExitOnDebugSuccess: true,
	})
)

func main() {
	fmt.Println("Start of the program")

	wrapingFunc()

	fmt.Println("End of the program")
}

func wrapingFunc() {
	someBigFunction()
}

func someBigFunction() {
	someDumbFunction()

	someSmallFunction()
	err := someNastyFunction()
	someDumbFunction()

	if debug.Debug(err) {
		return
	}

	someSmallFunction()

	someDumbFunction()
}

func someSmallFunction() {
	fmt.Println("I do things !")
}

func someNastyFunction() error {
	return errors.New("I'm failing for no reason")
}

func someDumbFunction() bool {
	return false
}
