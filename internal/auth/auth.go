package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the HTTP request headers
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(Headers http.Header) (string, error) {
	val := Headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authenication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
