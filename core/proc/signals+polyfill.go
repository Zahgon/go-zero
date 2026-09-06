//go:build windows

package proc

func Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }
