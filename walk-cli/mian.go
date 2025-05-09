package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type multiFlag []string

func (m *multiFlag) String() string {
	return fmt.Sprintf("%v", *m)
}

func (m *multiFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

type config struct {
	exts    []string
	size    int64
	list    bool
	del     bool
	archive string
	wLog    io.Writer
}

var (
	f   = os.Stdout
	err error
)

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filterOut(path, cfg.exts, cfg.size, info) {
			return nil
		}
		if cfg.list {
			return listFile(path, out)
		}
		if cfg.archive != "" {
			if err := archiveFile(cfg.archive, root, path); err != nil {
				return err
			}
		}
		if cfg.del {
			return delFile(path, delLogger)
		}
		return listFile(path, out)
	})
}

func main() {
	var extensions multiFlag
	flag.Var(&extensions, "ext", "File extensions to be filter out.")
	root := flag.String("root", ".", "Root directory to start.")
	logFile := flag.String("log", "", "Log deletes to this file")
	size := flag.Int64("size", 0, "Minimum File size to filter out.")
	list := flag.Bool("list", false, "List files only.")
	del := flag.Bool("del", false, "Delete files only.")
	archive := flag.String("archive", "", "Archive file.")
	flag.Parse()

	c := config{
		exts:    extensions,
		size:    *size,
		list:    *list,
		del:     *del,
		archive: *archive,
		wLog:    f,
	}

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
