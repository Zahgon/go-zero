package logx

type redirector struct{}

func CollectSysLog() { _ = "STUB: not implemented"; return }

func (r *redirector) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
