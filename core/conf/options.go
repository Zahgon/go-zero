package conf

type (
	Option func(opt *options)

	options struct {
		env bool
	}
)

func UseEnv() Option { _ = "STUB: not implemented"; return *new(Option) }
