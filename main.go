package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"path/filepath"
	"strings"
	"strconv"

	"dmarc-observer/parser"
	"dmarc-observer/extract"
)

func main() {
	fmt.Println("*** Dmarc Observer ***")
	fmt.Println("Please enter your downloaded XML files path: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	files, err := os.ReadDir(input)
	if err != nil {
		log.Fatal(err)
	}

	var	reports []parser.Feedback

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".zip") {
			xml, err := extract.ExtractZip(filepath.Join(input, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			report, err := parser.ParseReport(xml)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, *report)
		}

		if strings.HasSuffix(file.Name(), ".gz") {
			gz, err := extract.ExtractGz(filepath.Join(input, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			report, err := parser.ParseReport(gz)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, *report)
		}
	}
	

	fmt.Println(parser.PrintReportSummary(reports))

	for {
		fmt.Println("Choose a report or q to quit:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		report := reports[n]

		fmt.Println(parser.PrintFullReport(&report))
		fmt.Println("Press Enter to continue")
		reader.ReadString('\n')
		fmt.Println(parser.PrintReportSummary(reports))
	}
}
