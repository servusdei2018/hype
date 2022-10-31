package futil

import (
	"os"
	"path"
)

// Exists returns whether a file exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirMustExist creates a directory if it doesn't already exist.
func DirMustExist(fpath string, name string) error {
	if !Exists(path.Join(fpath, name)) {
		return os.Mkdir(path.Join(fpath, name), 0754)
	}
	return nil
}
