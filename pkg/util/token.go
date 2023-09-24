package util

import "strings"

func BuildBearerToken(token string) string {
	var builder strings.Builder
	builder.WriteString("Bearer")
	builder.WriteString(token)

	return builder.String()
}
