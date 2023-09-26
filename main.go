package main

import (
	"flag"
	"fmt"
	image2 "goresize/image"
	"goresize/parse"
	"log"
)

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

	imageFile := image2.OpenImageFile(information.InputFileName)
	defer imageFile.Close()
	decodedImage, err := image2.DecodeFileToImage(imageFile)
	resizedImage := image2.NewImageFrom(decodedImage, information.Width, information.Height)
	resizedFile, err := image2.EncodeImageFile(resizedImage, information)
	resizedFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image resized successfully!")
}
