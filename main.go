package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"

	"dmarc-observer/parser"
	"dmarc-observer/zip"
)

func main() {
	fmt.Println("*** Dmarc Observer ***")
	fmt.Println("Please enter your zip file path: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	xml, err := zip.ExtractXML(input)
	if err != nil {
		log.Fatal(err)
	}
	
	report, err := parser.ParseReport(xml)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parser.PrintReport(report))
}
