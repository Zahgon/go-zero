package token

import (
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const claimHistoryResetDuration = time.Hour * 24

type (
	ParseOption func(parser *TokenParser)

	TokenParser struct {
		resetTime     time.Duration
		resetDuration time.Duration
		history       sync.Map
	}
)

func NewTokenParser(opts ...ParseOption) *TokenParser { _ = "STUB: not implemented"; return nil }

func (tp *TokenParser) ParseToken(r *http.Request, secret, prevSecret string) (*jwt.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tp *TokenParser) doParseToken(r *http.Request, secret string) (*jwt.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tp *TokenParser) incrementCount(secret string) { _ = "STUB: not implemented"; return }

func (tp *TokenParser) loadCount(secret string) uint64 { _ = "STUB: not implemented"; return 0 }

func WithResetDuration(duration time.Duration) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func newParser() *jwt.Parser { _ = "STUB: not implemented"; return nil }
