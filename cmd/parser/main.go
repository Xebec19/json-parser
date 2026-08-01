package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func main() {

	argv := os.Args

	if len(argv) < 2 {
		slog.Error("Invalid args", slog.Any("arguments", argv))
		os.Exit(1)
	}

	fileName := argv[1]

	slog.Info("received file name for json parsing", slog.String("file", fileName))

	file, err := os.Open(fileName)
	if err != nil {
		slog.Error("failed to open file! could not read json!", slog.String("error", err.Error()))
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		slog.Info(fmt.Sprintf("Payload received from file %s", fileName), slog.String("payload", scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		slog.Error(fmt.Sprintf("Failed to read content of file %s", fileName), slog.String("error", err.Error()))
	}
}
