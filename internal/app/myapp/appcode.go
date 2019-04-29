package myapp

import (
	"fmt"
	"pkg/myprivatelib"
)

func Appfunc() {
	fmt.Println("Hello world, from myapp internals")
	myprivatelib.Libfunc("myapp internal")
}
