package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"

	)

func main() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()


	stream, err :=ioutil.ReadFile("sample.txt")

	if err != nil {
	  log.Fatal(err)
	}

	readStream := string(stream)
	fmt.Println(readStream)
}
