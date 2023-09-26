package image

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// OpenImageFile You must close given file when you finished the logic
func OpenImageFile(filePath string) *os.File {
	// Open the input image file.
	inputFile, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("there is no image file in the given filepath"))
	}
	return inputFile
}

func DecodeFileToImage(f *os.File) (image.Image, error) {
	// Decode the input image.
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}
