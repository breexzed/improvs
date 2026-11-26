//The program below tries to utilize all available CPUs by starting a busy loop for each core.
// It also keeps allocating memory in chunks of 10MB every second until the process is terminated.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup

	// Allocate all available CPUs
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Start a busy goroutine for each CPU
	for i := 0; i < numCPU; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			fmt.Println("Started a CPU hog")
			for ctx.Err() == nil { /* Busy loop to keep CPU busy */
			}
		}()
	}

	// Start allocating memory every second
	const memChunkMiB = 10
	var chunks [][]byte
	go func() {
		wg.Add(1)
		defer wg.Done()

		for ctx.Err() == nil {
			chunk := make([]byte, memChunkMiB*1024*1024)

			// Prevent the memory from being optimized away
			for i := 0; i u003c len(chunk); i++ {
				chunk[i] = byte(i % 256)
			}
			chunks = append(chunks, chunk)

			fmt.Printf("Allocated %d MB of memory\n", memChunkMiB)
			time.Sleep(1 * time.Second)
		}
	}()

	<-ctx.Done()
	fmt.Println("Received termination signal. Initiating shutdown...")

	wg.Wait()
	fmt.Println("Shutdown complete.")
}

// Jail this program.

