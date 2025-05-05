package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func filterOut(path string, ext string, size int64, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < size {
		return true
	}
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}
	return false
}
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}
