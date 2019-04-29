package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"

	"project-layout/internal/app/myapp"
	"project-layout/internal/pkg/myprivatelib"
)

func main() {
	flag.Parse()
	glog.Infoln("Log from myapp main")
	fmt.Println("Hello World, from myapp main")
	myapp.Appfunc()
	myprivatelib.Libfunc("myapp main")
}
