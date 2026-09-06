package logx

type Sensitive interface {
	MaskSensitive() any
}

func maskSensitive(v any) any { _ = "STUB: not implemented"; return *new(any) }
