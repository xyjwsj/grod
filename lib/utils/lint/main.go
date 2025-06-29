// Package main ...
package main

import (
	"github.com/xyjwsj/grod/lib/utils"
)

func main() {
	utils.UseNode(true)

	utils.Exec("npx -ys -- cspell@6.31.1 --no-progress **")

	utils.Exec("npx -ys -- eslint@8.41.0 --ext=.js,.html --fix --ignore-path=.gitignore .")

	utils.Exec("npx -ys -- prettier@2.8.8 --loglevel=error --write --ignore-path=.gitignore .")

	utils.Exec("go run github.com/ysmood/golangci-lint@latest")

	lintMustPrefix()

	checkGitClean()
}

func checkGitClean() {
	out := utils.ExecLine(false, "git status --porcelain")
	if out != "" {
		panic("Please run \"go generate\" on local and git commit the changes:\n" + out)
	}
}
