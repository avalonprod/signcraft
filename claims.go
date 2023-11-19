package signcraft

type Claims map[string]interface{}

func New() *Claims {
	return &Claims{}
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
