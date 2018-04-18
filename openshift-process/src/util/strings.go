package util

import (
	"regexp"
	"strings"
)

func EncodeProjectId(name string) string {
	r := regexp.MustCompile("[^\\w]")
	name = strings.ToLower(name)
	return r.ReplaceAllString(name, "")
}
