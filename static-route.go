package httpStaticRoute

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	lucieutils "gitlab.com/lucie_cupcakes/lucie-utils"
)

type StaticFile struct {
	Contents []byte
	Path     string
	MimeType *mimetype.MIME
}

func LoadStaticFiles(dirPath string,
	filterFilePath func(filePath string) bool) (*map[string]*StaticFile, error) {
	staticFiles := make(map[string]*StaticFile)
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filterFilePath(path) {
				fileBytes, err := lucieutils.OpenAndReadAllFile(path)
				if err != nil {
					return err
				}
				if strings.Contains(path, "/") {
					var f StaticFile
					f.Path = path
					f.Contents = fileBytes
					f.MimeType = mimetype.Detect(fileBytes)
					path = strings.SplitN(path, "/", 2)[1]
					staticFiles[path] = &f
				}
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("Error walking staticRoute Filesystem: %s", err.Error())
	}
	return &staticFiles, nil
}

func AddStaticRoutes(staticFiles *map[string]*StaticFile) {
	sf := *staticFiles
	for path := range sf {
		http.HandleFunc("/"+path, func(rw http.ResponseWriter, r *http.Request) {
			methodOk := false
			uri := r.URL.EscapedPath()
			if strings.HasPrefix(uri, "/") {
				uri = strings.TrimPrefix(uri, "/")
				if file, ok := sf[uri]; ok {
					rw.Header().Add("Content-Type", file.MimeType.String())
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
