package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"os/exec"
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
				files[k] = mtime
				fmt.Println(k)

				cmd := exec.Command(os.Args[1], os.Args[2:]...)
				out, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("%s", out)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
