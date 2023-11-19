package signcraft

import "errors"

var (
	ErrInvalidToken      = errors.New("Token is invalid")
	ErrNotFound          = errors.New("Claim key not found in claims")
	ErrClaimValueInvalid = errors.New("Claim value invalid")
	ErrTokenHasExpired   = errors.New("Token has expired")
)
