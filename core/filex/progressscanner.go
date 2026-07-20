package filex

import "gopkg.in/cheggaaa/pb.v1"

type (
	Scanner interface {
		Scan() bool

		Text() string
	}

	progressScanner struct {
		Scanner
		bar *pb.ProgressBar
	}
)

func NewProgressScanner(scanner Scanner, bar *pb.ProgressBar) Scanner {
	_ = "STUB: not implemented"
	return *new(Scanner)
}

func (ps *progressScanner) Text() string { _ = "STUB: not implemented"; return "" }
