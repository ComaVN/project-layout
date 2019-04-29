package main

import (
	"fmt"

	"project-layout/internal/app/myapp"
	"project-layout/internal/pkg/myprivatelib"
)

func main() {
	fmt.Println("Hello World, from myapp main")
	myapp.Appfunc()
	myprivatelib.Libfunc("myapp main")
}
