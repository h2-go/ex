package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepSomeTime(n int) {
	sleepTime := time.Duration(rand.Intn(10))
	fmt.Printf("job %2d start sleeping for %2d seconds\n", n, sleepTime)
	time.Sleep(sleepTime * time.Second)
}

func job(n int, chs chan string) {
	sleepSomeTime(n)
	chs <- fmt.Sprintf("job %d done", n)
}

func main() {
	nTask := 5

	chs := make(chan string)
	for i := 0; i < nTask; i++ {
		go job(i, chs)
	}
	for i := 0; i < nTask; i++ {
		fmt.Println(<-chs)
	}
}
