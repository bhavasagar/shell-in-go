package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(input)
}
