package filex

import (
	"errors"
	"os"
)

var errExceedFileSize = errors.New("exceed file size")

type RangeReader struct {
	file  *os.File
	start int64
	stop  int64
}

func NewRangeReader(file *os.File, start, stop int64) *RangeReader {
	_ = "STUB: not implemented"
	return nil
}

func (rr *RangeReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
