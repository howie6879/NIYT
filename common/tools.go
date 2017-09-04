package common

import (
	"net/http"
	"net/url"
)

// RequestURL return the search result
func RequestURL(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", GetUserAgent())
	response, err := client.Do(req)
	return response, err
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

//ReturnDomain parse url and return the domain
func ReturnDomain(currentURL string) string {
	urlParse, _ := url.Parse(currentURL)
	domain := urlParse.Host
	return domain
}
