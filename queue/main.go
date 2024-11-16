package main

import (
	"fmt"
)

func main() {
	nTask := 5

	chs := make(chan string)
	for i := 0; i < nTask; i++ {
		go JobCh(i, chs)
	}

	for i := 0; i < nTask; i++ {
		fmt.Println(<-chs)
	}
}
