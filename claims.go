package signcraft

import (
	"encoding/json"
)

type Claims map[string]interface{}

func New() *Claims {
	return &Claims{}
}

func ToClaims(payload interface{}) (Claims, error) {
	structBytes, _ := json.Marshal(payload)
	var claims Claims
	err := json.Unmarshal(structBytes, &claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (c Claims) ToStruct(payload interface{}) error {
	claimsBytes, _ := json.Marshal(c)
	err := json.Unmarshal(claimsBytes, payload)
	if err != nil {
		return err
	}
	return nil
}

// func (c Claims) Validate() error {
// 	now := time.Now().Unix()

// 	if c.Has(NotBeforeAt) {
// 		nbf, _ := c.GetNotBeforeAt()
// 		if now < nbf {
// 			return ErrTokenNotYetValid
// 		}
// 	}

// 	if c.Has(ExpiresAt) {
// 		exp, _ := c.GetExpiresAt()
// 		if now >= exp {
// 			return ErrTokenHasExpired
// 		}
// 	}

// 	return nil
// }
