package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countsForEachFile := make(map[string]map[string]int)
	cs := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, cs)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsForEachFile[file] = make(map[string]int)
			countLines(f, countsForEachFile[file])
			f.Close()
		}
	}
	for file, counts := range countsForEachFile {
		for _, count := range counts {
			if count > 1 {
				fmt.Println(file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
