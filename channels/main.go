package main

import (
	"fmt"
	"time"
)

func woker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs) //Example: 5 clients in restaurant
	result := make(chan int, numJobs)

	//Example : 3 workers
	for w := 1; w <= 3; w++ {
		go woker(w, jobs, result)
	}
	// we insert numjobs in jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
	close(result)

}
