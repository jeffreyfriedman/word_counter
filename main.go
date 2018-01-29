package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var website string

func init() {
	flag.StringVar(&website, "site", "everquote.com", "Website to be scraped")
	flag.Parse()
}

func main() {
	resp, err := http.Get("https://" + website)
	check(err)
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// check(err)
	shouldProcess := false
	z := html.NewTokenizer(resp.Body)
	count := 0
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken && z.Token().Data == "body":
			shouldProcess = true
		case tt == html.EndTagToken && z.Token().Data == "body":
			shouldProcess = false
		case tt == html.StartTagToken && z.Token().Data == "script":
			shouldProcess = false
		case tt == html.EndTagToken && z.Token().Data == "script":
			shouldProcess = true
		case tt == html.TextToken && shouldProcess:
			str := strings.TrimSpace(string(z.Text()))
			if str != "" {
				count += len(strings.Split(str, " "))
				// fmt.Println(str)
			}
		}
		fmt.Println("count: ", count)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
