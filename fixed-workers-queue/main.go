// SPDX-License-Identifier: Apache-2.0
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/HdrHistogram/hdrhistogram-go"
	how_much_is_a_goroutine "github.com/hspsh/how-much-is-a-goroutine"
)

var workers = flag.Int("workers", 10, "number of workers doing irrelevant work")
var iterations = flag.Int64("iterations", 1000000, "number of iterations to perform and goroutines to create")

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	wg.Add(int(*iterations) + 1)
	var queue = make(chan int64, 100)
	histogram := hdrhistogram.New(1, 500, 1)

	var i int64
	go func(count int64) {
		for i = 0; i < count; i++ {
			queue <- i
		}
		wg.Done()
	}(*iterations)

	for i := 0; i < *workers; i++ {
		go runWorker(queue, &wg, histogram)
	}

	wg.Wait()
	fmt.Println(i)
	fmt.Println(*iterations)
	close(queue)
	histogram.PercentilesPrint(os.Stderr, 2, 1)
	os.Stderr.Sync()
}

func runWorker(queue chan int64, wg *sync.WaitGroup, histogram *hdrhistogram.Histogram) {
	for {
		select {
		case _, ok := <-queue:
			if ok {
				how_much_is_a_goroutine.DoIrrelevantWork(wg, histogram)
			} else {
				break
			}
		}
	}
}
