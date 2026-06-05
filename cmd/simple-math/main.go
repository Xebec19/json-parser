package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/Xebec19/json-parser/internal"
)

func main() {

	fmt.Println("Input expressions")

	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	tokens, err := internal.Lexer(expr)
	if err != nil {
		slog.Error("parse token", "error", err)
		os.Exit(1)
	}

	fmt.Println(tokens)
}
