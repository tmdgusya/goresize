package image_test

import (
	"goresize/image"
	"goresize/parse"
	"os"
	"testing"
)

func shouldPanicWith(message string, f func(file string) *os.File, filePath string, t *testing.T) {
	defer func() {
		err := recover().(error)

		if err.Error() != message {
			t.Fatalf("Wrong panic message: %s", err.Error())
		}
	}()
	f(filePath)
}

func TestPanicIfThereIsNoFileInTheGivenFilePath(t *testing.T) {
	given := parse.ResizedImageInfo{
		InputFileName:  "test-input.jpeg",
		OutputFileName: "test.jpeg",
		Width:          300,
		Height:         300,
	}

	shouldPanicWith("there is no image file in the given filepath", image.OpenImageFile, given.InputFileName, t)
}

func TestTheInputFilePathExistence(t *testing.T) {
	given := parse.ResizedImageInfo{
		InputFileName:  "test-input.jpeg",
		OutputFileName: "test.jpeg",
		Width:          300,
		Height:         300,
	}
	os.Create(given.InputFileName)
	defer os.Remove(given.InputFileName)

	testImage := image.OpenImageFile(given.InputFileName)

	if testImage == nil {
		t.Fatal("there is no image file given filepath : ", given.InputFileName)
	}
}
