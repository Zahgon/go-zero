package swagger

func getBoolFromKVOrDefault(properties map[string]string, key string, def bool) bool {
	_ = "STUB: not implemented"
	return false
}

func getFirstUsableString(def ...string) string { _ = "STUB: not implemented"; return "" }

func getListFromInfoOrDefault(properties map[string]string, key string, def []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getOrDefault[T any](properties map[string]string, key string, def T, convert func(string, T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func getStringFromKVOrDefault(properties map[string]string, key string, def string) string {
	_ = "STUB: not implemented"
	return ""
}
