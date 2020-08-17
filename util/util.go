package util

import (
	"net/url"
)

// Test if URI is valid.
func IsValidURI(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}