package scanner

import (
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"
)

const (
	initMode mode = iota

	documentHalfOpen
	documentOpen
	documentHalfClose
	documentClose

	stringOpen
	stringClose
)

type mode int

type Scanner struct {
	filename string
	size     int

	data         []rune
	position     int
	readPosition int
	ch           rune

	lines []int
}

func (s *Scanner) NextToken() (token.Token, error) {
	_ = "STUB: not implemented"
	return *new(token.Token), nil
}

func (s *Scanner) newToken(tp token.Type) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) readRune() { _ = "STUB: not implemented"; return }

func (s *Scanner) peekRune() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) scanString(delim rune, tp token.Type) (token.Token, error) {
	_ = "STUB: not implemented"
	return *new(token.Token), nil
}

func (s *Scanner) scanAt() (token.Token, error) {
	_ = "STUB: not implemented"
	return *new(token.Token), nil
}

func (s *Scanner) scanIntOrDuration() token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanDuration(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanNanosecond(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanMicrosecond(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanMillisecondOrMinute(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanMillisecond(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanSecond(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanMinute(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanHour(bgPos int) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) illegalToken() token.Token { _ = "STUB: not implemented"; return *new(token.Token) }

func (s *Scanner) scanIdent() token.Token { _ = "STUB: not implemented"; return *new(token.Token) }

func (s *Scanner) scanLetterSet() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanLineComment() token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func (s *Scanner) scanDocument() (token.Token, error) {
	_ = "STUB: not implemented"
	return *new(token.Token), nil
}

func (s *Scanner) assertExpected(actual token.Type, expected ...token.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) assertExpectedString(actual string, expected ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) positionAt() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *Scanner) newPosition(position int) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *Scanner) lineCount() int { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) skipWhiteSpace() { _ = "STUB: not implemented"; return }

func (s *Scanner) isDigit(b rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isLetter(b rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isIdentifierLetter(b rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isWhiteSpace(b rune) bool { _ = "STUB: not implemented"; return false }

func MustNewScanner(filename string, src interface{}) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

func NewScanner(filename string, src interface{}) (*Scanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readData(filename string, src interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
