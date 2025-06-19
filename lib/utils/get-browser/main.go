// Package main ...
package main

import (
	"fmt"

	"github.com/xyjwsj/grod/lib/launcher"
	"github.com/xyjwsj/grod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
