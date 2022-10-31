package futil

import "os"

var WD = ""

func init() {
	var err error
	WD, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
