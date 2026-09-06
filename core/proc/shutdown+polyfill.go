//go:build windows

package proc

import "time"

type ShutdownConf struct{}

func AddShutdownListener(fn func()) func() { _ = "STUB: not implemented"; return nil }

func AddWrapUpListener(fn func()) func() { _ = "STUB: not implemented"; return nil }

func SetTimeToForceQuit(duration time.Duration) { _ = "STUB: not implemented"; return }

func Setup(conf ShutdownConf) { _ = "STUB: not implemented"; return }

func Shutdown() { _ = "STUB: not implemented"; return }

func WrapUp() { _ = "STUB: not implemented"; return }
