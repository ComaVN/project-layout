package myapp

import (
	"fmt"
	"project-layout/internal/pkg/myprivatelib"
)

func Appfunc() {
	fmt.Println("Hello world, from myapp internals")
	myprivatelib.Libfunc("myapp internal")
}
