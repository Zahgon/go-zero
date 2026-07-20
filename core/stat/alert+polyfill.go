//go:build !linux

package stat

func Report(string) { _ = "STUB: not implemented"; return }

func SetReporter(func(string)) { _ = "STUB: not implemented"; return }
