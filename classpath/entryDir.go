package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 14:42
 */

type DirEntry struct {
	absolutePath string
}

/* readClass first join the absolute path with the classname, then open the certain file */
func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(dirEntry.absolutePath + className)
	data, err := ioutil.ReadFile(fileName)
	return data, dirEntry, err
}
/* ToString return the absolute path of the Entry */
func (dirEntry *DirEntry) ToString() string {
	return dirEntry.absolutePath
}
/* constructor for DirEntry */
func NewDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absolutePath: absPath}
}