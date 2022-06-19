// path.Split()

package main

import (
	"fmt"
	"path"
)

func main() {
	// example 1
	// var dir, file string

	// dir, file = path.Split("css/main.css")

	// fmt.Println("dir: ", dir)
	// fmt.Println("file: ", file)

	// example 2
	// var file string
	// **use _
	// _, file = path.Split("css/main.css")

	// fmt.Println("file: ", file)

	// example 3
	// short declaration
	_, file := path.Split("css/main.css")

	fmt.Println("file: ", file)

}
