package myapp

import (
	"fmt"

	"github.com/golang/glog"

	"project-layout/internal/pkg/myprivatelib"
)

func Appfunc() {
	glog.Infoln("Log from myapp internals")
	fmt.Println("Hello world, from myapp internals")
	myprivatelib.Libfunc("myapp internal")
}
