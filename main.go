package main

import (
	"flag"
	"fmt"
	"goresize/parse"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func resizeImage(r *parse.ResizedImageInfo) error {
	// Open the input image file.
	inputFile, err := os.Open(r.InputFileName)
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
	newImg := resize(img, r.Width, r.Height)

	// Create the output image file.
	outputFile, err := os.Create(r.OutputFileName)
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

	information := &parse.ResizedImageInfo{
		Width:          *width,
		Height:         *height,
		InputFileName:  *inputFile,
		OutputFileName: *outputFile,
	}

	err := resizeImage(information)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image resized successfully!")
}
