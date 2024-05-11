package dg

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

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

// http://www.drogaraia.com.br/
func DCrawler() {

	//class="Menustyles__StyleSubLink-sc-1u1vz6z-2 ePRMZz"

	f, err := os.Create("texto2.txt")
	if err != nil {
		log.Fatalln("os.create:", err)
	}
	defer f.Close()

	c := colly.NewCollector()
	// c.OnHTML("a.Menustyles__StyleSubLink-sc-1u1vz6z-2 ePRMZz", func(h *colly.HTMLElement) {
	// c.OnHTML("a.Menustyles__StyleSubLink", func(h *colly.HTMLElement) {
	// 	link := h.Attr("href")
	// 	fmt.Println("Link encontrado:", link)
	// })

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		log.Println("AQUI")
		// log.Println(e.Text)
		_, err = f.WriteString(e.Text)
		if err != nil {
			log.Fatalln("err:", err)
		}
		// if e.Attr("href") == `/medicamentos/tratamento-em-casa/inaladores.html` {
		// log.Println("DEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
		// }
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.Visit("http://www.drogaraia.com.br/")
}
