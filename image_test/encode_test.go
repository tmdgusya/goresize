package image

import (
	"goresize/image"
	"goresize/parse"
	"os"
	"testing"
)

func TestEncodeNormalFileToImageFile(t *testing.T) {
	given := parse.ResizedImageInfo{
		InputFileName:  "test.jpeg",
		OutputFileName: "test-output.jpeg",
		Width:          300,
		Height:         30,
	}
	imageForTest := image.OpenImageFile(given.InputFileName)
	defer imageForTest.Close()
	decodedImage, _ := image.DecodeFileToImage(imageForTest)
	newImage := image.NewImageFrom(decodedImage, given.Width, given.Height)
	outputFile, err := image.EncodeImageFile(newImage, &given)
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	if err != nil {
		t.Fatalf("error : %+v\n", err)
	}

	if outputFile.Name() != given.OutputFileName {
		t.Fatal("The file was not created successfully.")
	}
}
