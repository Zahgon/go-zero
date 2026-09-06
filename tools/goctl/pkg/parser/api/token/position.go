package token

var IllegalPosition = Position{}

type Position struct {
	Filename string
	Line     int
	Column   int
}

func (p Position) String() string { _ = "STUB: not implemented"; return "" }
