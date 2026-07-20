package mapping

const notSymbol = '!'

type (
	fieldOptionsWithContext struct {
		Inherit    bool
		FromString bool
		Optional   bool
		Options    []string
		Default    string
		EnvVar     string
		Range      *numberRange
	}

	fieldOptions struct {
		fieldOptionsWithContext
		OptionalDep string
	}

	numberRange struct {
		left         float64
		leftInclude  bool
		right        float64
		rightInclude bool
	}
)

func (o *fieldOptionsWithContext) fromString() bool { _ = "STUB: not implemented"; return false }

func (o *fieldOptionsWithContext) getDefault() (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (o *fieldOptionsWithContext) inherit() bool { _ = "STUB: not implemented"; return false }

func (o *fieldOptionsWithContext) optional() bool { _ = "STUB: not implemented"; return false }

func (o *fieldOptionsWithContext) options() []string { _ = "STUB: not implemented"; return nil }

func (o *fieldOptions) optionalDep() string { _ = "STUB: not implemented"; return "" }

func (o *fieldOptions) toOptionsWithContext(key string, m Valuer, fullName string) (
	*fieldOptionsWithContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
