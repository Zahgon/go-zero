package timex

import "time"

var initTime = time.Now().AddDate(-1, -1, -1)

func Now() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func Since(d time.Duration) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
