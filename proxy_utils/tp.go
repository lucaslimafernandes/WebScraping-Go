package proxy_utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func CollyExample() {

	c := colly.NewCollector()

	// Rotate two socks5 proxies
	// http://pubproxy.com/api/proxy
	// https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all
	// rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")

	rp, err := proxy.RoundRobinProxySwitcher("http://185.217.199.114:4444", "http://185.232.169.108:4444")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnHTML("body", func(h *colly.HTMLElement) {
		// h.Request.Visit(h.Attr("href"))
		log.Println(h.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println(string(r.Body))
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
		log.Println(r.Headers)
	})

	c.Visit("http://httpbin.org/ip")
	// c.Visit("https://httpbin.org/ip")

}

func Prs() []string {

	var slice []string

	f, err := os.Open("ps.txt")
	if err != nil {
		log.Println("ps open err:", err)
		return slice
	}
	defer f.Close()

	fb := bufio.NewScanner(f)
	fb.Split(bufio.ScanLines)

	for fb.Scan() {
		slice = append(slice, fmt.Sprintf("http://%s", fb.Text()))
	}

	return slice

}

func Exx() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	prxs := Prs()
	// rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	rp, err := proxy.RoundRobinProxySwitcher(prxs...)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)
	c.SetRequestTimeout(time.Second * 2)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("err:", string(r.Body))
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < len(prxs); i++ {
		c.Visit("https://httpbin.org/ip")
	}
}
