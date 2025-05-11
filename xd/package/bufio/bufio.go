package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println(reader)
}
