package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func resizeImage(inputPath, outputPath string, width, height int) error {
	// Open the input image file.
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Decode the input image.
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Create a new image with the desired dimensions.
	newImg := resize(img, width, height)

	// Create the output image file.
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode the resized image as JPEG.
	err = jpeg.Encode(outputFile, newImg, nil)
	if err != nil {
		return err
	}

	return nil
}

func resize(img image.Image, width, height int) image.Image {
	// Create a new image with the desired dimensions.
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Use the Draw function to resize the image.
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImg
}

func main() {
	inputPath := "input.jpeg"   // Change this to your input image file path.
	outputPath := "output.jpeg" // Change this to your output image file path.
	width := 300                // Change this to your desired width.
	height := 200               // Change this to your desired height.

	err := resizeImage(inputPath, outputPath, width, height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image resized successfully!")
}
