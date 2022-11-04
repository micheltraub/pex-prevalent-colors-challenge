package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func AppendToCsvFile(filename string, data []string) error {
	log.Println("💾 Appending line: ", data, "to", filename)
	if len(data) == 0 {
		return fmt.Errorf("🚧🚨 Empty data received")
	}
	csvFile, err := OpenFileToAppend(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	//Append data to file.
	err = AppendToOsFile(csvFile, data)
	if err != nil {
		return err
	}

	return nil
}

func OpenFileToAppend(filename string) (io.WriteCloser, error) {
	//Open CSV file to append. Create new file if does not exist.
	csvFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("🚧🚨 Error opening file to append: " + err.Error())
	}

	return csvFile, nil
}

func AppendToOsFile(csvFile io.Writer, data []string) error {
	//Append data to file.
	csvwriter := csv.NewWriter(csvFile)
	err := csvwriter.Write(data)
	if err != nil {
		return fmt.Errorf("🚧🚨 Error writting to file: " + err.Error())
	}
	csvwriter.Flush()
	return nil
}
