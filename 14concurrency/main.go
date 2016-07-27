package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	runtime.GOMAXPROCS(1)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		defer wg.Done()

		for a := 100; a >= 0; a-- {
			fmt.Printf("X:%d", a)
		}

		wg.Done()
	}()

	go func() {

		for i := 0; i < 100; i++ {
			fmt.Printf("Y:%d", i)
		}

	}()

	wg.Done()

	wg.Wait()

	fmt.Println("terminating")

}
