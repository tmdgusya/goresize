package image_test

import (
	"goresize/image"
	"os"
	"testing"
)

// normal case
func TestDecodeFileToImage(t *testing.T) {
	imageForTest := image.OpenImageFile("test.jpeg")

	// shouldn't return error
	if _, err := image.DecodeFileToImage(imageForTest); err != nil {
		t.Fatalf("you must check the decode logic : %+v\n", err.Error())
	}
}

func TestDecodeNotImageFile(t *testing.T) {
	os.Create("test.txt")

	imageForTest := image.OpenImageFile("test.txt")

	defer os.Remove("test.txt")

	decodedImage, err := image.DecodeFileToImage(imageForTest)

	if decodedImage != nil && err == nil {
		t.Fatalf("text file was decoded as well... you should check the logic")
	}
}
