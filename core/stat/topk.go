package stat

type taskHeap []Task

func (h *taskHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *taskHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (h *taskHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h *taskHeap) Push(x any) { _ = "STUB: not implemented"; return }

func (h *taskHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }

func topK(all []Task, k int) []Task { _ = "STUB: not implemented"; return nil }
