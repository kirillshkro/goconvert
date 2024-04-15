package bmplib

import (
	"image"
	"os"

	"golang.org/x/image/bmp"
)

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
