package main

import (
	"fmt"
	"myuto.net/waitgroupexer"
)

func main() {
	sites := []string{
		"https://www.baidu.com/",
		//"https://www.iqiyi.com/",
		//"https://v.qq.com/",
		"https://www.jd.com/",
		//"https://www.tmall.com/",
	}

	resps, err := waitgroupexer.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resps received:", resps)
}
