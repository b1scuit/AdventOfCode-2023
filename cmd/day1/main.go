package main

import (
	"bytes"
	"flag"
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	// Some random shite
	now := time.Now()
	slog.Info("Starting Day 1 AoC")

	// Parse input
	var fileName string
	flag.StringVar(&fileName, "file", "./assets/day1/examples.txt", "File path for the task")
	flag.Parse()

	slog.Info("Parsing File", slog.String("file_path", fileName))

	f, err := os.Open(fileName)
	if err != nil {
		slog.Error("Error opening file", slog.Any("error", err))
		return
	}

	b := bytes.Buffer{}

	io.ReadAll(f)
	if _, err := f.Read(&b); err != nil {
		slog.Error("Error reading file contents", slog.Any("error", err))
	}

	// Finish up
	slog.Info("Completed Day 1 AoC", slog.Duration("took", time.Since(now)))
}
