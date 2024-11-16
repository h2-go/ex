package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func sleepSomeTime(n int) time.Duration {
	sleepTime := time.Duration(rand.Intn(10))
	fmt.Printf("job %d start sleeping for %ds\n", n, sleepTime)
	time.Sleep(sleepTime * time.Second)
	return sleepTime
}

func JobCh(n int, chs chan string) {
	sleepTime := sleepSomeTime(n)
	chs <- fmt.Sprintf("job %d done in %ds", n, sleepTime)
}

func JobQueue(n int, chs chan string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		sleepTime := sleepSomeTime(n)
		chs <- fmt.Sprintf("job %d done in %ds", n, sleepTime)
		return nil
	}
}
