package crawlerutils

import (
	"bufio"
	"log"
	"os"
)

func Send_pages(c chan<- string, filepath string) {

	f, err := os.Open(filepath)
	if err != nil {
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
