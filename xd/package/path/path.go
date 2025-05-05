package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	path := flag.String("path", ".", "File Path for manipulation.")
	flag.Parse()
	base := filepath.Dir(*path)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	fmt.Println(base)
}
