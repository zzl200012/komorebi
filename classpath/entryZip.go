package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 14:55
 */

type ZipEntry struct {
	absolutePath string
}

/* readClass first open the ZIP file, then in a for loop to find the class file, and finally get the data */
func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(zipEntry.absolutePath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, f := range reader.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, zipEntry, nil
		}
	}
	return nil, nil, errors.New("Class Not Found: " + className)
}

/* ToString return the absolute path of the Entry */
func (zipEntry *ZipEntry) ToString() string {
	return zipEntry.absolutePath
}

/* constructor for ZipEntry */
func NewZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absolutePath: absPath}
}
