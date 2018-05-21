package util

import (
	"regexp"
	"strings"
)

func EncodeProjectId(name string) string {
	if len(name) < 1 {
		// return placeholder
		return "placeholder-id"
	}

	r := regexp.MustCompile("[^\\w]")
	name = strings.ToLower(name)
	return r.ReplaceAllString(name, "")
}
