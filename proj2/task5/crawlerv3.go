package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// The url to get the main part of the web
var destUrl = "http://tumregels.github.io/Network-Programming-with-Go/"

// The url to get search_index.json used for search
var jasonUrl = "http://tumregels.github.io/Network-Programming-with-Go/search_index.json"

// The urls to get the icons
var iconUrl1 = "http://tumregels.github.io/Network-Programming-with-Go/gitbook/fonts/fontawesome/fontawesome-webfont.woff"
var iconUrl2 = "http://tumregels.github.io/Network-Programming-with-Go/gitbook/fonts/fontawesome/fontawesome-webfont.woff2"

// The url used for test
// var testUrl = "http://tumregels.github.io/Network-Programming-with-Go/gitbook/fonts/fontawesome/fontawesome-webfont.woff2"

func errCheck(err error) {
	if err != nil {
		panic(err)
	} // if
} // errCheck

// helper function to save the file given path and response
func saveFile(savedPath string, res *http.Response) {
	f, err := os.Create(savedPath)
	errCheck(err)
	io.Copy(f, res.Body)
} // saveFile

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: tumregels.github.io
		colly.AllowedDomains("tumregels.github.io"),
		colly.MaxDepth(1),
	)
	os.Mkdir("./Network-Programming-with-Go/", 0777)
	os.Mkdir("./Network-Programming-with-Go/assets", 0777)
	os.Mkdir("./Network-Programming-with-Go/gitbook", 0777)

	c.OnResponse(func(r *colly.Response) {
		// // 以下代码将打印得到的response body的全部内容
		// fmt.Println("body", string(r.Body))
		// 获取链接的绝对路径
		fullurl := r.Request.URL.String()
		// fmt.Println(fullurl)
		// fmt.Println("Hello this is OnResponse")
		res, _ := http.Get(fullurl)
		//解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(fullurl)
		errCheck(err)
		h := strings.Split(u.Path, "/")
		if h[len(h)-3] == "fonts" {
			// do noting
			// fmt.Println(h[len(h) - 2])
			// fmt.Printf("Response has tye: %T\n", res)
			dirPath1 := "./Network-Programming-with-Go/gitbook/" + h[len(h)-3]
			os.Mkdir(dirPath1, 0777)
			dirPath2 := dirPath1 + "/" + h[len(h)-2]
			os.Mkdir(dirPath2, 0777)
			savedPath := dirPath2 + "/" + h[len(h)-1]
			saveFile(savedPath, res)
		} else if (h[len(h)-2]) == "Network-Programming-with-Go" {
			savedPath := "./Network-Programming-with-Go/" + "index.html"
			saveFile(savedPath, res)
		} else if (h[len(h)-1]) == "" {
			dirPath := "./Network-Programming-with-Go/" + h[len(h)-2]
			os.Mkdir(dirPath, 0777)
			// savedPath := "./Network-Programming-with-Go/" + h[len(h)-2] + ".html"
			// savedPath := "./Network-Programming-with-Go/" + h[len(h)-2] + "/" + "index.html"
			savedPath := dirPath + "/index.html"
			saveFile(savedPath, res)
		} else {
			savedPath := "./Network-Programming-with-Go/" + h[len(h)-2] + "/" + h[len(h)-1]
			saveFile(savedPath, res)
		}
	})

	// search for all link tags that have a rel attribute that is equal to stylesheet - CSS
	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		// hyperlink reference
		link := e.Attr("href")
		// 获取css的绝对路径
		fullurl := e.Request.AbsoluteURL(link)
		// fmt.Println("CSS path is: " + fullurl)
		res, _ := http.Get(fullurl)
		//解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(fullurl)
		errCheck(err)
		h := strings.Split(u.Path, "/")
		// fmt.Println(u)
		// fmt.Println(h[1])
		if h[len(h)-2] != "gitbook" {
			dirPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-2]
			os.Mkdir(dirPath, 0777)
			// savedPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-2] + "/" + h[len(h)-1]
			savedPath := dirPath + "/" + h[len(h)-1]
			saveFile(savedPath, res)
		} else {
			savedPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-1]
			saveFile(savedPath, res)
		}
	})

	// search for all script tags with src attribute -- JS
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		// src attribute
		link := e.Attr("src")
		// 获取js的绝对路径
		fullurl := e.Request.AbsoluteURL(link)
		// fmt.Println("JavaScript path is: " + fullurl)
		res, _ := http.Get(fullurl)
		//解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(fullurl)
		errCheck(err)
		h := strings.Split(u.Path, "/")
		// fmt.Println(u)
		// fmt.Println(h[1])
		if h[len(h)-2] != "gitbook" {
			dirPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-2]
			os.Mkdir(dirPath, 0777)
			savedPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-2] + "/" + h[len(h)-1]
			saveFile(savedPath, res)
		} else {
			savedPath := "./Network-Programming-with-Go/gitbook/" + h[len(h)-1]
			saveFile(savedPath, res)
		}
	})

	// serach for all img tags with src attribute -- Images
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		srcRef := e.Attr("src")
		// 获取图片的绝对路径
		fullurl := e.Request.AbsoluteURL(srcRef)
		// fmt.Println("Image path is: " + fullurl)
		res, _ := http.Get(fullurl)
		//解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(fullurl)
		errCheck(err)
		h := strings.Split(u.Path, "/")
		// fmt.Println(u)
		// fmt.Println(h[1])
		if h[1] == "Network-Programming-with-Go" {
			savedPath := "./Network-Programming-with-Go/assets/" + h[len(h)-1]
			saveFile(savedPath, res)
		} // if
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// fmt.Println(link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		urlPath := r.URL.String()
		fmt.Println("Visiting", urlPath)
		//解析这个 URL 并确保解析没有出错。
		u, err := url.Parse(urlPath)
		errCheck(err)
		// split the url according to "/"
		h := strings.Split(u.Path, "/")
		// if the url is the path for search_index.json, save it to the file
		if h[len(h)-1] == "search_index.json" {
			res, _ := http.Get(urlPath)
			savedPath := "./Network-Programming-with-Go/" + h[len(h)-1]
			saveFile(savedPath, res)
		} // if
	})

	// Start scraping on the given url
	// c.Visit(testUrl)
	c.Visit(iconUrl1)
	c.Visit(iconUrl2)
	c.Visit(jasonUrl)
	c.Visit(destUrl)
} // main
