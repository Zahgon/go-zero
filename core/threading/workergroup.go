package threading

type WorkerGroup struct {
	job     func()
	workers int
}

func NewWorkerGroup(job func(), workers int) WorkerGroup {
	_ = "STUB: not implemented"
	return *new(WorkerGroup)
}

func (wg WorkerGroup) Start() { _ = "STUB: not implemented"; return }
