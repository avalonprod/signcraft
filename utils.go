package signcraft

import (
	"fmt"
	"strconv"
)

func BytesToString(b []byte) string {
	return string(b)
}

func (c Claims) GetStr(name string) (string, error) {
	if !c.Has(name) {
		return "", ErrNotFound
	}

	switch val := c[name].(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	}

	return fmt.Sprintf("%v", c[name]), nil
}

func (c Claims) GetInt(name string) (int, error) {
	if !c.Has(name) {
		return 0, ErrNotFound
	}

	switch val := c[name].(type) {
	case string:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, ErrClaimValueInvalid
		}
		return int(v), nil
	case float32:
		return int(val), nil
	case float64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case int:
		return int(val), nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	}

	return 0, ErrClaimValueInvalid
}

func (c Claims) GetFloat(name string) (float64, error) {
	if !c.Has(name) {
		return 0, ErrNotFound
	}

	switch val := c[name].(type) {
	case float32:
		return float64(val), nil
	case float64:
		return float64(val), nil
	case string:
		v, _ := strconv.ParseFloat(val, 64)
		return v, nil
	}

	return 0, ErrClaimValueInvalid
}

func (c Claims) GetBool(name string) (bool, error) {
	if !c.Has(name) {
		return false, ErrNotFound
	}

	switch val := c[name].(type) {
	case string:
		v, _ := strconv.ParseBool(val)
		return v, nil
	case bool:
		return val, nil
	}

	return false, ErrClaimValueInvalid
}
