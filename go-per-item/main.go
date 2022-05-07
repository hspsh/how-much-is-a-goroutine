package main

import (
	"flag"
	"os"
	"sync"

	"github.com/HdrHistogram/hdrhistogram-go"
	"github.com/hspsh/how-much-is-a-goroutine"
)

var iterations = flag.Int64("iterations", 1000, "number of iterations to perform and goroutines to create")

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	wg.Add(int(*iterations))
	var i int64
	histogram := hdrhistogram.New(1, 500, 1)
	for i = 0; i < *iterations; i++ {
		go how_much_is_a_goroutine.DoIrrelevantWork(&wg, histogram)
	}
	wg.Wait()
	histogram.PercentilesPrint(os.Stderr, 2, 1)
	os.Stderr.Sync()
}
