package httpx

import "github.com/zeromicro/go-zero/rest/internal/header"

const (
	ContentEncoding = "Content-Encoding"

	ContentSecurity = "X-Content-Security"

	ContentType = header.ContentType

	JsonContentType = header.ContentTypeJson

	KeyField = "key"

	SecretField = "secret"

	TypeField = "type"

	CryptionType = 1
)

const (
	CodeSignaturePass = iota

	CodeSignatureInvalidHeader

	CodeSignatureWrongTime

	CodeSignatureInvalidToken
)
