package spec

const (
	bodyTagKey        = "json"
	formTagKey        = "form"
	pathTagKey        = "path"
	headerTagKey      = "header"
	defaultSummaryKey = "summary"
)

var definedKeys = []string{bodyTagKey, formTagKey, pathTagKey, headerTagKey}

func (s Service) JoinPrefix() Service { _ = "STUB: not implemented"; return *new(Service) }

func (s Service) Routes() []Route { _ = "STUB: not implemented"; return nil }

func (m Member) Tags() []*Tag { _ = "STUB: not implemented"; return nil }

func (m Member) IsOptional() bool { _ = "STUB: not implemented"; return false }

func (m Member) IsOmitEmpty() bool { _ = "STUB: not implemented"; return false }

func (m Member) GetPropertyName() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (m Member) GetComment() string { _ = "STUB: not implemented"; return "" }

func (m Member) IsBodyMember() bool { _ = "STUB: not implemented"; return false }

func (m Member) IsFormMember() bool { _ = "STUB: not implemented"; return false }

func (m Member) IsTagMember(tagKey string) bool { _ = "STUB: not implemented"; return false }

func (m Member) GetEnumOptions() []string { _ = "STUB: not implemented"; return nil }

func (t DefineStruct) GetBodyMembers() []Member { _ = "STUB: not implemented"; return nil }

func (t DefineStruct) GetFormMembers() []Member { _ = "STUB: not implemented"; return nil }

func (t DefineStruct) GetNonBodyMembers() []Member { _ = "STUB: not implemented"; return nil }

func (t DefineStruct) GetTagMembers(tagKey string) []Member { _ = "STUB: not implemented"; return nil }

func (r Route) JoinedDoc() string { _ = "STUB: not implemented"; return "" }

func (r Route) GetAnnotation(key string) string { _ = "STUB: not implemented"; return "" }

func (g Group) GetAnnotation(key string) string { _ = "STUB: not implemented"; return "" }

func (r Route) ResponseTypeName() string { _ = "STUB: not implemented"; return "" }

func (r Route) RequestTypeName() string { _ = "STUB: not implemented"; return "" }
