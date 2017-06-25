package mintest

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type TestSlice struct {
	Input  interface{}
	Must   interface{}
	Error  bool
	result interface{}
}

type TestSlices []TestSlice

func Error(pattern TestSlice, result interface{}) string {
	return fmt.Sprintln("\nOutcome: Error", "\nResult :", spew.Sdump(result), "\nPattern:", spew.Sdump(pattern))
}

func Log(pattern TestSlice, result interface{}) string {
	return fmt.Sprintln("\nOutcome: Success", "\nResult :", spew.Sdump(result), "\nPattern:", spew.Sdump(pattern))
}
