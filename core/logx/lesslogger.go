package logx

type LessLogger struct {
	*limitedExecutor
}

func NewLessLogger(milliseconds int) *LessLogger { _ = "STUB: not implemented"; return nil }

func (logger *LessLogger) Error(v ...any) { _ = "STUB: not implemented"; return }

func (logger *LessLogger) Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }
