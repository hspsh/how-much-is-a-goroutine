// SPDX-License-Identifier: Apache-2.0
package how_much_is_a_goroutine

import (
	"flag"
	"sync"
	"time"

	"github.com/HdrHistogram/hdrhistogram-go"
)

var index = flag.Int("index", 10, "number of index iterations")

func DoIrrelevantWork(wg *sync.WaitGroup, histogram *hdrhistogram.Histogram) {
	start := time.Now()
	GetFibonacciByIndex(*index)
	end := time.Now()
	elapsed := end.Sub(start)
	histogram.RecordValue(elapsed.Nanoseconds())
	wg.Done()
}
