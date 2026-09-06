//go:build linux || darwin || freebsd

package pathx

func ReadLink(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }
