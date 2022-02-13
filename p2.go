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

	result := make([]string, 0)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Fatal(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	new_html := string(html1)
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", new_html)

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
	wg.Add(len(result))
	for _, result := range result {
		go func(result string) {
			defer wg.Done()
			tokens := strings.Split(result, "/")
			//fmt.Println(tokens)

			imageName := tokens[len(tokens)-1]
			fmt.Println("Found", result, "as", imageName)

			chk := tokens[0]
			if chk == "https:" {
				output, err := os.Create(imageName)
				if err != nil {
					log.Fatal(err)
				}
				defer output.Close()

				res, err := http.Get(result)
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

		}(result)
	}
	wg.Wait()
	fmt.Println("Done")

}
