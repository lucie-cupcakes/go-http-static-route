package httpStaticRoute

import (
	"net/http"
	"strings"
)

func AddStaticRoutes(staticFiles *map[string]*StaticFile) {
	sf := *staticFiles
	for path := range sf {
		http.HandleFunc("/"+path, func(rw http.ResponseWriter, r *http.Request) {
			methodOk := false
			uri := r.URL.EscapedPath()
			if strings.HasPrefix(uri, "/") {
				uri = strings.TrimPrefix(uri, "/")
				if file, ok := sf[uri]; ok {
					rw.Header().Add("Content-Type", file.ContentType)
					rw.Write(file.Contents)
					methodOk = true
				}
			}
			if !methodOk {
				rw.WriteHeader(500)
			}
		})
	}
}
