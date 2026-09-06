//go:build linux || darwin || freebsd

package proc

import (
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	defaultWrapUpTime = time.Second

	defaultWaitTime = 5500 * time.Millisecond
)

var (
	wrapUpListeners   = new(listenerManager)
	shutdownListeners = new(listenerManager)
	wrapUpTime        = defaultWrapUpTime
	waitTime          = defaultWaitTime
	shutdownLock      sync.Mutex
)

type ShutdownConf struct {
	WrapUpTime time.Duration `json:",default=1s"`

	WaitTime time.Duration `json:",default=5.5s"`
}

func AddShutdownListener(fn func()) (waitForCalled func()) { _ = "STUB: not implemented"; return nil }

func AddWrapUpListener(fn func()) (waitForCalled func()) { _ = "STUB: not implemented"; return nil }

func SetTimeToForceQuit(duration time.Duration) { _ = "STUB: not implemented"; return }

func Setup(conf ShutdownConf) { _ = "STUB: not implemented"; return }

func Shutdown() { _ = "STUB: not implemented"; return }

func WrapUp() { _ = "STUB: not implemented"; return }

func gracefulStop(signals chan os.Signal, sig syscall.Signal) { _ = "STUB: not implemented"; return }

type listenerManager struct {
	lock      sync.Mutex
	waitGroup sync.WaitGroup
	listeners []func()
}

func (lm *listenerManager) addListener(fn func()) (waitForCalled func()) {
	_ = "STUB: not implemented"
	return nil
}

func (lm *listenerManager) notifyListeners() { _ = "STUB: not implemented"; return }
