package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	fibPrev := 1
	fib := 1

	for i := 2; i < n; i++ {
		tmp := fib
		fib = fib + fibPrev
		fibPrev = tmp
	}

	return fib
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}
