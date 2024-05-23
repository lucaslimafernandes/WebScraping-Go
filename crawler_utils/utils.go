package crawlerutils

import (
	"bufio"
	"log"
	"os"
)

// Func to read file and
// sending URL's to chan
func Send_pages(c chan<- string, filepath string) {

	f, err := os.Open(filepath)
	if err != nil {
		if filepath == "data_scrap/Crawler_products_url.txt" {
			log.Println("Please, execute -crawler first!")
		}
		log.Fatalln(err)
	}
	defer f.Close()

	fb := bufio.NewScanner(f)
	fb.Split(bufio.ScanLines)

	for fb.Scan() {
		if fb.Text() != "" {
			c <- fb.Text()
		}
	}
	close(c)

}
