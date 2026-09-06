package prof

import "github.com/zeromicro/go-zero/core/utils"

type (
	ProfilePoint struct {
		*utils.ElapsedTimer
	}

	Profiler interface {
		Start() ProfilePoint
		Report(name string, point ProfilePoint)
	}

	realProfiler struct{}

	nullProfiler struct{}
)

var profiler = newNullProfiler()

func EnableProfiling() { _ = "STUB: not implemented"; return }

func Start() ProfilePoint { _ = "STUB: not implemented"; return *new(ProfilePoint) }

func Report(name string, point ProfilePoint) { _ = "STUB: not implemented"; return }

func newRealProfiler() Profiler { _ = "STUB: not implemented"; return *new(Profiler) }

func (rp *realProfiler) Start() ProfilePoint { _ = "STUB: not implemented"; return *new(ProfilePoint) }

func (rp *realProfiler) Report(name string, point ProfilePoint) { _ = "STUB: not implemented"; return }

func newNullProfiler() Profiler { _ = "STUB: not implemented"; return *new(Profiler) }

func (np *nullProfiler) Start() ProfilePoint { _ = "STUB: not implemented"; return *new(ProfilePoint) }

func (np *nullProfiler) Report(string, ProfilePoint) { _ = "STUB: not implemented"; return }
