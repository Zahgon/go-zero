package mapping

import (
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/lang"
)

const (
	defaultKeyName   = "key"
	delimiter        = '.'
	ignoreKey        = "-"
	numberTypeString = "number"
)

var (
	errTypeMismatch     = errors.New("type mismatch")
	errValueNotSettable = errors.New("value is not settable")
	errValueNotStruct   = errors.New("value type is not struct")
	keyUnmarshaler      = NewUnmarshaler(defaultKeyName)
	durationType        = reflect.TypeOf(time.Duration(0))
	cacheKeys           = make(map[string][]string)
	cacheKeysLock       sync.Mutex
	defaultCache        = make(map[string]any)
	defaultCacheLock    sync.Mutex
	emptyMap            = map[string]any{}
	emptyValue          = reflect.ValueOf(lang.Placeholder)
)

type (
	Unmarshaler struct {
		key  string
		opts unmarshalOptions
	}

	UnmarshalOption func(*unmarshalOptions)

	unmarshalOptions struct {
		fillDefault  bool
		fromArray    bool
		fromString   bool
		opaqueKeys   bool
		canonicalKey func(key string) string
	}
)

func NewUnmarshaler(key string, opts ...UnmarshalOption) *Unmarshaler {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalKey(m map[string]any, v any) error { _ = "STUB: not implemented"; return nil }

func (u *Unmarshaler) Unmarshal(i, v any) error { _ = "STUB: not implemented"; return nil }

func (u *Unmarshaler) UnmarshalValuer(m Valuer, v any) error { _ = "STUB: not implemented"; return nil }

func (u *Unmarshaler) fillMap(fieldType reflect.Type, value reflect.Value,
	mapValue any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillMapFromString(value reflect.Value, mapValue any) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillSlice(fieldType reflect.Type, value reflect.Value,
	mapValue any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillSliceFromString(fieldType reflect.Type, value reflect.Value,
	mapValue any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillSliceValue(slice reflect.Value, index int,
	baseKind reflect.Kind, value any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillSliceWithDefault(derefedType reflect.Type, value reflect.Value,
	defaultValue, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillStructElement(baseType reflect.Type, target reflect.Value,
	value any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) fillUnmarshalerStruct(fieldType reflect.Type,
	value reflect.Value, targetValue string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) generateMap(keyType, elemType reflect.Type, mapValue any,
	fullName string) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (u *Unmarshaler) implementsUnmarshaler(t reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (u *Unmarshaler) parseOptionsWithContext(field reflect.StructField, m Valuer, fullName string) (
	string, *fieldOptionsWithContext, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (u *Unmarshaler) processAnonymousField(field reflect.StructField, value reflect.Value,
	m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processAnonymousFieldOptional(field reflect.StructField, value reflect.Value,
	key string, m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processAnonymousFieldRequired(field reflect.StructField, value reflect.Value,
	m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processAnonymousStructFieldOptional(fieldType reflect.Type,
	value reflect.Value, key string, m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processField(field reflect.StructField, value reflect.Value,
	m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processFieldNotFromString(fieldType reflect.Type, value reflect.Value,
	vp valueWithParent, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processFieldPrimitive(fieldType reflect.Type, value reflect.Value,
	mapValue any, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processFieldPrimitiveWithJSONNumber(fieldType reflect.Type, value reflect.Value,
	v json.Number, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processFieldStruct(fieldType reflect.Type, value reflect.Value,
	m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processFieldTextUnmarshaler(fieldType reflect.Type, value reflect.Value,
	mapValue any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (u *Unmarshaler) processFieldWithEnvValue(fieldType reflect.Type, value reflect.Value,
	envVal string, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processNamedField(field reflect.StructField, value reflect.Value,
	m valuerWithParent, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processNamedFieldWithValue(fieldType reflect.Type, value reflect.Value,
	vp valueWithParent, key string, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processNamedFieldWithValueFromString(fieldType reflect.Type, value reflect.Value,
	mapValue any, key string, opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) processNamedFieldWithoutValue(fieldType reflect.Type, value reflect.Value,
	opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) unmarshal(i, v any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) unmarshalValuer(m Valuer, v any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Unmarshaler) unmarshalWithFullName(m valuerWithParent, v any, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func WithStringValues() UnmarshalOption { _ = "STUB: not implemented"; return *new(UnmarshalOption) }

func WithCanonicalKeyFunc(f func(string) string) UnmarshalOption {
	_ = "STUB: not implemented"
	return *new(UnmarshalOption)
}

func WithDefault() UnmarshalOption { _ = "STUB: not implemented"; return *new(UnmarshalOption) }

func WithFromArray() UnmarshalOption { _ = "STUB: not implemented"; return *new(UnmarshalOption) }

func WithOpaqueKeys() UnmarshalOption { _ = "STUB: not implemented"; return *new(UnmarshalOption) }

func createValuer(v valuerWithParent, opts *fieldOptionsWithContext) valuerWithParent {
	_ = "STUB: not implemented"
	return *new(valuerWithParent)
}

func fillDurationValue(fieldType reflect.Type, value reflect.Value, dur string) error {
	_ = "STUB: not implemented"
	return nil
}

func fillPrimitive(fieldType reflect.Type, value reflect.Value, mapValue any,
	opts *fieldOptionsWithContext, fullName string) error {
	_ = "STUB: not implemented"
	return nil
}

func fillWithSameType(fieldType reflect.Type, value reflect.Value, mapValue any,
	opts *fieldOptionsWithContext) error {
	_ = "STUB: not implemented"
	return nil
}

func getValue(m valuerWithParent, key string, opaque bool) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func getValueWithChainedKeys(m valuerWithParent, keys []string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func join(elem ...string) string { _ = "STUB: not implemented"; return "" }

func newInitError(name string) error { _ = "STUB: not implemented"; return nil }

func newTypeMismatchError(name string) error { _ = "STUB: not implemented"; return nil }

func newTypeMismatchErrorWithHint(name, expectType, actualType string) error {
	_ = "STUB: not implemented"
	return nil
}

func readKeys(key string, opaque bool) []string { _ = "STUB: not implemented"; return nil }

func setSameKindValue(targetType reflect.Type, target reflect.Value, value any) {
	_ = "STUB: not implemented"
	return
}
