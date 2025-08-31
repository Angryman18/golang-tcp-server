package server

import (
	"strings"
)

func ParsePathName(httpContext []byte) string {

	ctx := string(httpContext)

	firstIdx := strings.Index(ctx, "/")
	lastIdx := strings.Index(ctx, "HTTP")

	if firstIdx != -1 && lastIdx != -1 {
		return ctx[firstIdx:lastIdx]
	}
	return "/"
}
