package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

const (
	header = `<!DOCTYPE html>
 <html>
 <head>
 <meta http-equiv="content-type" content="text/html; charset=utf-8">
 <title>Markdown Preview Tool</title>
 </head>
 <body>
 `
	footer = `
 </body>
 </html>
 `
)

func parseContent(input []byte) []byte {
	var markdownBuf bytes.Buffer
	if err := goldmark.Convert(input, &markdownBuf); err != nil {
		panic(fmt.Sprintf("failed to convert markdown: %v", err))
	}

	safeHTML := bluemonday.UGCPolicy().Sanitize(markdownBuf.String())

	var finalBuf bytes.Buffer
	finalBuf.WriteString(header)
	finalBuf.WriteString(safeHTML)
	finalBuf.WriteString(footer)

	return finalBuf.Bytes()
}

func saveHTML(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

func run(fileName string) error {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	htmlData := parseContent(input)
	temp, err := os.CreateTemp("./testdata", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	return saveHTML(outName, htmlData)
}

func main() {
	fileName := flag.String("file", "", "Markdown file to preivew")
	flag.Parse()
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
