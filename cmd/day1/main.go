package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/b1scuit/adventofcode-2023/lineswapper"
)

// fivesevenfour9jslninesevenjtttt7oneightssr is the current problem
// should end up as 58 but instead comes out as 51 as numbers are parsed L->R
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
	var totalMutex sync.Mutex
	var total int = 0
	var wg sync.WaitGroup
	for _, line := range fileLines {
		wg.Add(1)
		go func(wg *sync.WaitGroup, line string) {
			defer wg.Done()
			// Walk through the line and pick out the numbers

			line1 := LexLine(line)

			lineValue := Walk(line1) + WalkBack(line1)

			value, err := strconv.Atoi(lineValue)
			if err != nil {
				slog.Error("Error converting string", slog.Any("error", err))
			}
			totalMutex.Lock()
			total = total + value
			totalMutex.Unlock()
		}(&wg, line)
	}

	wg.Wait()

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

func LexLine(s string) string {
	return lineswapper.New(lineswapper.WithInput(s)).Do()
}
