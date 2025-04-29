package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(r io.Reader, countLines bool, countBytes bool) (int, int) {
	var lines, words, bytes int
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if countLines {
			lines++
		}
		if !countLines {
			scanner := bufio.NewScanner(strings.NewReader(line))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				words++
			}
		}
		bytes += len(line)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %v\n", err)
			break
		}

	}
	if countLines && countBytes {
		return lines, bytes
	} else if countLines {
		return lines, 0
	} else if countBytes {
		return words, bytes
	}
	return words, 0
}

func main() {
	countLines := flag.Bool("l", false, "count lines")
	countBytes := flag.Bool("b", false, "count bytes")
	fmt.Println(count(os.Stdin, *countLines, *countBytes))
}
