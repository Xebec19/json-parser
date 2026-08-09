package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Xebec19/json-parser/internal"
	"github.com/Xebec19/json-parser/pkg/tokens"
)

func main() {

	argv := os.Args

	if len(argv) < 2 {
		slog.Error("Invalid args", slog.Any("arguments", argv))
		os.Exit(1)
	}

	fileName := argv[1]

	slog.Info("received file name for json parsing", slog.String("file", fileName))

	dir, err := os.Getwd()
	if err != nil {
		slog.Error("failed to get working directory", slog.String("error", err.Error()))
		os.Exit(1)
	}

	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		slog.Error("failed to open file! could not read json!", slog.String("error", err.Error()))
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var token []tokens.Token

	for scanner.Scan() {
		token = append(token, internal.Lexer(scanner.Text())...)
	}

	if err := scanner.Err(); err != nil {
		slog.Error(fmt.Sprintf("Failed to read content of file %s", fileName), slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("Tokens", slog.Any("token", token))
}
