package constants

import (
	"path/filepath"
	"runtime"
)

var _, base, _, ok = runtime.Caller(0)
var Basepath = filepath.Join(filepath.Dir(base), "../")
