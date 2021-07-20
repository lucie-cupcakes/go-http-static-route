package httpStaticRoute

import (
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func detectMimeType(path string, data *[]byte) string {
	path = strings.ToLower(path)
	// Fast path based on file name
	if strings.HasSuffix(path, ".html") {
		return "text/html"
	}
	if strings.HasSuffix(path, ".js") {
		return "text/javascript"
	}
	if strings.HasSuffix(path, ".css") {
		return "text/css"
	}
	if strings.HasSuffix(path, ".json") {
		return "application/json"
	}
	if strings.HasSuffix(path, ".xml") {
		return "application/xml"
	}
	if strings.HasSuffix(path, ".webp") {
		return "image/webp"
	}
	if strings.HasSuffix(path, ".jpg") {
		return "image/jpeg"
	}
	if strings.HasSuffix(path, ".png") {
		return "image/png"
	}
	if strings.HasSuffix(path, ".svg") {
		return "image/svg+xml"
	}
	if strings.HasPrefix(path, ".ico") {
		return "image/vnd.microsoft.icon"
	}

	// Other files need binary detection
	return mimetype.Detect(*data).String()
}
