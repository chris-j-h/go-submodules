package main

import (
	"github.com/chris-j-h/go-submodules/subpkg/subsubpkg"
	"fmt"
)

const Name = subsubpkg.Name

func main() {
	fmt.Println(Name)
}
