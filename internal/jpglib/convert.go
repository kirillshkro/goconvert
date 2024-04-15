package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func Jpeg2Png(fileName string, outFileName string) error {
	outFile, err := os.Create(outFileName)
	defer outFile.Sync()
	defer outFile.Close()
	defer convertError(err)
	if err != nil {
		panic(err)
	}
	image, err := jpeg2Image(fileName)
	defer convertError(err)
	if err != nil {
		panic(err)
	}
	return png.Encode(outFile, image)
}

func jpeg2Image(fileName string) (image.Image, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := file
	image, err := jpeg.Decode(r)
	if err != nil {
		panic(err)
	}
	return image, nil
}

func convertError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
