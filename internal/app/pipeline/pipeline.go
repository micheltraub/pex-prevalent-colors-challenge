package pipeline

import (
	"bufio"
	"log"
	"os"
	"pex-prevalent-colors-challenge/internal/app/accurateprevalent"
	"pex-prevalent-colors-challenge/internal/app/averageprevalent"
	"pex-prevalent-colors-challenge/internal/app/csv"
	"pex-prevalent-colors-challenge/internal/app/models"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"sync"
)

func Run(env *models.Env) {
	csvCh := make(chan []string)
	var wg sync.WaitGroup
	go processInputFile(csvCh, env, &wg)
	var wg2 sync.WaitGroup
	persistToCsvFile(csvCh, env, &wg2)
	wg2.Wait()
}

func persistToCsvFile(csvCh chan []string, env *models.Env, wg2 *sync.WaitGroup) {

	for c := range csvCh {
		csvLine := c
		wg2.Add(1)
		go func() {
			err := csv.AppendToCsvFile(env.OUTPUT_PATH+env.CSV_OUTPUT_FILENAME, csvLine)
			if err != nil {
				log.Println(err)
			}
			wg2.Done()
		}()
	}

}

func processInputFile(csvCh chan []string, env *models.Env, wg *sync.WaitGroup) {
	reduce := false
	if env.REDUCE_IMAGES == "true" {
		reduce = true
	}

	accurateMode := false
	if env.PREVALENT_MODE == "ACCURATE" {
		accurateMode = true
	}

	// Open input file
	file, err := os.Open(env.INPUT_PATH + env.INPUT_FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		wg.Add(1)
		imgUrl := scanner.Text()
		// ignore empty lines
		if len(imgUrl) == 0 {
			continue
		}
		go func() {
			processLine(csvCh, imgUrl, reduce, accurateMode)
			wg.Done()
		}()

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	close(csvCh)
}

func processLine(csvCh chan []string, imgUrl string, reduce bool, accurateMode bool) {
	var prevalentColor models.PrevalentColor
	if accurateMode {
		prevalentColor = accurateprevalent.NewAccuratePrevalentColor(imgUrl, "-", "-", "-", reduce)
	} else {
		prevalentColor = averageprevalent.NewAveragePrevalentColor(imgUrl, "-", "-", "-", reduce)
	}
	prevalentcolors.ProcessPrevalentColors(prevalentColor, csvCh)
}
