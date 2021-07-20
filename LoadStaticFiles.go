package httpStaticRoute

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lucieutils "gitlab.com/lucie_cupcakes/lucie-utils"
)

type StaticFile struct {
	Contents    []byte
	Path        string
	ContentType string
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
					f.ContentType = detectMimeType(path, &fileBytes)
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
