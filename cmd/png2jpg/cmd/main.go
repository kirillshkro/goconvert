package main

//Using: convert --input inputfile --output outputfile
//То, что после --input, если не пусто, полное имя входного файла
//Тип файла определяется по его расширению из возможных, иначе - ошибка

import (
	"fmt"
	"os"
	"strings"

	"github.com/h2non/filetype"
)

const (
	NotfoundParameter      = 1
	InputFileNotRecognized = 10
)

func main() {
	var ext string = ""
	if os.Args[1] != "--input" {
		fmt.Println("not found parameter --input")
		os.Exit(NotfoundParameter)
	}
	var inputFileName = os.Args[2]
	if inputFileName != "" {
		fmt.Println("input file not recognized")
		os.Exit(InputFileNotRecognized)
	}
	ext = strings.Split(inputFileName, ".")[1]
	if getFileType(inputFileName) {
		switch ext {
		case "jpg":
		case "jpeg":
		case "png":
		}
	}

}

func getFileType(fname string) bool {
	file, err := os.Open(fname)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var buf [100]byte
	file.Read(buf[:])
	return filetype.IsImage(buf[:])
}
