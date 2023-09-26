package main

import (
	"flag"
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
	inputFile := flag.String("i", "", "input file name")
	outputFile := flag.String("o", "", "output file name")
	width := flag.Int("w", 300, "width for resized image")
	height := flag.Int("h", 300, "height for resized image")
	flag.Parse()

	if *inputFile == "" {
		panic(fmt.Errorf("you must write the input file path"))
	}

	if *outputFile == "" {
		panic(fmt.Errorf("you must write the output file path"))
	}

	if *width == 0 || *height == 0 {
		panic(fmt.Errorf("you must write the width and height of resized iamge"))
	}

	err := resizeImage(*inputFile, *outputFile, *width, *height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image resized successfully!")
}
