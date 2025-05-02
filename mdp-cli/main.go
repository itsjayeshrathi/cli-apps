package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

const (
	defaultTemplate = `<!DOCTYPE html>
 <html>
 <head>
 <meta http-equiv="content-type" content="text/html; charset=utf-8">
 <title>{{ .Title }}</title>
 </head>
 <body>
 {{ .Body }}
 </body>
 </html>
 `
)

type content struct {
	Title string
	Body  template.HTML
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	var markdownBuf bytes.Buffer
	if err := goldmark.Convert(input, &markdownBuf); err != nil {
		panic(fmt.Sprintf("failed to convert markdown: %v", err))
	}

	body := bluemonday.UGCPolicy().Sanitize(markdownBuf.String())

	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}
	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer
	// Execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil

}

func saveHTML(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

func run(fileName string, tFname string, out io.Writer, skipPreview bool) error {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}
	temp, err := os.CreateTemp("./testdata", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprint(out, outName)
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	if skipPreview {
		return nil
	}
	defer os.Remove(outName)
	return preview(outName)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}
	cParams = append(cParams, fname)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	return exec.Command(cPath, cParams...).Run()
}

func main() {
	fileName := flag.String("file", "", "Markdown file to preivew")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*fileName, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
