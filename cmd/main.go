package main

//Using: convert --input inputfile --output outputfile
//То, что после --input, если не пусто, полное имя входного файла
//Тип файла определяется по его расширению из возможных, иначе - ошибка

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
)

const (
	NotfoundParameter      = 1
	InputFileNotRecognized = 10
)

func main() {
	if os.Args[1] != "--input" {
		fmt.Println("not found parameter --input")
		os.Exit(NotfoundParameter)
	}
	var inputFileName = os.Args[2]
	if inputFileName != "" {
		fmt.Println("input file not recognized")
		os.Exit(InputFileNotRecognized)
	}
	switch getFileType(inputFileName) {
	case "image/jpeg":
	case "image/png":
	case "image/bmp":
	}

}

func getFileType(fname string) string {
	var kind string
	buf, isImg := isValid(fname)
	if isImg {
		kind, _ := filetype.Match(buf[:])
		return kind.MIME.Value
	}
	return kind
}

func isValid(fname string) (buf []byte, valid bool) {
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
	buf = make([]byte, 100)
	file.Read(buf)
	return buf, filetype.IsImage(buf)
}
