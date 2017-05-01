package main

import (
	_ "expvar"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	runtime.SetBlockProfileRate(1)
	var m sync.Mutex
	ch := make(chan int)

	outFile, _ := os.Create("./trace.out")
	trace.Start(outFile)
	time.AfterFunc(time.Second*5, func() {
		trace.Stop()
	})

	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			fmt.Println("got lock")
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			<-ch
		}()
	}

	go func() {
		for {
			time.Sleep(time.Second)
			os.Hostname()
		}
	}()

	http.ListenAndServe(":3000", nil)
}
