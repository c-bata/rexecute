package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const interval int = 1

// RunCmd is a function to run a external command.
func RunCmd(c []string) {
	cmd := exec.Command(c[0], c[1:]...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", out)
}

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
				fmt.Printf("%s is changed!\n", k)
				RunCmd(os.Args[1:])
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
