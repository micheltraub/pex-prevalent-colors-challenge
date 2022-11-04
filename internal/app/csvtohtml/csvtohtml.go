package csvtohtml

import (
	"log"
	"os"
	"pex-prevalent-colors-challenge/pkg/html"
)

func OpenCsvOnBrowser(csvPath string, htmlTemplatePath string, htmlPath string) {
	err := html.CreateHtmlFromCsv(csvPath, htmlTemplatePath, htmlPath)
	if err != nil {
		log.Fatal(err)
	}
	fullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	html.Open("file:///" + fullPath + htmlPath[1:])
}
