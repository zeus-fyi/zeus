package cookbooks

import (
	"os"
	"path"
	"runtime"
)

func ChangeToCookbookDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}
