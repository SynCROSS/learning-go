package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://gall.dcinside.com/board/lists/?id=baseball_new9"

func main() {
  fmt.Println(getPages())
}

func getPages() int {
  numberPages:=0

	res, err := http.Get(baseURL)

	checkError(err)
	checkStatus(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	checkError(err)

	doc.Find(".bottom_paging_box").Each(func(i int, s *goquery.Selection) {
		numberPages=s.Find("a").Length() - 1
	})

	return numberPages
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkStatus(res *http.Response) {
	if res.StatusCode >= 400 {
		log.Fatalln("Requst is Failed Status:", res.StatusCode)
	}
}
