package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/*
	//パイプライン(バッチ処理形式)
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i,v := range values {
			multipliedValues[i] = v * multiplier
		}

		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i,v := range values {
			addedValues[i] = v + additive
		}

		return addedValues
	}


	ints := []int{1, 2, 3, 4, 5, 6, 7}

	for _,v := range multiply(add(multiply(ints,2), 1),2) {  //<-ここがパイプライン結合している部分
		fmt.Println(v)
	}

	fmt.Println("")


	//パイプライン(ストリーム処理形式)
	multiply2 := func(value, multiplier int) int {
		return value * multiplier
	}

	add2 := func(value, additive int) int {
		return value + additive
	}


	ints2 := []int{1, 2, 3, 4, 5, 6, 7}

	for _,v := range ints2 {
		println(multiply2(add2(multiply2(v,2), 1),2) )
	}

	 */

	//パイルラインをチャネルで構築
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
					case <-done:
						return
					case intStream <- i:
				}
			}
		}()

		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
		) <-chan int {
			multipliedStream := make(chan int)
			go func() {
				defer close(multipliedStream)
				for i := range intStream {
					select {
						case <-done:
							return
						case multipliedStream <- i * multiplier:
					}
				}
			}()

			return multipliedStream
	}


	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
		) <- chan int {
			addedStream := make(chan int)
			go func() {
				defer close(addedStream)
				for i := range intStream {
					select {
						case <-done:
							return
						case addedStream <- i + additive:
					}
				}
			}()

			return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}


	fmt.Println("")


	//便利なジェネレータ関数
	/*
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <- chan interface{} {
			valueStream := make(chan interface{})
			go func() {
				defer close(valueStream)
				for {
					for _, v := range values {
						select {
							case <- done:
								return
							case valueStream <- v:
						}
					}
				}
			}()

			return valueStream
	}

	 */

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
		) <- chan interface{} {
			valueStream := make(chan interface{})
			go func() {
				defer close(valueStream)
				for {
					select {
						case <-done:
							return
						case valueStream <- fn():
					}
				}
			}()

			return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
		) <- chan interface{} {
			takeStream := make(chan interface{})
			go func() {
				defer close(takeStream)
				for i := 0; i < num; i++ {
					select {
						case <-done:
							return
						case takeStream <- <- valueStream:
					}
				}
			}()

			return takeStream
	}

	//乱数を指定数分だけ生成する

	done2 := make(chan interface{})
	defer close(done2)

	rand := func() interface{} {return rand.Int() }

	for num := range take(done, repeatFn(done, rand),100) {
		fmt.Println(num)
	}


	//ファンアウトとファンインの例
	/*
	rand := func() interface{} {return rand.Intn(50000000)}

	done3 := make(chan interface{})
	defer close(done3)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done,rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done,randIntStream),10) {
		fmt.Printf("\t%d\n",prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
	*/

}

