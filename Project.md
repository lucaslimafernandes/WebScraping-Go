
https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/

https://pkg.go.dev/github.com/gocolly/colly

Level up even further!
Scrape more data:

Generating paginated URLs based on the input URL, such as example.com/blog/1, example.com/blog/2, etc
Following links
Add CLI arguments and flags to allow the user to customize the application:

Specify the URLs to scrape
Use custom CSS selectors
Enable auto-pagination based on an input pattern
Enable link following using a CSS selector
Implement rate limiting per domain so the user doesn't get blocked if they make lots of requests

Cache pages so subsequent runs don't need to download the same page again:

Download a fresh copy if the cached page is over 1 day old
Add a CLI flag to clear the cache
Add a CLI flag to ignore the cache and download the page anyway
Save the new page in the cache