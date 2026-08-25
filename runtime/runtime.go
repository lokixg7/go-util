package runtime

import (
	"runtime"
	"strings"
)

// GetCurCalleeFunc returns the name of its direct caller.
func GetCurCalleeFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	pathList := strings.Split(name, ".")
	return pathList[len(pathList)-1]
}

// GetParentCallFunc returns the name of its caller's parent.
func GetParentCallFunc() string {
	pc, _, _, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	pathList := strings.Split(name, ".")
	return pathList[len(pathList)-1]
}
