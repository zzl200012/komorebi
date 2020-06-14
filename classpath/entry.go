package classpath

import (
	"os"
	"strings"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 14:37
 */
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	ToString() string
}

/* constructor for Entry */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return NewCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return NewWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return NewZipEntry(path)
	}
	return NewDirEntry(path)
}
