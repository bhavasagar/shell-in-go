package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading the input: ", err)
			os.Exit(1)
		}

		runtime := runtime.GOOS
		max_cap := len(input) - 1
		if runtime == "windows" {
			max_cap--
		}

		input_eval := strings.Split(input[:max_cap], " ")
		if len(input_eval) == 0 {
			continue
		}

		cmd := input_eval[0]
		args := input_eval[1:]

		switch cmd {
		case "TEST_CMD":
			fmt.Println("Executing TEST_CMD")
		case "exit":
			exit_status, _ := strconv.Atoi(args[0])
			os.Exit(exit_status)
		default:
			fmt.Println(cmd + ": command not found")
		}
	}
}
