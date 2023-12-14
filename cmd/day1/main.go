package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
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
		swap := &LineSwapper{State: StartHell}
		swappedLine := swap.Input(line).Process().Output()
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

type LineFunc func(*LineSwapper) LineFunc

type LineSwapper struct {
	Pos int

	Buf string

	State LineFunc
}

func (ls *LineSwapper) Input(s string) *LineSwapper {
	ls.Buf = s

	return ls
}

func (ls *LineSwapper) Process() *LineSwapper {
	state := ls.State

	for state := state; state != nil; {
		state = state(ls)
	}

	return ls
}

func (ls *LineSwapper) Output() string {
	return ls.Buf
}

// Lots of functions
func (ls *LineSwapper) IsEof() bool {
	return (ls.Pos >= len(ls.Buf))
}

func (ls *LineSwapper) Inc() {
	ls.Pos++
}

func (ls *LineSwapper) ResetPos() {
	ls.Pos = 0
}

func (ls *LineSwapper) InputToEnd() string {
	return ls.Buf[ls.Pos:]
}

// Oh god, i'm sorry John
// Guess you got your overengineered example eventuially ;-)

func StartHell(l *LineSwapper) LineFunc {
	for {
		restOfLine := l.InputToEnd()

		if isOne(restOfLine) {
			return LexOne(l)
		}

		if isTwo(restOfLine) {
			return LexTwo(l)
		}

		if isThree(restOfLine) {
			return LexThree(l)
		}

		if isFour(restOfLine) {
			return LexFour(l)
		}

		if isFive(restOfLine) {
			return LexFive(l)
		}
		if isSix(restOfLine) {
			return LexSix(l)
		}

		if isSeven(restOfLine) {
			return LexSeven(l)
		}

		if isEight(restOfLine) {
			return LexEight(l)
		}

		if isNine(restOfLine) {
			return LexNine(l)
		}

		if isZero(restOfLine) {
			return LexZero(l)
		}
		if l.IsEof() {
			break
		}

		l.Inc()
	}

	return nil
}

func isOne(s string) bool {
	return strings.HasPrefix(s, "one")
}

func LexOne(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "one", "1", 1)
	l.ResetPos()

	return StartHell
}

func isTwo(s string) bool {
	return strings.HasPrefix(s, "two")
}
func LexTwo(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "two", "2", 1)
	l.ResetPos()

	return StartHell
}

func isThree(s string) bool {
	return strings.HasPrefix(s, "three")
}
func LexThree(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "three", "3", 1)
	l.ResetPos()

	return StartHell
}

func isFour(s string) bool {
	return strings.HasPrefix(s, "four")
}
func LexFour(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "four", "4", 1)
	l.ResetPos()

	return StartHell
}
func isFive(s string) bool {
	return strings.HasPrefix(s, "five")
}
func LexFive(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "five", "5", 1)
	l.ResetPos()

	return StartHell
}
func isSix(s string) bool {
	return strings.HasPrefix(s, "six")
}
func LexSix(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "six", "6", 1)
	l.ResetPos()

	return StartHell
}
func isSeven(s string) bool {
	return strings.HasPrefix(s, "seven")
}
func LexSeven(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "seven", "7", 1)
	l.ResetPos()

	return StartHell
}
func isEight(s string) bool {
	return strings.HasPrefix(s, "eight")
}
func LexEight(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "eight", "8", 1)
	l.ResetPos()

	return StartHell
}
func isNine(s string) bool {
	return strings.HasPrefix(s, "nine")
}
func LexNine(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "nine", "9", 1)
	l.ResetPos()

	return StartHell
}
func isZero(s string) bool {
	return strings.HasPrefix(s, "zero")
}

func LexZero(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "zero", "0", 1)
	l.ResetPos()

	return StartHell
}
