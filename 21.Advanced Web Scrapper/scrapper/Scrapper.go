package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://gall.dcinside.com/board/lists/?id=baseball_new9"

type extractedPosts struct {
	id          string
	postType    string
	title       string
	replies     string
	writer      string
	dateWriting string
	views       string
	recommend   string
}

func main(searchString string) {
	var postSlice []extractedPosts
	c := make(chan []extractedPosts)
	totalPages := getTotalPages()

	for i := 0; i < totalPages; i++ {
		go getPost(i+1, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedPost := <-c
		postSlice = append(postSlice, extractedPost...)

	}

	exportPosts(postSlice)
}

func getPost(pageNum int, mainChannel chan<- []extractedPosts) {
	var postSlice []extractedPosts
	c := make(chan extractedPosts)
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
		go extractPost(post, c)
	})

	for i := 0; i < posts.Length(); i++ {
		p := <-c
		postSlice = append(postSlice, p)
	}

	mainChannel <- postSlice
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

func decodeString(str string) string {
	b := []byte(str)
	str = ""

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		str += string(r)
		b = b[size:]
	}

	return str
}

func extractPost(post *goquery.Selection, c chan<- extractedPosts) {
	id := post.Find(".gall_num").Text()
	titleClass, _ := post.Find(".gall_tit > a > em").Attr("class")
	titleType := []rune(titleClass)
	postType := string(titleType[14:17])
	title := post.Find(".gall_tit > a ").Text()
	replies := post.Find(".gall_tit > .reply_numbox > .reply_num").Text()

	if replies == "" {
		replies = "[0]"
	} else {
		title = title[:len(title)-3]
		title = decodeString(title)
	}

	writer := post.Find(".gall_writer > .nickname").Text()
	ip := post.Find(".gall_writer > .ip").Text()
	mixedWriter := writer + " " + ip

	if mixedWriter == " " {
		operator, _ := post.Find(".gall_writer").Attr("user_name")
		if operator == "" {
			other, _ := post.Find(".gall_writer>b > .nickname").Attr("title")
			if other == "" {
				another, _ := post.Find(".gall_writer").Attr("data-nick")
				other = another
			}
			operator = other
		}
		mixedWriter = decodeString(operator)
	}

	dateWriting, _ := post.Find(".gall_date").Attr("title")

	if dateWriting == "" {
		date := post.Find(".gall_date").Text()
		t, err := time.Parse("06/02/02", date) // ! must write like that jan 2 2006
		checkError(err)
		dateWriting = t.Format("2006-02-02")
	} else {
		dateWriting = dateWriting[:10]
	}

	views := post.Find(".gall_count").Text()
	recommend := post.Find(".gall_recommend").Text()

	c <- extractedPosts{
		id:          id,
		postType:    postType,
		title:       title,
		replies:     replies,
		writer:      mixedWriter,
		dateWriting: dateWriting,
		views:       views,
		recommend:   recommend}
}

func exportPosts(postSlice []extractedPosts) {
	file, err := os.Create("DCInside.csv")

	checkError(err)

	w := csv.NewWriter(file)

	headers := []string{
		"id",
		"postType",
		"title",
		"replies",
		"writer",
		"dateWriting",
		"views",
		"recommend",
	}

	writeErr := w.Write(headers)

	checkError(writeErr)

	for _, post := range postSlice {
		extractedPostsSlice := []string{
			post.id,
			post.postType,
			post.title,
			post.replies,
			post.writer,
			post.dateWriting,
			post.views,
			post.recommend,
		}
		postWritingErr := w.Write(extractedPostsSlice)

		checkError(postWritingErr)
	}

	defer func() {
		w.Flush()
		fmt.Println("File Writing is Successfully Done:", len(postSlice))
	}()
}
