package lib

import (
	"strings"
)

func Truncate(s string, maxLen int) string {
	var sb strings.Builder
	trailChar := "..."
	if len(s) > maxLen {
		sb.WriteString(s[:maxLen])
		sb.WriteString(trailChar)
	} else {
		sb.WriteString(s)
	}
	return sb.String()
}
