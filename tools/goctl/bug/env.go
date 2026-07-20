package bug

type env map[string]string

func (e env) string() string { _ = "STUB: not implemented"; return "" }

func getEnv() env { _ = "STUB: not implemented"; return *new(env) }
