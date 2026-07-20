//go:build linux || darwin || freebsd

package proc

const DefaultMemProfileRate = 4096

var started uint32

type Profile struct {
	closers []func()

	stopped uint32
}

func (p *Profile) close() { _ = "STUB: not implemented"; return }

func (p *Profile) startBlockProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) startCpuProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) startMemProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) startMutexProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) startThreadCreateProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) startTraceProfile() { _ = "STUB: not implemented"; return }

func (p *Profile) Stop() { _ = "STUB: not implemented"; return }

func StartProfile() Stopper { _ = "STUB: not implemented"; return *new(Stopper) }

func createDumpFile(kind string) string { _ = "STUB: not implemented"; return "" }
