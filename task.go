package waitgroupexer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetURL 计算请求用时
func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
	log.Printf("Completed getting %s in %s", url, time.Since(start))
	return resp, err
}

// CrawlError 自定义错误类型
type CrawlError struct {
	Errors []string
}

func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

// 实现错误的接口
func (c *CrawlError) Error() string {
	return fmt.Sprintf("All Errors: %s", strings.Join(c.Errors, ","))
}

// Present 决定是否应该返回
func (c *CrawlError) Present() bool {
	return len(c.Errors) != 0
}
