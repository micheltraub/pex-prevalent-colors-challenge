package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func AppendToCsvFile(filename string, data []string) {

	if len(data) == 0 {
		log.Println("Empty data received")
		return
	}
	log.Println("💾 Appending line: ", data)

	//Open CSV file to append. Create new file if does not exist.
	csvFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("Error opening file to append: ", err.Error())
		return
	}

	//Append data to file.
	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(data)
	csvwriter.Flush()
	csvFile.Sync()
}
