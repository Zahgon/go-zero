package logx

import (
	"io"
	"os"
)

var fileSys realFileSystem

type (
	fileSystem interface {
		Close(closer io.Closer) error
		Copy(writer io.Writer, reader io.Reader) (int64, error)
		Create(name string) (*os.File, error)
		Open(name string) (*os.File, error)
		Remove(name string) error
	}

	realFileSystem struct{}
)

func (fs realFileSystem) Close(closer io.Closer) error { _ = "STUB: not implemented"; return nil }

func (fs realFileSystem) Copy(writer io.Writer, reader io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (fs realFileSystem) Create(name string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs realFileSystem) Open(name string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs realFileSystem) Remove(name string) error { _ = "STUB: not implemented"; return nil }
