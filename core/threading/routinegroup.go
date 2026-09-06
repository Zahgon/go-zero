package threading

import "sync"

type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup { _ = "STUB: not implemented"; return nil }

func (g *RoutineGroup) Run(fn func()) { _ = "STUB: not implemented"; return }

func (g *RoutineGroup) RunSafe(fn func()) { _ = "STUB: not implemented"; return }

func (g *RoutineGroup) Wait() { _ = "STUB: not implemented"; return }
