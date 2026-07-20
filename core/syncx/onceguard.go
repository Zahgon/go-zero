package syncx

type OnceGuard struct {
	done uint32
}

func (og *OnceGuard) Taken() bool { _ = "STUB: not implemented"; return false }

func (og *OnceGuard) Take() bool { _ = "STUB: not implemented"; return false }
