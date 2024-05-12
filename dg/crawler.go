package dg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
	"time"
	"web-scraping-go/proxy_utils"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

// https://go-colly.org/docs/examples/rate_limit/

func CollyExample() {

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
	})

	c.Visit("http://go-colly.org/")

}

// with proxies http://www.drogaraia.com.br/
func DR_Crawler() {

	f, err := os.Create("Crawler_links.txt")
	if err != nil {
		log.Fatalln("os.create:", err)
	}
	defer f.Close()

	c := colly.NewCollector()
	c.SetRequestTimeout(time.Second * 10)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		log.Println("AQUI")
		// fmt.Println(e.Text)
		var sl []string
		text := e.Text
		com := 0

		for {

			text = text[com:]
			s1 := s.Index(text, "url_path")
			if s1 <= 0 {
				break
			}
			s2 := s.Index(text[s1:], `",`)
			sl = append(sl, text[s1+11:s1+s2])

			com = s1 + s2

		}

		for _, v := range sl {
			if v != "medicamentos/manipulados.html" {
				_, err := fmt.Fprintln(f, v)
				if err != nil {
					log.Fatalln("write", err)
				}
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		// fmt.Println("ON ERR:", string(r.Body))
		fmt.Println("ON ERR:", r.StatusCode)
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("OnScraped", string(r.Body))
	// })

	c.Visit("http://www.drogaraia.com.br/")

}

func DR_SubCrawler() {

	URL := "http://www.drogaraia.com.br/"
	var sl []string

	f, err := os.Open("Crawler_links.txt")
	if err != nil {
		log.Fatalln("DR_SubCrawler Read:", err)
	}
	defer f.Close()

	f_all_pages, err := os.Create("Crawler_all_pages.txt")
	if err != nil {
		log.Fatalln("DR_SubCrawler Create all pg:", err)
	}
	defer f_all_pages.Close()

	fb := bufio.NewScanner(f)
	fb.Split(bufio.ScanLines)

	for fb.Scan() {
		// Paginação
		// collyCrawler(sl[0], ".Found__FoundStyles-sc-62hzma-0 bjMYbQ")
		paginas := collyCrawler(fmt.Sprintf("http://www.drogaraia.com.br/%s", fb.Text()), "script#__NEXT_DATA__")

		for i := 1; i <= paginas; i++ {
			sl = append(sl, fmt.Sprintf("%s%s?page=%v", URL, fb.Text(), i))
		}
	}

	for _, v := range sl {
		_, err = fmt.Fprintln(f_all_pages, v)
		if err != nil {
			log.Println(err)
		}
	}

}

func collyCrawler(url, qs string) int {

	f, err := os.Create("collyC.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var paginacao string

	fmt.Println("collyCrawler", url, qs)
	c := colly.NewCollector()

	c.OnHTML(qs, func(e *colly.HTMLElement) {
		// h.Request.Visit(h.Attr(qsr))

		text := e.Text
		pages := s.Index(text, "total_pages")
		endpages := s.Index(text[pages:], ",")

		paginacao = text[pages+13 : pages+endpages]
		fmt.Println("Paginação:", paginacao)

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
	})

	c.Visit(url)
	paginacao_int, err := strconv.Atoi(paginacao)
	if err != nil {
		log.Println(err)
	}

	return paginacao_int

}

// with proxies http://www.drogaraia.com.br/
func DR_Crawler_proxies() {

	f, err := os.Create("Crawler_links.txt")
	if err != nil {
		log.Fatalln("os.create:", err)
	}
	defer f.Close()

	c := colly.NewCollector(colly.AllowURLRevisit())

	prxies := proxy_utils.Prs()

	// rp, err := proxy.RoundRobinProxySwitcher("http://185.217.199.114:4444", "http://185.232.169.108:4444")
	rp, err := proxy.RoundRobinProxySwitcher(prxies...)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)
	c.SetRequestTimeout(time.Second * 10)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		// log.Println("AQUI")
		// fmt.Println(e.Text)
		var sl []string
		text := e.Text
		com := 0

		for {

			text = text[com:]
			s1 := s.Index(text, "url_path")
			if s1 <= 0 {
				break
			}
			s2 := s.Index(text[s1:], `",`)
			sl = append(sl, text[s1+11:s1+s2])

			com = s1 + s2

		}

		for _, v := range sl {
			_, err := fmt.Fprintln(f, v)
			if err != nil {
				log.Fatalln("write", err)
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		// fmt.Println("ON ERR:", string(r.Body))
		fmt.Println("ON ERR:", r.StatusCode)
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("OnScraped", string(r.Body))
	// })

	for i := 0; i < len(prxies); i++ {
		c.Visit("http://www.drogaraia.com.br/")
	}

}
