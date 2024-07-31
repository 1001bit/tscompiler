package main

import (
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, path := range os.Args {
		wg.Add(1)
		go func() {
			RecursiveCompileIn(path)
			wg.Done()
		}()
	}

	wg.Wait()
}
