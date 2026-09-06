package proc

var noopStopper nilStopper

type (
	Stopper interface {
		Stop()
	}

	nilStopper struct{}
)

func (ns nilStopper) Stop() { _ = "STUB: not implemented"; return }
