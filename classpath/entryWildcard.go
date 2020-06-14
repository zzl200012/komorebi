package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 15:42
 */

/* constructor for WildcardEntry, it's actually a composite entry, so we just remove the '*' at the end, then
call the 'Walk' to traverse baseDir and build the zip entries. During the 'WalkFn', children dir will be skipped */
func NewWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := make([]Entry, 0)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := NewZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		panic(err)
	}
	return compositeEntry
}
