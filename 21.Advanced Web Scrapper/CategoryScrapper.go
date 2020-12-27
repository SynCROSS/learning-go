package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://gall.dcinside.com/"

type extractedCategory struct {
	gallName string
	url      string
}

var cateListNum = []int{2, 26, 21, 18, 22, 12, 31, 7, 23, 39, 24, 40, 4, 41, 19, 13, 27, 20, 29, 8, 15, 14, 11, 6, 17, 16, 37, 28, 5, 9, 30, 35, 10, 38, 3, 25, 33}

func main() {
	var categorySlice []extractedCategory
	c := make(chan []extractedCategory)

	for _, num := range cateListNum {
		go getCategoryName(num, c)
	}

	for i := 0; i < len(cateListNum); i++ {
		extractedCategory := <-c
		categorySlice = append(categorySlice, extractedCategory...)
	}

	fmt.Println(len(categorySlice))
}

func getCategoryName(num int, mainChannel chan<- []extractedCategory) {
	var categorySlice []extractedCategory
	c := make(chan extractedCategory)

	res, err := http.Get(baseURL)

	checkError(err)
	checkStatus(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)

	defer res.Body.Close()
	checkError(err)

	categories := doc.Find("#categ_listwrap")

	categories.Each(func(i int, category *goquery.Selection) {
		go extractCategory(category, num, "bind", c)
		go extractCategory(category, num, "", c)
	})

	for len(cateListNum) > 0 {
		cate := <-c
		categorySlice = append(categorySlice, cate)
	}

	mainChannel <- categorySlice
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

func extractCategory(category *goquery.Selection, num int, queryPath string, c chan<- extractedCategory) {

	// * 2, 26, 21, 18, 22
	// * 12, 31, 7, 23, 39, 24, 40, 4, 41, 19, 13, 27, 20, 29, 8, 15, 14, 11, 6, 17, 16, 37, 28, 5, 9, 30, 35, 10, 38, 3, 25, 33

	findingPath := queryPath + " .cate_box > #cateList" + strconv.Itoa(num) + " > .cate > ul > li > a" + ""

	gallName := category.Find(findingPath).Text()
	url, _ := category.Find(findingPath).Attr("href")

	c <- extractedCategory{
		gallName: gallName,
		url:      url,
	}
}

// func exportPosts(postSlice []extractedCategory) {
// 	file, err := os.Create("DCInside.csv")

// 	checkError(err)

// 	w := csv.NewWriter(file)

// 	headers := []string{
// 		"id",
// 		"postType",
// 		"title",
// 		"replies",
// 		"writer",
// 		"dateWriting",
// 		"views",
// 		"recommend",
// 	}

// 	writeErr := w.Write(headers)

// 	checkError(writeErr)

// 	for _, post := range postSlice {
// 		extractedCategorySlice := []string{
// 			post.id,
// 			post.postType,
// 			post.title,
// 			post.replies,
// 			post.writer,
// 			post.dateWriting,
// 			post.views,
// 			post.recommend,
// 		}
// 		postWritingErr := w.Write(extractedCategorySlice)

// 		checkError(postWritingErr)
// 	}

// 	defer func() {
// 		w.Flush()
// 		fmt.Println("File Writing is Successfully Done:", len(postSlice))
// 	}()
// }
