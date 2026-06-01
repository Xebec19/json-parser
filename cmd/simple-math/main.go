package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Input expressions")

	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	fmt.Println(expr)
}
