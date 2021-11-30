package titulosPageHTML

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtém o título de uma página HTML
func Titulo(urls ...string) <-chan string {
	canal := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			regex, _ := regexp.Compile("<title>(.*?)<\\/title>")
			canal <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return canal
}
