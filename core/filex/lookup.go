package filex

import (
	"os"
)

type OffsetRange struct {
	File  string
	Start int64
	Stop  int64
}

func SplitLineChunks(filename string, chunks int) ([]OffsetRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextRange(file *os.File, start, stop int64) (OffsetRange, error) {
	_ = "STUB: not implemented"
	return *new(OffsetRange), nil
}

func skipPartialLine(file *os.File, offset int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
