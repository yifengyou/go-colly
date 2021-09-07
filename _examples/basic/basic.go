package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector 创建采集器
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		// 根据url规则匹配，白名单，只允许指定域名
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		// 有允许就有禁止，黑名单
		//colly.DisallowedDomains()
	)

	// On every a element which has href attribute call callback
	// 回调函数，当有html数据时候，过滤HTML元素
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	// 访问url之前
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})


	// Start scraping on https://hackerspaces.org
	// 执行爬虫入口，必须在注册回调函数之后执行
	c.Visit("https://hackerspaces.org/")

	// Visit 会阻塞，直到结束
	fmt.Println("All done!")
}
