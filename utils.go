package runsass

import (
	"path"
	"strings"
)

type argString []string

func (a argString) Get(i int, args ...string) (r string) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

func SanitizePath(fileOrDirPath string) string {
	return path.Clean(strings.Replace(fileOrDirPath, "\\", "/", -1))
}
