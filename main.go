package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	s "strings"
	"time"
	"web-scraping-go/dg"
)

// CLI init WebScraping
func main() {

	t0 := time.Now()

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	flag.Parse()

	log.Println(*wordPtr)
	log.Println(*numbPtr)
	log.Println(*boolPtr)

	// dg.CollyExample()
	// dg.DR_Crawler()
	// dg.DR_SubCrawler()

	// dg.DR_Crawler_Products()
	// 30 pg, 1435 urls products
	// Iniciou: 21:14:21
	// Terminou: 21:15:21
	// Levou: 1m1.213956847s

	// dg.DRCrawlerUrisProducts()
	// 30 pg, 1435 urls products
	// Iniciou: 21:16:21
	// Terminou: 21:16:21
	// Levou: 1.672431276s

	// 3125 pg, 131302 products url's, 11M archive size
	// Iniciou: 21:20:21
	// Terminou: 21:25:21
	// Levou: 4m29.739605204s

	dg.Scrape_produtos()
	// 132.365 products - 100 nthreads
	// Iniciou: 21:59:21
	// Terminou: 22:22:22
	// Levou: 22m50.503719862s

	// gg()
	// gg2()
	// proxy_utils.CollyExample()
	// proxy_utils.Exx()

	t1 := time.Now()

	fmt.Println("Iniciou:", t0.Format("15:04:15"))
	fmt.Println("Terminou:", t1.Format("15:04:15"))

	fmt.Println("Levou:", t1.Sub(t0))

}

// main ✗ $ go run main.go -word=Davi -numb=7 -fork
// 2024/05/11 12:55:05 Davi
// 2024/05/11 12:55:05 7
// 2024/05/11 12:55:05 true

func gg() {

	f, err := os.ReadFile("texto2.txt")
	if err != nil {
		log.Println("open:", err)
	}
	// defer f.Close()

	// var tex string

	// _, _ = f.Read([]byte(tex))
	// log.Println(string(f))
	str_f := string(f)

	s1 := s.Index(str_f, "url_path")
	s2 := s.Index(str_f[s1:], `",`)

	fmt.Println(s1)
	// url_path":"medicamentos.html
	log.Println(string(f)[s1+11 : s1+s2])

}

func gg2() {

	f, err := os.ReadFile("texto2.txt")
	if err != nil {
		log.Fatalln("open", err)
	}

	var sl []string
	com := 0
	str_f := string(f)

	fmt.Println(len(str_f))
	// fmt.Println("slice:", str_f[70777:70768])

	// for i := 0; i <= 10; i++ {
	for {

		str_f = str_f[com:]
		// log.Println(i)
		s1 := s.Index(str_f, "url_path")
		if s1 < 0 {
			break
		}
		fmt.Println("s1:", s1)
		s2 := s.Index(str_f[s1:], `",`)

		fmt.Println("texto:", str_f[s1+11:s1+s2])
		sl = append(sl, str_f[s1+11:s1+s2])

		com = s1 + s2
		fmt.Println("s1+s2:", s1+s2)

	}

	fmt.Println(sl)
	fmt.Println(len(sl))
}
