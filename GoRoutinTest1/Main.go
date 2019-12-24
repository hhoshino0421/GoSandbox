package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	//読み込み並列
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

	fmt.Println("")

	//書き込み並列
	newRandStream := func(done2 <-chan interface{} ) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
					case randStream <- rand.Int():
						case <-done2:
							return
				}
			}

		}()

		return randStream
	}

	done2 := make(chan interface{})
	randStream := newRandStream(done2)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	close(done2)

	//処理が実行中であることをシミュレート
	time.Sleep(1 * time.Second)

	fmt.Println("")

	//ORチャネル
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
					case <- channels[0]:
						case <- channels[1]:
				}
			default:
				select {
					case <-channels[0]:
						case <-channels[1]:
							case <- channels[2]:
								case <- or(append(channels[3:],orDone)...):
				}
			}

		}()

		return orDone

	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		} ()

		return c
	}

	start := time.Now()
	<- or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		)

	fmt.Printf("done after %v", time.Since(start))

	fmt.Println("")

	//エラーハンドリング
	type Result struct {
		Error error
		Response *http.Response
	}

	checkStatus := func(done3 <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
					case <- done3:
						return
						case results <- result:
				}
			 }
		}()

		return results
	}

	done3 := make(chan interface{})
	defer close(done3)

	errCount := 0
	urls := []string{"a", "https://www.google.co.jp", "b", "c", "d"}
	for result := range checkStatus(done3, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if (errCount >= 3) {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)

	}


}


