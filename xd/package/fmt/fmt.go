package main

import "fmt"

const (
	// \033 is a octal number
	// [
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func main() {
	fmt.Println(Red + "This is Red" + Reset + " " + Green + "This is Green")
}
