package main

import (
	"flag"
	"log"
	"web-scraping-go/dg"
)

// CLI init WebScraping
func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	flag.Parse()

	log.Println(*wordPtr)
	log.Println(*numbPtr)
	log.Println(*boolPtr)

	// dg.CollyExample()
	dg.DCrawler()

}

// main ✗ $ go run main.go -word=Davi -numb=7 -fork
// 2024/05/11 12:55:05 Davi
// 2024/05/11 12:55:05 7
// 2024/05/11 12:55:05 true
