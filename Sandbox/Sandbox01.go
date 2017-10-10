package main

import (
	"os"
	"fmt"
	"path/filepath"
	"os/exec"
)

func main() {

	// 取得
	p, _ := os.Getwd()
	fmt.Println(p)

	// 再帰なし
	fmt.Println("No Recursion-----")
	files, _ := filepath.Glob(p + "/*")
	for _, f := range files {
		//fmt.Println(f)

		//別プログラム起動
		//out, _ := exec.Command("Exec Sandbox02_go.exe", f).Output()
		//fmt.Println(len(out))
		//fmt.Println(err.)
		//stdinObj := os.Stdin
		cmd := exec.Command("Sandbox02_go.exe", f)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//func Run(name string, argv, envv []string, dir string, stdin, stdout, stderr int) (p *Cmd, err os.Error)

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println("------------------")

	//再帰あり
	//fmt.Println("Recursion")
	//filepath.Walk(p, visit)
	//fmt.Println("------------------")




}

func visit(path string, info os.FileInfo, err error) error {

	fmt.Println(path)
	return nil
}


