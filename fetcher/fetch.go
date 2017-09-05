package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/howie6879/NIYT/common"
	readability "github.com/mauidude/go-readability"
)

// Novel contains information about the novel taht you searched
type Novel struct {
	Title    string
	URL      string
	Chapters []ChapterItem
}

// ChapterItem contains information about the chapter
type ChapterItem struct {
	ChapterName string
	Href        string
	Content     string
}

// LatestChapterItem contains information about the latest chapter
type LatestChapterItem struct {
	LatestChapterName string
	LatestChapterURL  string
}

// FetchResult get the result that you need
func FetchResult(query string) ([]Novel, error) {
	var resultData []Novel
	config, err := common.LoadConfiguration()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return resultData, err
	}
	requestURL := config.SoURL + "?q=" + url.QueryEscape(query) + "&ie=utf-8"
	response, err := common.RequestURL(requestURL)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return resultData, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		doc, _ := goquery.NewDocumentFromReader(io.Reader(response.Body))
		doc.Find(".res-list").Each(func(i int, s *goquery.Selection) {
			var title = s.Find("h3 a").Text()
			var currentURL, _ = s.Find("h3 a").Attr("href")
			if strings.Contains(currentURL, config.DomainFlagM) {
				currentURL, _ = s.Find("h3 a").Attr("data-url")
			}
			if strings.Contains(currentURL, config.DomainFlagL) {
				urlParse, _ := url.Parse(currentURL)
				query, _ := url.ParseQuery(urlParse.RawQuery)
				currentURL = query["url"][0]
			}
			domain := common.ReturnDomain(currentURL)
			isParsed := common.StringInSlice(domain, config.Sites)
			if isParsed {
				currentURL = strings.Replace(currentURL, "index.html", "", -1)
				if !strings.Contains(currentURL, ".html") {
					currentItem := Novel{Title: title, URL: currentURL}
					resultData = append(resultData, currentItem)
				}
			}
		})
	}
	return resultData, nil
}

// FetchContent get the chapter's content
func (chapter *ChapterItem) FetchContent() {
	if len(chapter.Content) > 0 {
		return
	}
	var (
		chapterContent []string
		chapterString  string
		html           string
	)
	response, err := common.RequestURL(chapter.Href)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	html = common.DetectBody(body)
	readabilityDoc, err := readability.NewDocument(html)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	content := readabilityDoc.Content()
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		chapterContent = append(chapterContent, s.Text())
	})
	chapterString = strings.Join(chapterContent, "\n")
	chapterString = strings.Replace(chapterString, "  　　", "\n", -1)
	chapter.Content = chapterString
}

// FetchChapters search all chapters
func (novel *Novel) FetchChapters() {
	if len(novel.Chapters) > 0 {
		return
	}
	var (
		content     string
		chapterData []ChapterItem
		chapters    []string
	)
	response, err := common.RequestURL(novel.URL)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	content = common.DetectBody(body)
	valid := regexp.MustCompile(`<a\s+.*?>.*第?\s*[一二两三四五六七八九十○零百千万亿0-9１２３４５６７８９０]{1,6}\s*[章回卷节折篇幕集].*?</a>`)
	chapters = valid.FindAllString(content, -1)
	chapterString := strings.Join(chapters, ",")
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(chapterString))
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		href, _ := s.Attr("href")
		u, err := url.Parse(href)
		if err != nil {
			log.Fatal(err)
		}
		base, err := url.Parse(novel.URL)
		if err != nil {
			log.Fatal(err)
		}
		href = base.ResolveReference(u).String()
		currentChapterItem := ChapterItem{ChapterName: title, Href: href}
		chapterData = append(chapterData, currentChapterItem)
	})
	chapterData = ReverseSlice(chapterData)
	if len(chapterData) > 9 {
		novel.Chapters = chapterData[:10]
	} else {
		novel.Chapters = chapterData
	}

}

// ReverseSlice https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
func ReverseSlice(s []ChapterItem) []ChapterItem {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
