package tsgen

import (
	_ "embed"
)

//go:embed request.ts
var requestTemplate string

func genRequest(dir string) error { _ = "STUB: not implemented"; return nil }
