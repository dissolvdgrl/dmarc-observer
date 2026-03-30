package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("Dmarc Observer build in progress...")
	files, err := os.ReadDir("xml-reports")


	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
