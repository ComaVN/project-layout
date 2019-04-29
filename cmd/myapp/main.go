package main

import (
	"fmt"

	"app/myapp"
	"pkg/myprivatelib"
)

func main() {
	fmt.Println("Hello World, from myapp main")
	myapp.Appfunc()
	myprivatelib.Libfunc()
}
