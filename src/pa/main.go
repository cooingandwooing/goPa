package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func qiubai_parse(url string) bool {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	// array := make([]map[string]string, 100)
	hash := make(map[string]string)
	fmt.Println(doc.Text())
	fmt.Println(url + " 大小：")

	fmt.Println(doc.Find(".yw-con li a").Length())

	if doc.Find(".yw-con li a ").Length() > 0 {
		doc.Find(".yw-con li a ").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			hash["http://www.sxdi.gov.cn"+url] = s.Text()
			// array = append(array, hash)
		})
		print(hash)
		printTxt(hash)
		printZF(hash)
		return true
	} else {
		return false
	}
}

func print(s map[string]string) {
	fileName := "审查调查.txt"
	dstFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //打开文件
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	if s != nil {
		for href, t := range s {
			dstFile.WriteString(href + " : " + t + "\n")
		}
	}
}
func printTxt(s map[string]string) {
	fileName := "审查调查(仅文本).txt"
	dstFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //打开文件
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	if s != nil {
		for _, t := range s {
			dstFile.WriteString(t + "\n")
		}
	}
}
func printZF(s map[string]string) {
	fileName := "审查调查(法院检察院仅文本).txt"
	dstFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //打开文件
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	if s != nil {
		for _, t := range s {
			if strings.Contains(t, "法院") || strings.Contains(t, "检察院") {
				dstFile.WriteString(t + "\n")
			}
		}
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func do() {
	sggbZjsc := "http://www.sxdi.gov.cn/scdc/sggb/zjsc/"         // 249
	sggbDjzwcf := "http://www.sxdi.gov.cn/scdc/sggb/djzwcf/"     // 250
	shiggbZjsc := "http://www.sxdi.gov.cn/scdc/shiggb/zjsc/"     // 251
	shiggbDjzwcf := "http://www.sxdi.gov.cn/scdc/shiggb/djzwcf/" // 252

	url := make(map[string]string)
	url[sggbZjsc] = "249"
	url[sggbDjzwcf] = "250"
	url[shiggbZjsc] = "251"
	url[shiggbDjzwcf] = "252"

	for href, suffix := range url {
		var num bool = true
		for i := 1; num; i++ {
			if i == 1 {
				num = qiubai_parse(href)
			}
			if i > 1 {
				// list_251_11.html
				num = qiubai_parse(href + "list_" + suffix + "_" + strconv.Itoa(i) + ".html")
			}
			fmt.Println(suffix + " " + strconv.Itoa(i) + strconv.FormatBool(num))
		}
	}
}
func main() {
	do()
}

func test(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// array := make([]map[string]string, 100)
	hash := make(map[string]string)

	fmt.Println(doc.Find(".yw-con li a").Length())

	if doc.Find(".yw-con li a ").Length() > 0 {
		doc.Find(".yw-con li a ").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			hash["https://www.qiushibaike.com"+url] = s.Text()
			// array = append(array, hash)
		})
		print(hash)
	}
}
