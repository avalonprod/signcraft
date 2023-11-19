package signcraft

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// Generate token using claims
func (c *Claims) Sign(secret []byte) (string, error) {
	header, err := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256"})
	if err != nil {
		return "", err
	}

	claims, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	token := fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString(header),
		base64.RawStdEncoding.EncodeToString(claims),
	)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(token))
	return fmt.Sprintf("%s.%s", token, base64.RawStdEncoding.EncodeToString(mac.Sum(nil))), nil
}

// Generate token
func Sign(secret []byte, payload interface{}) (string, error) {
	header, err := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256"})
	if err != nil {
		return "", err
	}

	claims, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	token := fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString(header),
		base64.RawStdEncoding.EncodeToString(claims),
	)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(token))
	return fmt.Sprintf("%s.%s", token, base64.RawStdEncoding.EncodeToString(mac.Sum(nil))), nil
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

func GenerateUUID() (string, error) {
	version := byte(4)
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | (version << 4)

	uuid[8] = (uuid[8] & 0xbf) | 0x80

	buf := make([]byte, 36)
	var dash byte = '-'
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], uuid[10:])

	return string(buf), nil
}
