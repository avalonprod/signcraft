package signcraft

import (
	"time"
)

const (
	ID        = "jti"
	NotBefore = "nbf"
	Issuer    = "iss"
	Subject   = "sub"
	IssuedAt  = "iat"
	OriginID  = "origin_id"
	Expiry    = "exp"
)

type Claims map[string]interface{}

type StandartClaims struct {
	ID        string `json:"jti,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Expiry    int64  `json:"exp,omitempty"`
	OriginID  string `json:"origin_id,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func New() *Claims {
	return &Claims{}
}

func NewWithClaims(standartClaims StandartClaims) *Claims {
	claims := Claims{}
	if standartClaims.ID != "" {
		claims[ID] = standartClaims.ID
	}
	if standartClaims.Expiry != 0 {
		claims[Expiry] = standartClaims.Expiry
	}
	if standartClaims.NotBefore != 0 {
		claims[NotBefore] = standartClaims.NotBefore
	}
	if standartClaims.OriginID != "" {
		claims[OriginID] = standartClaims.OriginID
	}
	if standartClaims.Issuer != "" {
		claims[Issuer] = standartClaims.Issuer
	}
	if standartClaims.IssuedAt != 0 {
		claims[IssuedAt] = standartClaims.IssuedAt
	}
	if standartClaims.Subject != "" {
		claims[Subject] = standartClaims.Subject
	}
	return &claims
}

func (c Claims) Set(name string, value interface{}) { c[name] = value }
func (c Claims) Del(name string)                    { delete(c, name) }
func (c Claims) Has(name string) bool               { _, ok := c[name]; return ok }
func (c Claims) Get(name string) (interface{}, error) {
	if !c.Has(name) {
		return nil, ErrNotFound
	}
	return c[name], nil
}

func (c Claims) SetTokenID(tokenID string)            { c[ID] = tokenID }
func (c Claims) SetIssuer(issuer string)              { c[Issuer] = issuer }
func (c Claims) SetSubject(subject string)            { c[Subject] = subject }
func (c Claims) SetIssuedAt(issuedAt time.Time)       { c[IssuedAt] = issuedAt.Unix() }
func (c Claims) SetExpiresAt(expiry time.Time)        { c[Expiry] = expiry.Unix() }
func (c Claims) SetNotBeforeAt(notbeforeAt time.Time) { c[NotBefore] = notbeforeAt.Unix() }
func (c Claims) SetOriginID(originId string)          { c[OriginID] = originId }

func (c Claims) GetExpiresAt() (int64, error) {
	if !c.Has(Expiry) {
		return 0, ErrNotFound
	}

	expiry, err := c.GetInt(Expiry)
	if err != nil {
		return 0, ErrClaimValueInvalid
	}

	return int64(expiry), nil
}

func (c Claims) Validate() error {
	now := time.Now().Unix()

	if exp, _ := c.GetExpiresAt(); exp == 0 {
		return nil
	}

	if c.Has(Expiry) {
		exp, _ := c.GetExpiresAt()
		if now >= exp {
			return ErrTokenHasExpired
		}
	}

	return nil
}
