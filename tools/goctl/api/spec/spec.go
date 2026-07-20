package spec

const RoutePrefixKey = "prefix"

type (
	Doc []string

	Annotation struct {
		Properties map[string]string
	}

	ApiSyntax struct {
		Version string
		Doc     Doc
		Comment Doc
	}

	ApiSpec struct {
		Info    Info
		Syntax  ApiSyntax
		Imports []Import
		Types   []Type
		Service Service
	}

	Import struct {
		Value   string
		Doc     Doc
		Comment Doc
	}

	Group struct {
		Annotation Annotation
		Routes     []Route
	}

	Info struct {
		Title string

		Desc string

		Version string

		Author string

		Email      string
		Properties map[string]string
	}

	Member struct {
		Name string

		Type    Type
		Tag     string
		Comment string

		Docs     Doc
		IsInline bool
	}

	Route struct {
		AtServerAnnotation Annotation
		Method             string
		Path               string
		RequestType        Type
		ResponseType       Type
		Docs               Doc
		Handler            string
		AtDoc              AtDoc
		HandlerDoc         Doc
		HandlerComment     Doc
		Doc                Doc
		Comment            Doc
	}

	Service struct {
		Name   string
		Groups []Group
	}

	Type interface {
		Name() string
		Comments() []string
		Documents() []string
	}

	DefineStruct struct {
		RawName string
		Members []Member
		Docs    Doc
	}

	NestedStruct struct {
		RawName string
		Members []Member
		Docs    Doc
	}

	PrimitiveType struct {
		RawName string
	}

	MapType struct {
		RawName string

		Key string

		Value Type
	}

	ArrayType struct {
		RawName string
		Value   Type
	}

	InterfaceType struct {
		RawName string
	}

	PointerType struct {
		RawName string
		Type    Type
	}

	AtDoc struct {
		Properties map[string]string
		Text       string
	}
)
