package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnResponse(func(r *colly.Response) {
		// 以下代码将打印得到的response body的全部内容
		// fmt.Println("body", string(r.Body))
		// 解析response body
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error message: %v\n", err)
			os.Exit(1)
		} // if
		// 用htmlquery.Find寻找class="tb-rmb-num"的em
		list := htmlquery.Find(doc, "//em[@class=\"tb-rmb-num\"]")
		// 使用htmlquery.InnerText()获取内容
		// htmlquery.Innertext()输入一个*html.Node，返回一个string
		fmt.Println("The price is: " + htmlquery.InnerText(list[0]))
		// 以下代码将html.Node转化为string
		var b bytes.Buffer
		errR := html.Render(&b, list[0])
		if errR != nil {
			fmt.Fprintf(os.Stderr, "Error message: %v\n", errR)
			os.Exit(1)
		} // if
		fmt.Println(b.String())
	})

	c.Visit("https://item.taobao.com/item.htm?id=622329071715")
} // main
