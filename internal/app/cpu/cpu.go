package cpu

import (
	"os"
	"runtime/pprof"
)

func MonitorCPU(cpuPprofFilename string) {
	cpufile, err := os.Create(cpuPprofFilename)
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	defer pprof.StopCPUProfile()
}
