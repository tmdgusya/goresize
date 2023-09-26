package image_test

import (
	"goresize/image"
	"testing"
)

func TestResizeImageToGivenWidthAndHeight(t *testing.T) {
	width, height := 300, 30
	imageForTest := image.OpenImageFile("test.jpeg")
	decodedImage, _ := image.DecodeFileToImage(imageForTest)
	newImage := image.NewImageFrom(decodedImage, width, height)
	rect := newImage.Bounds()

	if rect.Max.X != width || rect.Max.Y != height {
		t.Fatalf("the resized image size is difference from your given information %+v\n", rect)
	}
}
