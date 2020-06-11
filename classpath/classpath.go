package classpath

import (
	"os"
	"path/filepath"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 18:18
 */

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
/* ReadClass seek for classes in boot, ext and user class path in turn */
func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := classpath.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := classpath.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return classpath.userClasspath.readClass(className)
}
/* ToString return the string of the user class path */
func (classpath *Classpath) ToString() string {
	return classpath.userClasspath.ToString()
}
/* parseBootAndExtClasspath parse the boot and extension classes by -Xjre option */
func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	classpath.bootClasspath = NewWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	classpath.extClasspath = NewWildcardEntry(jreExtPath)
}
/* parseUserClasspath parse the user classes by -cp/-classpath option, if no such
 arg, use the current directory as the path for user classes */
func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newEntry(cpOption)
}
/* getJreDir use the input arg -Xjre as jre directory first, if no such arg, search
 the jre directory from the current directory, and finally try the JAVA_HOME env */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Jre Folder Not Found")
}
/* exists return true if the directory exists, or the false if not */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}