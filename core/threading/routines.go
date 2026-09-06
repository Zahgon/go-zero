package threading

import (
	"context"
)

func GoSafe(fn func()) { _ = "STUB: not implemented"; return }

func GoSafeCtx(ctx context.Context, fn func()) { _ = "STUB: not implemented"; return }

func RoutineId() uint64 { _ = "STUB: not implemented"; return 0 }

func RunSafe(fn func()) { _ = "STUB: not implemented"; return }

func RunSafeCtx(ctx context.Context, fn func()) { _ = "STUB: not implemented"; return }
