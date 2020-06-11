package classpath

import (
	"errors"
	"strings"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 15:07
 */

type CompositeEntry []Entry

/* constructor for CompositeEntry, split the path list and add all entries to the composite entry returned */
func NewCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := make([]Entry, 0)
	for _, v := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(v)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
/* readClass traverse all the entries for the class, once found, return the value */
func (compositeEntry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range compositeEntry {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("Class Not Found: " + className)
}
/* ToString return the union of the absolute path of every entries joined with pathListSeparator */
func (compositeEntry CompositeEntry) ToString() string {
	strs := make([]string, len(compositeEntry))
	for i, entry := range compositeEntry {
		strs[i] = entry.ToString()
	}
	return strings.Join(strs, pathListSeparator)
}