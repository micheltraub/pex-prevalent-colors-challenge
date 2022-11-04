package cpu

import (
	"os"
	"runtime/pprof"
)

func StartMonitorCPU(cpuPprofFilename string) *os.File {
	cpufile, err := os.Create(cpuPprofFilename)
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	return cpufile
}

func StopMonitorCPU() {
	pprof.StopCPUProfile()
}
