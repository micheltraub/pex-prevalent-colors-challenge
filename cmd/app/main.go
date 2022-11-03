package main

import (
	"bufio"
	"log"
	"os"
	"pex-prevalent-colors-challenge/internal/app/accurateprevalent"
	"pex-prevalent-colors-challenge/internal/app/csv"
	"pex-prevalent-colors-challenge/internal/app/memory"
	"pex-prevalent-colors-challenge/internal/app/models"
	"pex-prevalent-colors-challenge/pkg/html"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	env := models.NewEnv()

	/*if env.ENABLE_CPU_MONITOR == "true" {
		cpu.MonitorCPU(env.CPU_PPROF_FILENAME)
	}*/
	cpufile, err := os.Create(env.CPU_PPROF_FILENAME)
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	defer pprof.StopCPUProfile()

	csvCh := make(chan []string, 50)
	jobs := make(chan int, 5)
	done := make(chan bool)
	jobsAppend := make(chan int, 5)
	doneWritingCsvFile := make(chan bool)
	var wg sync.WaitGroup

	go ProcessInputFile(done, jobs, csvCh, env, &wg)
	counter := 0
	for csvLine := range csvCh {
		counter++
		jobsAppend <- counter
		go csv.AppendToCsvFile(env.OUTPUT_PATH+env.CSV_OUTPUT_FILENAME, csvLine, jobsAppend, doneWritingCsvFile)
	}

	//OpenCsvOnBrowser(env.OUTPUT_PATH+env.CSV_OUTPUT_FILENAME, env.HTML_TEMPLATE_FILENAME, env.OUTPUT_PATH+env.HTML_OUTPUT_FILENAME)
	/* use number of cpus for goroutines
		var wg2 sync.WaitGroup
	    for i := 0; i < runtime.NumCPU(); i++ {
	        wg2.Add(1)
	        go ProcessInputFile(done, jobs, csvCh, env, &wg2)
	    }
	*/
	log.Println("Done")
	elapsed := time.Since(start)
	log.Printf("Process took %s", elapsed)

	if env.ENABLE_MEMORY_MONITOR == "true" {
		memory.MonitorMemory()
	}
}

func ProcessInputFile(done chan bool, jobs chan int, csvCh chan []string, env *models.Env, wg *sync.WaitGroup) {
	// Open input file
	file, err := os.Open(env.INPUT_PATH + env.INPUT_FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		wg.Add(1)
		imgUrl := scanner.Text()
		counter++
		jobs <- counter
		// ignore empty lines
		if len(imgUrl) == 0 {
			continue
		}
		go processLine(jobs, done, csvCh, imgUrl, wg)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	close(csvCh)

}

func processLine(jobs chan int, done chan bool, csvCh chan []string, imgUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, more := <-jobs
	prevalentColor := accurateprevalent.NewAccuratePrevalentColor(imgUrl, "-", "-", "-")
	prevalentcolors.ProcessPrevalentColors(prevalentColor, csvCh)
	if !more {
		done <- true
		return
	}

}

func OpenCsvOnBrowser(csvPath string, htmlTemplatePath string, htmlPath string) {
	err := html.CreateHtmlFromCsv(csvPath, htmlTemplatePath, htmlPath)
	fullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	html.Open("file:///" + fullPath + htmlPath[1:])
}
