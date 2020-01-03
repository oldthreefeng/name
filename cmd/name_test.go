/*
Copyright 2020 louis.
@Time : 2020/1/2 21:06
@Author : louis
@File : name_test.go
@Software: GoLand
*/
package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"testing"
)

/*
	html := `<body>

				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`
*/
func TestMeimingteng(t *testing.T) {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	dom.Find("div[lang=zh]+p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}