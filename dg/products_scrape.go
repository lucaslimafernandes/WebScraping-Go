package dg

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

// TODO
// Scrapping data from products
// "productData": {
// 	"productBySku": {
// 		"id": 168173,
// 		"sku": "168173",
// 		"name": "Faxxa Rivaroxabana 10mg 30 Comprimidos",
// 		"attribute_set_id": null,
// 		"price": {
// 			"originalPrice": 187.54,
// 			"finalPrice": 175.99,
// 			"tierPrice": 0,
// 			"tierQty": 0,
// 			"type": "DE_POR",
// 			"discountPercentage": 6.1586
// 		},

func Scrape_produtos() {

	channel := make(chan string)
	filepath := "Crawler_products_url.txt"
	go crawlerutils.Send_pages(channel, filepath)
	scraper2(channel)
	log.Println("foi")
}

func scraper(r <-chan string) {

	f, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	nthreads := 10
	var wg sync.WaitGroup
	var records [][]string

	for i := 0; i < nthreads; i++ {
		wg.Add(1)

		go func() {

			for v := range r {

				// fmt.Println("recebido do chan:", v)

				c := colly.NewCollector()

				c.Limit(&colly.LimitRule{
					// DomainGlob:  "*httpbin.*",
					// Parallelism: 2,
					RandomDelay: 5 * time.Second,
					Delay:       800 * time.Millisecond,
				})

				c.OnHTML("head", func(e *colly.HTMLElement) {

					text := e.Text

					// s1 := s.Index(text, `<script type="application/ld+json">`)
					// s2 := s.Index(text[s1:], "</script>")
					// text = text[s1:s2]

					name1 := s.Index(text, "name")
					// name1 := s.Index(text, `"@context":"https://schema.org/","@type":"Product","name":"`)
					log.Println(name1)
					if name1 > 0 {
						fmt.Println(">1")

						// name1 := 110
						name2 := s.Index(text[name1:], `image`) - 5

						if name2 > name1 {
							log.Panicln("PANIC")
						}
						// if name1 > name1+name2 {
						// 	log.Println("name1 > name2", name1, name2, name1+name2)
						// 	fmt.Println(text[name1 : name1+100])
						// 	time.Sleep(time.Second * 3)
						// }
						fmt.Println(text[name1:400])
						// name := text[name1+59 : name1+name2]
						name := text[name1+8 : name1+name2]

						preco1 := s.Index(text, `"price":`)
						preco2 := s.Index(text[preco1:], `",`)
						preco := text[preco1+10 : preco1+preco2]

						sku1 := s.Index(text, "sku")
						sku2 := s.Index(text[sku1:], `",`)
						sku := text[sku1+7 : sku1+sku2]

						records = append(records, []string{name, preco, sku})
						// fmt.Println(name, preco, sku)
					}

				})

				c.OnError(func(r *colly.Response, err error) {
					log.Println("err", err, r.StatusCode)
				})

				c.OnRequest(func(r *colly.Request) {
					log.Println("Visiting:", r.URL)
				})

				c.Visit(v)

			}
			wg.Done()
		}()
	}

	wg.Wait()

	err = csv.NewWriter(f).WriteAll(records)
	if err != nil {
		log.Fatalln(err)
	}

}

func scraper2(r <-chan string) {

	f, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	fe, err := os.Create("products_err.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	nthreads := 100
	var wg sync.WaitGroup
	// var records [][]string
	var urisErr []string

	for i := 0; i < nthreads; i++ {
		wg.Add(1)

		go func() {

			for v := range r {

				var text string
				// fmt.Println("recebido do chan:", v)

				c := colly.NewCollector()

				c.Limit(&colly.LimitRule{
					// DomainGlob:  "*httpbin.*",
					// Parallelism: 2,
					RandomDelay: 5 * time.Second,
					Delay:       800 * time.Millisecond,
				})

				c.OnHTML("head", func(e *colly.HTMLElement) {

					text = e.Text

					// s1 := s.Index(text, `<script type="application/ld+json">`)
					// s2 := s.Index(text[s1:], "</script>")
					// text = text[s1:s2]

				})

				c.OnError(func(r *colly.Response, err error) {
					urisErr = append(urisErr, fmt.Sprintf("%v\t%v\t%v", v, err, r.StatusCode))
					log.Println("err", err, r.StatusCode)
				})

				c.OnRequest(func(r *colly.Request) {
					log.Println("Visiting:", r.URL)
				})

				c.Visit(v)

				name1 := s.Index(text, "name")
				// name1 := s.Index(text, `"@context":"https://schema.org/","@type":"Product","name":"`)
				log.Println(name1)
				if name1 > 0 {
					fmt.Println(">1")

					// name1 := 110
					name2 := s.Index(text[name1:], `image`) - 5

					if name1 > (name1 + name2) {
						// log.Panicln("PANIC")
						urisErr = append(urisErr, v)
						continue
					}
					// if name1 > name1+name2 {
					// 	log.Println("name1 > name2", name1, name2, name1+name2)
					// 	fmt.Println(text[name1 : name1+100])
					// 	time.Sleep(time.Second * 3)
					// }
					fmt.Println(text[name1:400])
					// name := text[name1+59 : name1+name2]
					name := text[name1+8 : name1+name2]

					preco1 := s.Index(text, `"price":`)
					preco2 := s.Index(text[preco1:], `",`)
					preco := text[preco1+10 : preco1+preco2]

					sku1 := s.Index(text, "sku")
					sku2 := s.Index(text[sku1:], `",`)
					sku := text[sku1+7 : sku1+sku2]

					// records = append(records, []string{name, preco, sku})

					err := writer.Write([]string{name, preco, sku})
					if err != nil {
						log.Fatalln("write err", err)
					}
					// fmt.Println(name, preco, sku)
				}
				// continue
				// }

			}
			wg.Done()
		}()
	}

	wg.Wait()

	for _, ue := range urisErr {
		fmt.Fprintln(fe, ue)
	}

	// err = csv.NewWriter(f).WriteAll(records)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
