package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://gall.dcinside.com/board/lists/?id=baseball_new9"

type p struct {
	id          string
	postType    string
	title       string
	replies     string
	writer      string
	dateWriting string
	views       string
	recommend   string
}

func main() {
	totalPages := getTotalPages()

	// fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		getPage(i + 1)
	}
}

func getPage(pageNum int) {
	pageURL := baseURL + "&page=" + strconv.Itoa(pageNum)

	fmt.Println("Requesting: " + pageURL)

	res, err := http.Get(pageURL)

	checkError(err)
	checkStatus(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)

	defer res.Body.Close()
	checkError(err)

	posts := doc.Find(".ub-content")

	posts.Each(func(i int, post *goquery.Selection) {
		id := post.Find(".gall_num").Text()
		titleClass, _ := post.Find(".gall_tit > a > em").Attr("class")
		titleType := []rune(titleClass)
		postType := string(titleType[14:17])
		title := post.Find(".gall_tit > a ").Text()

		replies := post.Find(".gall_tit > .reply_numbox > .reply_num").Text()

		if replies == "" {
			replies = "[0]"
		} else {
			b := []byte(title)
			title = ""
			for i := 0; i < len(b); i += 3 {
				r, _ := utf8.DecodeRune(b[i:])
				title += string(r)
			}

      // title = returnUnicodeString(title)
		}

		writer := post.Find(".gall_writer > .nickname").Text()
		// writer := returnUnicodeString(strangeWriter)

		mixedWriter := writer + " " + post.Find(".gall_writer > .ip").Text()

		if mixedWriter == "" {
			operator, _ := post.Find(".gall_writer").Attr("user_name")
			if operator == "" {
				other, _ := post.Find(".gall_writer>b > .nickname").Attr("title")
				operator = other
			}
			mixedWriter = string(operator)
		}

		dateWriting, _ := post.Find(".gall_date").Attr("title")
		views := post.Find(".gall_count").Text()
		recommend := post.Find(".gall_recommend").Text()

		fmt.Println(id)
		fmt.Println(postType)
		fmt.Println(string([]rune(title)))
		// fmt.Printf("%c\n", title[len(title)-3])
		fmt.Println(replies)
		fmt.Println(mixedWriter)
		fmt.Println(dateWriting)
		fmt.Println(views)
		fmt.Println(recommend)
	})
}

func getTotalPages() int {
	totalPages := 0

	res, err := http.Get(baseURL)

	checkError(err)
	checkStatus(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	checkError(err)

	doc.Find(".bottom_paging_box").Each(func(i int, s *goquery.Selection) {
		totalPages = s.Find("a").Length() - 1
	})

	return totalPages
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
