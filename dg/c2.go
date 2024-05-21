package dg

import (
	"fmt"
	"log"
	"os"
	sl "slices"
	s "strings"
	"sync"
	"time"
	crawlerutils "web-scraping-go/crawler_utils"

	"github.com/gocolly/colly/v2"
)

func DRCrawlerUrisProducts() {

	channel := make(chan string)
	filepath := "Crawler_all_pages.txt"
	go crawlerutils.Send_pages(channel, filepath)

	urisCrawler(channel)

}

func urisCrawler(r <-chan string) {

	nthreads := 20
	var wg sync.WaitGroup

	f, err := os.Create("Crawler_products_url.txt")
	if err != nil {
		log.Fatalln("os create txt", err)
	}
	defer f.Close()

	for i := 0; i < nthreads; i++ {
		wg.Add(1)

		go func() {

			for v := range r {

				fmt.Println("Recebido do chan:", v)

				var uris []string

				c := colly.NewCollector()

				c.Limit(&colly.LimitRule{
					// DomainGlob:  "*httpbin.*",
					// Parallelism: 2,
					RandomDelay: 5 * time.Second,
				})

				c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {

					text := e.Text
					com := 0

					for {

						text = text[com:]
						s1 := s.Index(text, "url_key")
						if s1 <= 0 {
							break
						}
						s2 := s.Index(text[s1:], "],") - 1

						uri := text[s1+26 : s1+s2]
						if uri != `ildren":` && uri != `,"children":` {
							uri = fmt.Sprintf("http://www.drogaraia.com.br/%v.html", uri)

							if len(uri) < 200 && !sl.Contains(uris, uri) {
								fmt.Println("Paginação:", uri)
								uris = append(uris, uri)
							}
						}
						com = s1 + s2
					}

				})

				c.OnRequest(func(r *colly.Request) {
					log.Println("Visiting:", r.URL)
				})

				c.Visit(v)

				for _, v := range uris {
					fmt.Fprintln(f, v)
				}
			}

			wg.Done()
		}()

	}
	wg.Wait()

}
