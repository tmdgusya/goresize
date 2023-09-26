package main

import (
	"testing"
)

func ShouldPanicWith(message string, f func() error, t *testing.T) {
	defer func() {
		err := recover().(error)

		if err.Error() != message {
			t.Fatalf("Wrong panic message: %s", err.Error())
		}
	}()
	f()
}

func TestPanicIfTheClientDidntTypeInputFilePath(t *testing.T) {
	given := ResizedImageInfo{}
	ShouldPanicWith("you must write the input file name", given.Validate, t)
}

func TestPanicIfTheClientDidntTypeOutputFilePath(t *testing.T) {
	given := ResizedImageInfo{
		InputFileName: "input.jpeg",
	}
	ShouldPanicWith("you must write the output file name", given.Validate, t)
}

func TestPanicIfTheClientDidntTypeWidth(t *testing.T) {
	given := ResizedImageInfo{
		InputFileName:  "input.jpeg",
		OutputFileName: "output.jpeg",
	}
	ShouldPanicWith("you must write the width of the resized image", given.Validate, t)
}

func TestPanicIfTheClientDidntTypeHeight(t *testing.T) {
	given := ResizedImageInfo{
		InputFileName:  "input.jpeg",
		OutputFileName: "output.jpeg",
		Width:          300,
	}
	ShouldPanicWith("you must write the height of the resized image", given.Validate, t)
}
