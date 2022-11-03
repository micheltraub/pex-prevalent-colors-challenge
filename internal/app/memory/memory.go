package memory

import (
	"log"
	"runtime"
)

func MonitorMemory() {
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)
	log.Printf("\n")
	log.Printf("Alloc: %d MB, TotalAlloc: %d MB, Sys: %d MB\n",
		ms.Alloc/1024/1024, ms.TotalAlloc/1024/1024, ms.Sys/1024/1024)
	log.Printf("Mallocs: %d, Frees: %d\n",
		ms.Mallocs, ms.Frees)
	log.Printf("HeapAlloc: %d MB, HeapSys: %d MB, HeapIdle: %d MB\n",
		ms.HeapAlloc/1024/1024, ms.HeapSys/1024/1024, ms.HeapIdle/1024/1024)
	log.Printf("HeapObjects: %d\n", ms.HeapObjects)
	log.Printf("\n")
}
