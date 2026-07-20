package ast

import "errors"

var ErrImportCycleNotAllowed = errors.New("import cycle not allowed")

type importStack []string

func (s *importStack) push(p string) error { _ = "STUB: not implemented"; return nil }

func (s *importStack) pop() { _ = "STUB: not implemented"; return }
