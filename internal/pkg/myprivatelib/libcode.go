package myprivatelib

import (
	"fmt"

	"github.com/golang/glog"
)

func Libfunc(message string) {
	glog.Infof("Log from myprivatelib: '%s'\n", message)
	fmt.Printf("Hello world, from myprivatelib: '%s'\n", message)
}
