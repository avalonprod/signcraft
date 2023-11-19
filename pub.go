package signcraft

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func (c Claims) Set(name string, value interface{}) { c[name] = value }
func (c Claims) Del(name string)                    { delete(c, name) }
func (c Claims) Has(name string) bool               { _, ok := c[name]; return ok }
func (c Claims) Get(name string) (interface{}, error) {
	if !c.Has(name) {
		return nil, ErrNotFound
	}
	return c[name], nil
}

// Generate token
func (c *Claims) Generate(secret []byte) string {
	header, _ := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256"})
	claims, _ := json.Marshal(c)
	token := fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString(header),
		base64.RawStdEncoding.EncodeToString(claims),
	)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(token))
	return fmt.Sprintf("%s.%s", token, base64.RawStdEncoding.EncodeToString(mac.Sum(nil)))
}

// Parsing token
func Parse(token string) (Claims, error) {
	tokenArray := strings.Split(token, ".")

	if len(tokenArray) != 3 {
		return nil, ErrInvalidToken
	}
	claimsByte, err := base64.RawURLEncoding.DecodeString(tokenArray[1])
	if err != nil {
		return nil, err
	}

	var claims Claims
	err = json.Unmarshal(claimsByte, &claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// Verify token
func Verify(token string, secret []byte) bool {
	tokenArray := strings.Split(token, ".")
	if len(tokenArray) != 3 {
		return false
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(fmt.Sprintf("%s.%s", tokenArray[0], tokenArray[1])))
	sig, _ := base64.RawURLEncoding.DecodeString(tokenArray[2])
	return hmac.Equal(sig, mac.Sum(nil))
}
