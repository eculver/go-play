package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	defaultNumDsts     = 100
	defaultConcurrency = 5
)

// keep track tof the total amount of time spent just for copying so that we
// can report on an "improvement" metric that is just the difference between
// the total copy time and the actual time it took to complete
var totalCopyTime = time.Duration(0)

// ncopy is a simulation of a coordinated file copy mechanism from a single source
// node to N destination nodes. The only assumption is that you have a copyFile
// method that does the actual copying of the file from src to dst. For the sake
// of this example copyFile blocks for a random amount of time.
func copyFile(src, dst int) {
	// sleep for a random duration between 3 and 15 seconds
	secs := rand.Intn(12) + 3
	durs := time.Duration(secs) * time.Second
	totalCopyTime = totalCopyTime + durs
	fmt.Printf("start copy from %d to %d (%d secs)\n", src, dst, secs)
	time.Sleep(durs)
	fmt.Printf("done copying from %d to %d\n", src, dst)
}

func main() {
	var (
		wg         sync.WaitGroup
		resultLock sync.RWMutex

		numDsts     = flag.Int("num-dsts", defaultNumDsts, "Number of destinations to send to")
		concurrency = flag.Int("concurrency", defaultConcurrency, "Max number of concurrent sends from a source")
	)
	flag.Parse()

	results := make(map[int][]int)
	seeds := make(chan int)
	sends := make(chan int)
	dsts := make([]int, *numDsts)
	// generate a list of integer to use as destination IDs
	// node 0 will act as the initial seed so start at 1
	for i := 1; i <= *numDsts; i++ {
		dsts[i-1] = i
	}

	// seed makes a source available for x concurrent sends
	seed := func(src, x int) {
		for i := 0; i < x; i++ {
			seeds <- src
		}
	}

	saveResult := func(src, dst int) {
		resultLock.Lock()
		defer resultLock.Unlock()
		if _, ok := results[src]; !ok {
			results[src] = []int{dst}
			return
		}
		results[src] = append(results[src], dst)
	}

	// send initiates a send to dst from the next available source
	sendTo := func(dst int) {
		src := <-seeds
		copyFile(src, dst)
		saveResult(src, dst)
		wg.Done()
		// dst is now available for sending, add number of sends matching the concurrency factor
		// from it as a new source
		go seed(dst, *concurrency)
		// and we also put the src back out there too since it's available again
		go seed(src, 1)
	}

	// mark this as our "start" time since everything else is setup
	start := time.Now()

	// populate the channel of sends that need to happen
	go func() {
		for _, dst := range dsts {
			sends <- dst
		}
		close(sends)
	}()

	// populate the initial "set" of sources
	go seed(0, *concurrency)

	// block until all dsts have been sent to
	wg.Add(*numDsts)
	for dst := range sends {
		go sendTo(dst)
	}
	wg.Wait()

	timing := time.Now().Sub(start)

	// spit out some values humans may care about
	fmt.Printf("\n\nresults:\n")

	timingDiffAbs := totalCopyTime - timing
	timingDiffPct := (float64(timingDiffAbs) / float64(totalCopyTime)) * 100
	max := 0
	maxSrc := -1
	for src, dsts := range results {
		if len(dsts) > max {
			max = len(dsts)
			maxSrc = src
		}
		fmt.Printf("%d sent to %d dsts: %v\n", src, len(dsts), dsts)
	}
	fmt.Println("--------")
	fmt.Printf("max sends: %d sent to %d nodes\n", maxSrc, max)
	fmt.Printf("improvement over baseline: %v (%.2f%%)\n", timingDiffAbs, timingDiffPct)
}
