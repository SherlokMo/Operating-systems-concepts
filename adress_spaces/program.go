package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		panic("Please insert an argument with the seconds you want this code to run for")
	}
	fmt.Printf("Process ID = %d \n", os.Getegid())
	PrintMemUsage()
	fmt.Println("Memory after:")
	seconds, _ := strconv.Atoi(os.Args[1])
	var overall [][]int
	for i := 0; i < seconds; i++ {
		a := make([]int, 0, 999999)
		overall = append(overall, a)
		PrintMemUsage()
		time.Sleep(time.Second)
	}

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
