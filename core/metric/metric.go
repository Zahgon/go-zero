package metric

type VectorOpts struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}

func update(fn func()) { _ = "STUB: not implemented"; return }
