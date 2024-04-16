package bmplib

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
)

func Bmp2Jpeg(fName string, outFName string) error {
	img, err := bmp2Image(fName)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(outFName)
	if err != nil {
		panic(err)
	}
	defer outFile.Sync()
	defer outFile.Close()
	return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
}

func Bmp2Png(fName string, outFName string) error {
	img, err := bmp2Image(fName)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(outFName)
	if err != nil {
		panic(err)
	}
	defer outFile.Sync()
	defer outFile.Close()
	return png.Encode(outFile, img)
}

func bmp2Image(fName string) (image.Image, error) {
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	image, err := bmp.Decode(file)
	if err != nil {
		panic(err)
	}
	return image, nil
}
