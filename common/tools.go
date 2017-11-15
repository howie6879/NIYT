package common

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/saintfish/chardet"
)

// RequestURL return the search result
func RequestURL(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", GetUserAgent())
	response, err := client.Do(req)
	return response, err
}

// QuickestURL return the quickest response
func QuickestURL(index int, url string) int {
	// start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return -1
	}
	// timeUsed := strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64)
	// fmt.Println(url, timeUsed)
	return index
}

// DetectBody gbk convert to utf-8
func DetectBody(body []byte) string {
	var bodyString string
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(body)
	if err != nil {
		return string(body)
	}
	if strings.Contains(strings.ToLower(result.Charset), "utf") {
		bodyString = string(body)
	} else {
		bodyString = mahonia.NewDecoder("gbk").ConvertString(string(body))
	}
	return bodyString
}

// StringInSlice search for an element in a golang slice
func StringInSlice(domain string, list []string) bool {
	for _, eachDomain := range list {
		if domain == eachDomain {
			return true
		}
	}
	return false
}

// ReturnDomain parse url and return the domain
func ReturnDomain(currentURL string) string {
	urlParse, _ := url.Parse(currentURL)
	domain := urlParse.Host
	return domain
}
