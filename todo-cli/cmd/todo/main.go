package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/itsjayeshrathi/todo-cli"
)

var todoFileName = ".todo.json"

func getTask(r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}
	var tasks []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			tasks = append(tasks, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, fmt.Errorf("No tasks provided	")
	}
	return tasks, nil
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	if _, err := os.Stat(todoFileName); os.IsNotExist(err) {
		file, err := os.Create(todoFileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create file %s: %v\n", todoFileName, err)
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to initialize file content: %v\n", err)
			os.Exit(1)
		}
	}

	add := flag.Bool("add", false, "Add task to the Todo list")
	list := flag.Bool("list", false, "List all tasks")
	ul := flag.Bool("ul", false, "List of all uncompleted tasks.")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be deleted")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, "Error loading tasks: ", err)
		os.Exit(1)

	}

	switch {
	case *list:
		if *ul {
			l.PrintIncomplete()
		} else {
			fmt.Print(l)
		}

	case *complete > 0:

		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, "Error finishing task: ", err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks: ", err)
			os.Exit(1)
		}
		fmt.Println("Task completed sucessfully.")

	case *delete > 0:

		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting task: ", err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks: ", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted sucessfully.")
	case *add:
		fmt.Println("after writing all tasks, press CTRL+D")
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, item := range t {
			l.Add(item)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks: ", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)

	}
}
