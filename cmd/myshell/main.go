package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the input: ", err)
		os.Exit(1)
	}

	os := runtime.GOOS
	max_cap := len(input) - 1
	if os == "windows" {
		max_cap--
	}
	cmd := input[:max_cap]

	switch cmd {
	case "TEST_CMD":
		fmt.Println("Executing TEST_CMD")
	default:
		fmt.Println(cmd + ": Command not found")
	}

}
