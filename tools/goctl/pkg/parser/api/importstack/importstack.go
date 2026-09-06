package importstack

import "errors"

var ErrImportCycleNotAllowed = errors.New("import cycle not allowed")

type ImportStack []string

func New() *ImportStack { _ = "STUB: not implemented"; return nil }

func (s *ImportStack) Push(p string) error { _ = "STUB: not implemented"; return nil }

func (s *ImportStack) Pop() { _ = "STUB: not implemented"; return }

func (s *ImportStack) List() []string { _ = "STUB: not implemented"; return nil }
