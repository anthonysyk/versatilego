package fs

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		rootPath   = fmt.Sprintf("%v/", filepath.Dir(filepath.Dir(b)))
	)

	return rootPath
}

func GetResourcesFile(filename string) string {
	var (
		_, b, _, _    = runtime.Caller(0)
		resourcesPath = fmt.Sprintf("%v/resources/", filepath.Dir(filepath.Dir(b)))
	)

	return fmt.Sprintf("%v/%v", resourcesPath, filename)
}
