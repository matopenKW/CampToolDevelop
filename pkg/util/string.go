package util

import (
	"strings"
)

func SubstrBefore(str, sep string) string {
	return substr(str, sep, true)
}

func SubstrAfter(str, sep string) string {
	return substr(str, sep, false)
}

func substr(str, sep string, upToStart bool) string {
	index := strings.Index(str, sep)
	if index > -1 {
		if upToStart {
			str = str[1:index]
		} else {
			str = str[index+1 : len(str)]
		}
	}
	return str
}
