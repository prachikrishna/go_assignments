//given a url,fetch all images and download them concurrently
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"sync"

	"golang.org/x/net/html"
)

var wg sync.WaitGroup

func main() {
	url := "https://golang.org"
	fmt.Println("Processing...")
	message := FindImages1(url)
	fmt.Println(message)

}

//this function fetches all image urls from the given url and stores them in result
func FindImages1(url string) string { //[]string
	result := make([]string, 0)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	// reads html as a slice of bytes
	html1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	new_html := string(html1) // convert slice of bytes to string

	//fmt.Printf("%s\n", new_html)

	//parsing
	doc, err := html.Parse(strings.NewReader(new_html))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	//fmt.Println(result)
	images := result
	//we pass the image urls to DownloadImages1 which is used to
	//concurrently download all images with valid urls
	//and returns a completion message after the images
	//have been successfully downloaded
	message := DownloadImages1(images)
	return message

}
func DownloadImages1(images []string) string {

	wg.Add(len(images))
	//we have used a buffered channel named limit_concurrency
	//to implement bounded concurrency so that we can
	//abstain from generating too many requests at any given instance
	limit_concurrency := make(chan int, 10)
	defer close(limit_concurrency)
	for _, images := range images {
		limit_concurrency <- 1
		go func(images string) {
			defer wg.Done()

			tokens := strings.Split(images, "/")
			//fmt.Println(tokens)

			imageName := tokens[len(tokens)-1]
			fmt.Println("Found", images, "as", imageName)

			//to check valid url
			u, err := url.Parse(images)
			if err != nil {
				panic(err)
			}
			//and create a file to download image if it has https scheme
			if u.Scheme == "https" {
				output, err := os.Create(imageName)
				if err != nil {
					log.Fatal(err)
				}
				defer output.Close()

				res, err := http.Get(images)
				if err != nil {
					log.Fatal(err)
				} else {
					defer res.Body.Close()
					_, err = io.Copy(output, res.Body)
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Downloaded", imageName)
					}
				}
			}
			<-limit_concurrency

		}(images)
	}
	wg.Wait()
	return "done"
}
