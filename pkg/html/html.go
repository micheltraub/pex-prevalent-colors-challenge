package html

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
)

type CsvLine struct {
	Img    string
	Color1 string
	Color2 string
	Color3 string
}

// - EXTRA FEATURE -
// create a visual representation of the output CSV
func CreateHtmlFromCsv(csvPath string, htmlTemplatePath string, htmlPath string) error {

	//Open CSV file
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not open CSV file (%s): %s", csvPath, err)
	}
	defer csvFile.Close()

	//Read CSV file content
	csvReader := csv.NewReader(csvFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not read CSV file (%s): %s", csvPath, err)
	}
	//Parse CSV content into
	csvLines := createCsvLines(data)

	//Read HTML template
	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not read HTML template (%s): %s", htmlTemplatePath, err)
	}
	var processed bytes.Buffer
	//Execute template with CSV content
	err = tmpl.ExecuteTemplate(&processed, "result", csvLines)
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not execute HTML template (%s): %s", htmlTemplatePath, err)
	}
	//Save to HTML file
	f, err := os.Create(htmlPath)
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not create HTML file (%s): %s", htmlPath, err)
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(processed.String())
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not write to bufio HTML file (%s): %s", htmlPath, err)
	}
	err = w.Flush()
	if err != nil {
		return fmt.Errorf("🚧🚨 Could not write buffered data from HTML file (%s): %s", htmlPath, err)
	}

	log.Printf("HTML file saved :%s \n", htmlPath)
	return nil
}

// foreach line from CSV content, parse the line as CsvLine type
func createCsvLines(data [][]string) []CsvLine {
	var csvLines []CsvLine
	for _, line := range data {
		var rec CsvLine
		for j, field := range line {
			if j == 0 {
				rec.Img = field
			} else if j == 1 {
				rec.Color1 = field
			} else if j == 2 {
				rec.Color2 = field
			} else if j == 3 {
				rec.Color3 = field
			}
		}
		csvLines = append(csvLines, rec)
	}
	return csvLines
}
