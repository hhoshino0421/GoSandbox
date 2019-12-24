package main

import "fmt"



func main() {

	//パイプライン
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

	for _,v := range add(multiply(ints,2), 1) {  //<-ここがパイプライン結合している部分
		fmt.Println(v)
	}



}

