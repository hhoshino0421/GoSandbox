package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}

			}
		}()

		return terminated

	}

	done := make(chan interface{})
	terminated := doWork(done, nil)


	go func() {
		//１秒後に操作をキャンセルする
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork gorutine...")
		close(done)
	}()


	<- terminated
	fmt.Println("Done.")

}

