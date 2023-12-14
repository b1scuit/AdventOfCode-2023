package lineswapper

import "strings"

type LineFunc func(*LineSwapper) LineFunc

type LineSwapper struct {
	Pos int

	Buf string

	State LineFunc
}

type LineOption func(*LineSwapper)

func WithInput(s string) LineOption {
	return func(ls *LineSwapper) {
		ls.Buf = s
	}
}

func New(opts ...LineOption) *LineSwapper {
	l := LineSwapper{
		State: StartHell,
	}

	for _, f := range opts {
		f(&l)
	}

	return &l
}

func (ls *LineSwapper) Do() string {
	state := ls.State

	for state := state; state != nil; {
		state = state(ls)
	}

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

	return StartHell
}

func isTwo(s string) bool {
	return strings.HasPrefix(s, "two")
}
func LexTwo(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "two", "2", 1)

	return StartHell
}

func isThree(s string) bool {
	return strings.HasPrefix(s, "three")
}
func LexThree(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "three", "3", 1)

	return StartHell
}

func isFour(s string) bool {
	return strings.HasPrefix(s, "four")
}
func LexFour(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "four", "4", 1)

	return StartHell
}
func isFive(s string) bool {
	return strings.HasPrefix(s, "five")
}
func LexFive(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "five", "5", 1)

	return StartHell
}
func isSix(s string) bool {
	return strings.HasPrefix(s, "six")
}
func LexSix(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "six", "6", 1)

	return StartHell
}
func isSeven(s string) bool {
	return strings.HasPrefix(s, "seven")
}
func LexSeven(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "seven", "7", 1)

	return StartHell
}
func isEight(s string) bool {
	return strings.HasPrefix(s, "eight")
}
func LexEight(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "eight", "8", 1)

	return StartHell
}
func isNine(s string) bool {
	return strings.HasPrefix(s, "nine")
}
func LexNine(l *LineSwapper) LineFunc {
	l.Buf = strings.Replace(l.Buf, "nine", "9", 1)

	return StartHell
}
