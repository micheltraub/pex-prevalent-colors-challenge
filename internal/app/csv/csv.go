package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type AppendCsv interface {
	AppendToCsvFile() error
	OpenFileToAppend() error
	AppendToOsFile() error
}

type AppendCsvImpl struct {
	filename string
	data     []string
	csvFile  io.WriteCloser
}

func NewAppendCsvImpl(filename string, data []string) AppendCsv {

	return &AppendCsvImpl{
		filename: filename,
		data:     data,
	}
}

func (a *AppendCsvImpl) AppendToCsvFile() error {
	log.Println("💾 Appending line: ", a.data)
	if len(a.data) == 0 {
		return fmt.Errorf("🚧🚨 Empty data received")
	}
	err := a.OpenFileToAppend()
	if err != nil {
		return err
	}
	defer a.csvFile.Close()

	//Append data to file.
	err = a.AppendToOsFile()
	if err != nil {
		return err
	}

	return nil
}

func (a *AppendCsvImpl) OpenFileToAppend() error {
	//Open CSV file to append. Create new file if does not exist.
	csvFile, err := os.OpenFile(a.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("🚧🚨 Error opening file to append: " + err.Error())
	}

	a.csvFile = csvFile
	return nil
}

func (a *AppendCsvImpl) AppendToOsFile() error {
	//Append data to file.
	csvwriter := csv.NewWriter(a.csvFile)
	err := csvwriter.Write(a.data)
	if err != nil {
		return fmt.Errorf("🚧🚨 Error writting to file: " + err.Error())
	}
	csvwriter.Flush()
	return nil
}
