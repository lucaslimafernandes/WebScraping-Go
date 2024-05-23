package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"web-scraping-go/drogaraia"

	"github.com/briandowns/spinner"
)

// Inicializando
func init() {

	os.Mkdir("data_scrap", 0777)

}

// CLI init WebScraping
func main() {

	s := spinner.New(spinner.CharSets[36], time.Millisecond*100)
	s.Start()

	t0 := time.Now()

	crw_links := flag.Bool("crawler", false, "a bool")
	scr_data := flag.Bool("scrape", false, "a bool")
	nthreads := flag.Int("nthreads", 100, "an int")

	flag.Parse()

	if *crw_links {
		drogaraia.DRCrawler()
	}

	if *scr_data {
		drogaraia.Scrape_produtcs(*nthreads)
	}

	if !*crw_links && !*scr_data {
		fmt.Printf("\n\tUsage go run main.go [OPTIONS]\n\tor ./main [OPTIONS]\n")
		fmt.Printf("\n\tExecute first crawler and then scrape!\n\n")
		fmt.Printf("-crawler\t\tRun the crawler to discover all product URLs!\n")
		fmt.Printf("-scrape\t\t\tRun the scraper to collect all products data!\n")
		fmt.Printf("-nthreads 100\t\tSet threads number, optional, default 100\n\n")
	} else {

		s.Stop()

		t1 := time.Now()

		fmt.Println("Iniciou:", t0.Format("15:04:15"))
		fmt.Println("Terminou:", t1.Format("15:04:15"))

		fmt.Println("Levou:", t1.Sub(t0))
	}

}
