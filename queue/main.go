package main

import (
	"fmt"

	"github.com/golang-queue/queue"
)

func mainCh(nTask int) {
	chs := make(chan string, nTask)

	for i := 0; i < nTask; i++ {
		go JobCh(i, chs)
	}

	for i := 0; i < nTask; i++ {
		fmt.Println(<-chs)
	}
}

func mainQueue(nTask int) {
	chs := make(chan string, nTask)
	q := queue.NewPool(nTask)
	defer q.Release()

	for i := 0; i < nTask; i++ {
		go q.QueueTask(JobQueue(i, chs))
	}

	for i := 0; i < nTask; i++ {
		fmt.Println(<-chs)
	}
}

func main() {
	nTask := 5
	mainCh(nTask)
	fmt.Printf("\nUsing Queue...\n")
	mainQueue(nTask)
}
