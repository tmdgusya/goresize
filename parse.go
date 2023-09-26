package main

import "fmt"

type ResizedImageInfo struct {
	InputFileName  string
	OutputFileName string
	Width          int
	Height         int
}

func (r *ResizedImageInfo) Validate() error {
	if r.InputFileName == "" {
		panic(fmt.Errorf("you must write the input file name"))
	}
	if r.OutputFileName == "" {
		panic(fmt.Errorf("you must write the output file name"))
	}
	if r.Width == 0 {
		panic(fmt.Errorf("you must write the width of the resized image"))
	}
	if r.Height == 0 {
		panic(fmt.Errorf("you must write the height of the resized image"))
	}
	return nil
}
