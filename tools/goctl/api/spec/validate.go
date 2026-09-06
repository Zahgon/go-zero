package spec

import "errors"

var ErrMissingService = errors.New("missing service")

func (s *ApiSpec) Validate() error { _ = "STUB: not implemented"; return nil }
