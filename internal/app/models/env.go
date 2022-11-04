package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	CPU_PPROF_FILENAME     string
	CSV_OUTPUT_FILENAME    string
	ENABLE_CPU_MONITOR     string
	ENABLE_MEMORY_MONITOR  string
	GENERATE_HTML          string
	HTML_OUTPUT_FILENAME   string
	HTML_TEMPLATE_FILENAME string
	INPUT_FILENAME         string
	INPUT_PATH             string
	OUTPUT_PATH            string
	PREVALENT_MODE         string
	REDUCE_IMAGES          string
}

func NewEnv() *Env {
	// load .env file from given path
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return &Env{
		INPUT_FILENAME:         os.Getenv("INPUT_FILENAME"),
		INPUT_PATH:             os.Getenv("INPUT_PATH"),
		CSV_OUTPUT_FILENAME:    os.Getenv("CSV_OUTPUT_FILENAME"),
		HTML_OUTPUT_FILENAME:   os.Getenv("HTML_OUTPUT_FILENAME"),
		OUTPUT_PATH:            os.Getenv("OUTPUT_PATH"),
		HTML_TEMPLATE_FILENAME: os.Getenv("HTML_TEMPLATE_FILENAME"),
		ENABLE_CPU_MONITOR:     os.Getenv("ENABLE_CPU_MONITOR"),
		CPU_PPROF_FILENAME:     os.Getenv("CPU_PPROF_FILENAME"),
		ENABLE_MEMORY_MONITOR:  os.Getenv("ENABLE_MEMORY_MONITOR"),
		REDUCE_IMAGES:          os.Getenv("REDUCE_IMAGES"),
		PREVALENT_MODE:         os.Getenv("PREVALENT_MODE"),
		GENERATE_HTML:          os.Getenv("GENERATE_HTML"),
	}
}
