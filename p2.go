//given a url,fetch all images and download them concurrently
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var wg sync.WaitGroup

func main() {
	url := "https://golang.org"
	images := findImages(url)
	downloadImages(images)
}
func findImages(url string) []string {
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
	return result
}
func downloadImages(images []string) {

	wg.Add(len(images))
	for _, images := range images {
		go func(images string) {
			defer wg.Done()
			tokens := strings.Split(images, "/")
			//fmt.Println(tokens)

			imageName := tokens[len(tokens)-1]
			fmt.Println("Found", images, "as", imageName)

			chk := tokens[0]
			if chk == "https:" {
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

		}(images)
	}
	wg.Wait()
	fmt.Println("Done")
}
