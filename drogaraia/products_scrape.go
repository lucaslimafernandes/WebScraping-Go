package drogaraia

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	s "strings"
	"sync"
	"time"
	crawlerutils "web-scraping-go/crawler_utils"

	"github.com/gocolly/colly/v2"
)

// Scrape products data
func Scrape_produtcs(nthreads int) {

	channel := make(chan string)
	filepath := "data_scrap/Crawler_products_url.txt"
	go crawlerutils.Send_pages(channel, filepath)
	scraper(channel, nthreads)

}

// go scraper
func scraper(r <-chan string, nthreads int) {

	f, err := os.Create("data_scrap/products.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	err = writer.Write([]string{"Nome", "Preço", "SKU"})
	if err != nil {
		log.Fatalln("write err", err)
	}

	fe, err := os.Create("data_scrap/products_err.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fe.Close()

	var wg sync.WaitGroup

	for i := 0; i < nthreads; i++ {
		wg.Add(1)

		go func() {

			for v := range r {

				var text string

				c := colly.NewCollector()
				c.SetRequestTimeout(time.Second * 10)

				c.Limit(&colly.LimitRule{
					RandomDelay: 5 * time.Second,
					Delay:       800 * time.Millisecond,
				})

				c.OnHTML("head", func(e *colly.HTMLElement) {
					text = e.Text
				})

				c.OnError(func(r *colly.Response, err error) {
					_, err = fe.WriteString(fmt.Sprintf("%v\t%v\t%v", v, err, r.StatusCode))
					if err != nil {
						log.Fatalln(err)
					}
				})

				c.Visit(v)

				name1 := s.Index(text, "name")
				if name1 > 0 {

					name2 := s.Index(text[name1:], `image`) - 5

					if name1 > (name1 + name2) {
						_, err = fe.WriteString(fmt.Sprintf("%v\tif name1 > (name1 + name2)", v))
						if err != nil {
							log.Fatalln(err)
						}
						continue
					}

					name := text[name1+8 : name1+name2]

					preco1 := s.Index(text, `"price":`)
					preco2 := s.Index(text[preco1:], `",`)
					preco := text[preco1+10 : preco1+preco2]

					sku1 := s.Index(text, "sku")
					sku2 := s.Index(text[sku1:], `",`)
					sku := text[sku1+7 : sku1+sku2]

					err := writer.Write([]string{name, preco, sku})
					if err != nil {
						log.Fatalln("write err", err)
					}
				}

			}
			wg.Done()
		}()
	}

	wg.Wait()

}
