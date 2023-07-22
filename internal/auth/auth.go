package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from
// header of an HTTP request
// Authorization : APIKey {api_key of 64 characters}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("inappropriate authorization info found")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
