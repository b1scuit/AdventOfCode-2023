package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"time"
	"unicode"

	"github.com/b1scuit/adventofcode-2023/lineswapper"
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
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fileLines := []string{}
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	// Now the file is something neat and simple to work with, we can sort things out
	var total int = 0
	for _, line := range fileLines {
		swappedLine := lineswapper.New(lineswapper.WithInput(line)).Do()

		// Walk through the line and pick out the numbers
		lineValue := Walk(swappedLine) + WalkBack(swappedLine)

		value, err := strconv.Atoi(lineValue)
		if err != nil {
			slog.Error("Error converting string", slog.Any("error", err))
		}

		total = total + value
	}

	slog.Info("Calculation Total", slog.Int("total", total))

	// Finish up
	slog.Info("Completed Day 1 AoC", slog.Duration("took", time.Since(now)))
}

func Walk(s string) string {
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			return string(s[i])
		}
	}

	return ""
}

func WalkBack(s string) string {
	b := []byte(s)
	slices.Reverse(b)
	return Walk(string(b))
}

// Overkill, but fun
