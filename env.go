package main

var rootPath string

func SetRootPath(path string) {
	rootPath = path
}
func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	} else {
		return "."
	}
}
