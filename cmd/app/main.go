package main

import (
	"log"
	"pex-prevalent-colors-challenge/internal/app/cpu"
	"pex-prevalent-colors-challenge/internal/app/csvtohtml"
	"pex-prevalent-colors-challenge/internal/app/memory"
	"pex-prevalent-colors-challenge/internal/app/models"
	"pex-prevalent-colors-challenge/internal/app/pipeline"
	"time"
)

func main() {
	start := time.Now()
	env := models.NewEnv()

	//If the environment is set to monitor the CPU usage
	if env.ENABLE_CPU_MONITOR == "true" {
		cpufile := cpu.StartMonitorCPU(env.CPU_PPROF_FILENAME)
		defer cpufile.Close()
		defer cpu.StopMonitorCPU()
	}

	pipeline.Run(env)

	if env.GENERATE_HTML == "true" {
		csvtohtml.OpenCsvOnBrowser(env.OUTPUT_PATH+env.CSV_OUTPUT_FILENAME, env.HTML_TEMPLATE_FILENAME, env.OUTPUT_PATH+env.HTML_OUTPUT_FILENAME)
	}

	elapsed := time.Since(start)
	log.Printf("Process took %s", elapsed)

	if env.ENABLE_MEMORY_MONITOR == "true" {
		memory.MonitorMemory()
	}
}
