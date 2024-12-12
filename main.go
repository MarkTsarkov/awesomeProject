// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("done")
	var animals []string = []string{"dog", "cat", "bird", "seal", "cow", "platypus"}
	var wg sync.WaitGroup

	workC := 3
	jobC := 6
	lineID := 1

	jobs := make(chan int, jobC)
	res := make(chan string, workC)

	for w := 1; w <= workC; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			wrk(w, animals, jobs, res)
		}(w)
	}

	for j := 0; j < jobC; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Printf("line %v, %s\n", lineID, r)
		lineID++
	}
}

func wrk(worker int, animals []string, jobs <-chan int, res chan<- string) {
	for j := range jobs {
		line := "thread " + strconv.Itoa(worker) + ", animal " + "\"" + animals[j] + "\""
		time.Sleep(time.Second * 1)
		res <- line
	}
}
