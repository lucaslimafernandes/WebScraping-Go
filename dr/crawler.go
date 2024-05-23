package dr

import (
	"fmt"
	"log"
	"os"
	sl "slices"
	"strconv"
	s "strings"
	"sync"
	"time"
	crawlerutils "web-scraping-go/crawler_utils"

	"github.com/gocolly/colly/v2"
)

// Principal function to Crawler DrogaRaia
func DRCrawler() {

	chPag := make(chan string)
	chPrd := make(chan string)

	crawlerCategories()
	time.Sleep(time.Second)

	go crawlerutils.Send_pages(chPag, "data_scrap/Crawler_links.txt")
	categoriesPagination(chPag)
	time.Sleep(time.Second)

	go crawlerutils.Send_pages(chPrd, "data_scrap/Crawler_all_pages.txt")
	crawlerProductsUris(chPrd)
	time.Sleep(time.Second)

}

// DrogaRaia homepage categories/subcategories crawler urls
func crawlerCategories() {

	f, err := os.Create("data_scrap/Crawler_links.txt")
	if err != nil {
		log.Fatalln("os.create:", err)
	}
	defer f.Close()

	c := colly.NewCollector()
	// c.SetRequestTimeout(time.Second * 10)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {

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
			// Pular página de manipulados
			if v != "medicamentos/manipulados.html" {
				_, err := fmt.Fprintln(f, v)
				if err != nil {
					log.Fatalln("write", err)
				}
			}
		}

	})

	c.Visit("http://www.drogaraia.com.br/")

}

func categoriesPagination(channel <-chan string) {

	URL := "http://www.drogaraia.com.br/"
	nth := 5

	var wg sync.WaitGroup

	f_all_pages, err := os.Create("data_scrap/Crawler_all_pages.txt")
	if err != nil {
		log.Fatalln("DR_SubCrawler Create all pg:", err)
	}
	defer f_all_pages.Close()

	for t := 0; t < nth; t++ {
		wg.Add(1)

		go func() {
			for text := range channel {
				paginas := total_pages(fmt.Sprintf("http://www.drogaraia.com.br/%v", text), "script#__NEXT_DATA__")

				for i := 1; i <= paginas; i++ {
					_, err := f_all_pages.WriteString(fmt.Sprintf("%v%v?page=%v\n", URL, text, i))
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func crawlerProductsUris(channel <-chan string) {

	f, err := os.Create("data_scrap/Crawler_products_url.txt")
	if err != nil {
		log.Fatalln("os create txt", err)
	}
	defer f.Close()

	fe, err := os.Create("data_scrap/crawlerPrds_err.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fe.Close()

	var wg sync.WaitGroup
	nth := 5

	for t := 0; t < nth; t++ {
		wg.Add(1)
		go func() {

			for ch_url := range channel {

				var uris []string

				c := colly.NewCollector()

				c.Limit(&colly.LimitRule{
					RandomDelay: 5 * time.Second,
					Delay:       800 * time.Millisecond,
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
								uris = append(uris, uri)
							}
						}
						com = s1 + s2
					}

				})

				c.OnError(func(r *colly.Response, err error) {
					fe.WriteString(fmt.Sprintf("%v\t%v\n", err, r.StatusCode))
				})

				c.Visit(ch_url)

				for _, v := range uris {
					fmt.Fprintln(f, v)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

// Capture total pages from subcategories page
func total_pages(url, qs string) int {

	var paginacao string

	fe, err := os.Create("data_scrap/categoriesPag_err.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fe.Close()

	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
		Delay:       800 * time.Millisecond,
	})

	c.OnHTML(qs, func(e *colly.HTMLElement) {

		text := e.Text
		pages := s.Index(text, "total_pages")
		endpages := s.Index(text[pages:], ",")

		paginacao = text[pages+13 : pages+endpages]

	})

	c.Visit(url)
	total_pages, err := strconv.Atoi(paginacao)
	if err != nil {
		log.Println(err)
		return 1
	}

	c.OnError(func(r *colly.Response, err error) {
		fe.WriteString(fmt.Sprintf("%v\t%v\n", err, r.StatusCode))
	})

	return total_pages

}
