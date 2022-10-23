package main

import (
    "fmt"

    "github.com/gocolly/colly"
)
// Falabella website - https://www.bancofalabella.com.co

// Will get all the information of the day of the transactions
// of the company Falabella

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

    // Find and print all links
    c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
        fmt.Println(links)
    })
    c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}

