package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	files := make(map[string]time.Time)
	for s.Scan() {
		f := s.Text()
		fi, err := os.Stat(f)
		if err != nil {
			return
		}
		files[f] = fi.ModTime()
		fmt.Println(f)
	}

	for {
		for k, v := range files {
			fi, err := os.Stat(k)
			if err != nil {
				return
			}
			mtime := fi.ModTime()
			if mtime.After(v) {
				// TODO: Run cmd.
				fmt.Println(k)
				files[k] = mtime
			}
		}
		time.Sleep(1 * time.Second)
	}
}
