package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {

	f := flag.Int("flag1", 0, "flag 1")
	flag.Parse()
	fmt.Println(f)

	fmt.Println("os.Args: ", os.Args)

	fmt.Println("サーバルちゃん")

}
