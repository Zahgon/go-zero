package token

import (
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/placeholder"
)

const (
	Syntax        = "syntax"
	Info          = "info"
	Service       = "service"
	Returns       = "returns"
	Any           = "any"
	TypeKeyword   = "type"
	MapKeyword    = "map"
	ImportKeyword = "import"
)

type Type int

var EofToken = Token{Type: EOF}

var ErrorToken = Token{Type: error}

type Token struct {
	Type     Type
	Text     string
	Position Position
}

func (t Token) Fork(tp Type) Token { _ = "STUB: not implemented"; return *new(Token) }

func (t Token) IsEmptyString() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsComment() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsDocument() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsType(tp Type) bool { _ = "STUB: not implemented"; return false }

func (t Token) Line() int { _ = "STUB: not implemented"; return 0 }

func (t Token) String() string { _ = "STUB: not implemented"; return "" }

func (t Token) Valid() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsKeyword() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsBaseType() bool { _ = "STUB: not implemented"; return false }

func (t Token) IsHttpMethod() bool { _ = "STUB: not implemented"; return false }

func (t Token) Is(text ...string) bool { _ = "STUB: not implemented"; return false }

const (
	token_bg Type = iota
	error
	ILLEGAL
	EOF
	COMMENT
	DOCUMENT

	literal_beg
	IDENT
	INT
	DURATION
	STRING
	RAW_STRING
	PATH
	literal_end

	operator_beg
	SUB
	MUL
	QUO
	ASSIGN

	LPAREN
	LBRACK
	LBRACE
	COMMA
	DOT

	RPAREN
	RBRACE
	RBRACK
	SEMICOLON
	COLON
	ELLIPSIS
	operator_end

	golang_keyword_beg
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
	golang_keyword_end

	api_keyword_bg
	AT_DOC
	AT_HANDLER
	AT_SERVER
	ANY

	api_keyword_end
	token_end
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:      "EOF",
	COMMENT:  "COMMENT",
	DOCUMENT: "DOCUMENT",

	IDENT:      "IDENT",
	INT:        "INT",
	DURATION:   "DURATION",
	STRING:     "STRING",
	RAW_STRING: "RAW_STRING",
	PATH:       "PATH",

	SUB:    "-",
	MUL:    "*",
	QUO:    "/",
	ASSIGN: "=",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	DOT:    ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
	ELLIPSIS:  "...",

	BREAK:    "break",
	CASE:     "case",
	CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT:     "default",
	DEFER:       "defer",
	ELSE:        "else",
	FALLTHROUGH: "fallthrough",
	FOR:         "for",

	FUNC:   "func",
	GO:     "go",
	GOTO:   "goto",
	IF:     "if",
	IMPORT: "import",

	INTERFACE: "interface",
	MAP:       "map",
	PACKAGE:   "package",
	RANGE:     "range",
	RETURN:    "return",

	SELECT: "select",
	STRUCT: "struct",
	SWITCH: "switch",
	TYPE:   "type",
	VAR:    "var",

	AT_DOC:     "@doc",
	AT_HANDLER: "@handler",
	AT_SERVER:  "@server",
	ANY:        "interface{}",
}

var HttpMethods = []interface{}{"get", "head", "post", "put", "patch", "delete", "connect", "options", "trace"}

var httpMethod = map[string]placeholder.Type{
	"get":     placeholder.PlaceHolder,
	"head":    placeholder.PlaceHolder,
	"post":    placeholder.PlaceHolder,
	"put":     placeholder.PlaceHolder,
	"patch":   placeholder.PlaceHolder,
	"delete":  placeholder.PlaceHolder,
	"connect": placeholder.PlaceHolder,
	"options": placeholder.PlaceHolder,
	"trace":   placeholder.PlaceHolder,
}

var keywords = map[string]Type{

	"break":    BREAK,
	"case":     CASE,
	"chan":     CHAN,
	"const":    CONST,
	"continue": CONTINUE,

	"default":     DEFAULT,
	"defer":       DEFER,
	"else":        ELSE,
	"fallthrough": FALLTHROUGH,
	"for":         FOR,

	"func":   FUNC,
	"go":     GO,
	"goto":   GOTO,
	"if":     IF,
	"import": IMPORT,

	"interface": INTERFACE,
	"map":       MAP,
	"package":   PACKAGE,
	"range":     RANGE,
	"return":    RETURN,

	"select": SELECT,
	"struct": STRUCT,
	"switch": SWITCH,
	"type":   TYPE,
	"var":    VAR,
}

var baseDataType = map[string]placeholder.Type{
	"bool":       placeholder.PlaceHolder,
	"uint8":      placeholder.PlaceHolder,
	"uint16":     placeholder.PlaceHolder,
	"uint32":     placeholder.PlaceHolder,
	"uint64":     placeholder.PlaceHolder,
	"int8":       placeholder.PlaceHolder,
	"int16":      placeholder.PlaceHolder,
	"int32":      placeholder.PlaceHolder,
	"int64":      placeholder.PlaceHolder,
	"float32":    placeholder.PlaceHolder,
	"float64":    placeholder.PlaceHolder,
	"complex64":  placeholder.PlaceHolder,
	"complex128": placeholder.PlaceHolder,
	"string":     placeholder.PlaceHolder,
	"int":        placeholder.PlaceHolder,
	"uint":       placeholder.PlaceHolder,
	"uintptr":    placeholder.PlaceHolder,
	"byte":       placeholder.PlaceHolder,
	"rune":       placeholder.PlaceHolder,
	"any":        placeholder.PlaceHolder,
}

func LookupKeyword(ident string) (Type, bool) { _ = "STUB: not implemented"; return *new(Type), false }

func NewIllegalToken(b rune, pos Position) Token { _ = "STUB: not implemented"; return *new(Token) }
