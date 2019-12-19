package util

import (
	"strings"
)

func SubstrBefore(str, sep string) string {
	return substr(str, true)
}

func SubstrAfter(str, sep string) string {
	return substr(str, false)
}

func substr(str string, upToStart bool) string {
	index := strings.Index(str, ":")
	if index > 0 {
		if upToStart {
			str = str[1:index]
		} else {
			str = str[index+1 : len(str)]
		}
	}
	return str
}
