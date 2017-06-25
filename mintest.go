package mintest

import "fmt"

type TestSlice struct {
	Input  interface{}
	Must   interface{}
	Error  bool
	result interface{}
}

type TestSlices []TestSlice

func Error(pattern TestSlice, result interface{}) string {
	return fmt.Sprintln("\nOutcome: Error", "\nResult:", fmt.Sprintf("%#v", result), "\nPattern:", fmt.Sprintf("%#v", pattern))
}

func Log(pattern TestSlice, result interface{}) string {
	return fmt.Sprintln("\nOutcome: Success", "\nResult :", fmt.Sprintf("%#v", result), "\nPattern:", fmt.Sprintf("%#v", pattern))
}
