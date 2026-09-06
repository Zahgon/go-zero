package iox

import "io"

func LimitTeeReader(r io.Reader, w io.Writer, n int64) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

type limitTeeReader struct {
	r io.Reader
	w io.Writer
	n int64
}

func (t *limitTeeReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
