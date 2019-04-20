package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const intervalDuration = 1 * time.Second

func run(c []string) {
	cmd := exec.Command(c[0], c[1:]...)
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "[rexecute] error:", err)
		os.Exit(1)
	}
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
				_, _ = fmt.Fprintln(os.Stderr, "[rexecute] changed:", k)
				run(os.Args[1:])
			}
		}
		time.Sleep(intervalDuration)
	}
}
