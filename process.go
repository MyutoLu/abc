package waitgroupexer

import (
	"log"
	"sync"
	"time"
)

// Crawl 从url清单切片中收集响应， 返回前等待所有的请求完成
func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Printf("staring crawling")
	wg := &sync.WaitGroup{}

	var resps []int
	cerr := &CrawlError{}
	for _, v := range sites {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			resp, err := GetURL(v)
			if err != nil {
				cerr.Add(err)
				return
			}
			resps = append(resps, resp.StatusCode)
		}(v)
	}
	wg.Wait()
	if cerr.Present() {
		return resps, cerr
	}

	log.Printf("completed crawling in %s.", time.Since(start))
	return resps, nil
}
