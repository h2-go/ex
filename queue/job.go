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

func JobCh(n int, chs chan string) {
	sleepSomeTime(n)
	chs <- fmt.Sprintf("job %d done", n)
}
