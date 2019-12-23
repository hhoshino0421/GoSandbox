package main

import "fmt"

var data int
var errorcnt int
var LoopCnt int

func main() {

	data = 0
	errorcnt = 0
	LoopCnt = 100000000


	threadtest(LoopCnt)

	fmt.Println(errorcnt)
	
}

func threadtest(loop_cnt int) {

	for i := 0; i< loop_cnt; i++ {

		data = 0

		go func() {
			data++
		} ()

		go func() {
			data--
		} ()

		if data != 0 {
			errorcnt++
		}


	}

}
