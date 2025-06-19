// Package main ...
package main

import (
	"fmt"

	"github.com/xyjwsj/grod"
)

func main() {
	rod.New().MustConnect().MustPage("https://www.google.com/").MustWaitLoad().MustPDF("sample.pdf")
	fmt.Println("wrote sample.pdf")
}
