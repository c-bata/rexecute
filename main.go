package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fi, err := os.Stat(s.Text())
		if err != nil {
			return
		}
		mtime := fi.ModTime()
		fmt.Println(mtime)
	}
}
