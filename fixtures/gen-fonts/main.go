// generates the fixtures/fonts.html for testing the fonts in docker.
// Use the google translate to translate "test" into all the languages, print the result into a html page.
// By reviewing the generated pdf we can find out what font is missing for a specific language.

// Package main ...
package main

import (
	"fmt"
	"log"

	"github.com/xyjwsj/grod"
	"github.com/xyjwsj/grod/lib/launcher"
	"github.com/xyjwsj/grod/lib/utils"
)

func main() {
	url := launcher.New().MustLaunch()
	b := rod.New().ControlURL(url).MustConnect()
	defer b.MustClose()

	p := b.MustPage("https://translate.google.com/")

	p.MustElement("#source").MustInput("Test the google translate.")

	if p.MustHas(".tlid-dismiss-button") {
		p.MustElement(".tlid-dismiss-button").MustClick()
	}

	showList := p.MustElement(".tlid-open-target-language-list")
	list := p.MustElements(".language-list:nth-child(2) .language_list_section:nth-child(2) .language_list_item_language_name")

	html := ""

	for _, lang := range list {
		showList.MustClick()
		wait := p.MustWaitRequestIdle()
		lang.MustClick()
		wait()
		name := lang.MustText()
		result := p.MustElement(".tlid-translation").MustText()
		log.Println(name, result)
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", name, result)
	}

	html = fmt.Sprintf(`<html>
		<p style="font-family: serif;">
			This file is generated by <code>"fixtures/gen-fonts"</code>
		</p>
		<p>Test smileys: 😀 😁 😂 🤣 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘 🥰 😗 😙 😚</p>
		<table>
		%s
		</table></html>`,
		html,
	)

	utils.E(utils.OutputFile("fixtures/fonts.html", html))
}
