package main

import (
	"fmt"
	"os"
	"log"
	"dmarc-observer/parser"
)

func main() {
	fmt.Println("Dmarc Observer build in progress...")
	file, err := os.Open("xml-reports/report.xml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parser.ParseReport(file))

}
