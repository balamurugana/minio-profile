package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/minio/minio/pkg/bpool"
)

func main() {
	// testSyncPoolBytes()
	testBpool()
}

func printMemUsage(message string) {
	bToMb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("%s\tAlloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v\n", message, bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func testSyncPoolBytes() {
	bufPool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 2*10*1024*1024)
		},
	}

	printMemUsage("Before using pool:")

	bufs := make([][]byte, 100)

	for i := 0; i < 100; i++ {
		bufs[i] = bufPool.Get().([]byte)
	}

	for i := 0; i < 100; i++ {
		bufPool.Put(bufs[i])
		bufs[i] = nil
	}

	bufs = nil

	printMemUsage("After using pool:")
	for i := 0; i < 5; i++ {
		fmt.Println("Attempt", i+1, "sleeping for one minute")
		time.Sleep(1 * time.Minute)
		printMemUsage("")
	}

	runtime.GC()
	printMemUsage("After calling GC forcefully:")
}

func testBpool() {
	setCount := 1
	drivesPerSet := 8
	blockSizeV1 := 10 * 1024 * 1024
	bp := bpool.NewBytePoolCap(setCount*drivesPerSet, blockSizeV1, blockSizeV1*2)

	printMemUsage("Before using pool:")

	bufs := make([][]byte, 100)
	for i := 0; i < 100; i++ {
		bufs[i] = bp.Get()
	}

	for i := 0; i < 100; i++ {
		bp.Put(bufs[i])
		bufs[i] = nil
	}

	bufs = nil
	printMemUsage("After using pool:")

	for i := 0; i < 5; i++ {
		fmt.Println("Attempt", i+1, "sleeping for one minute")
		time.Sleep(1 * time.Minute)
		printMemUsage("")
	}

	runtime.GC()
	printMemUsage("After calling GC forcefully:")
}
