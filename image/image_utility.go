package image

import (
	"fmt"
	"golang.org/x/image/draw"
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

func NewImageFrom(img image.Image, width, height int) image.Image {
	// Create a new image with the desired dimensions.
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Use the Draw function to resize the image.
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImg
}
