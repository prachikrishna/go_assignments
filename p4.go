//given a url,fetch all images and download them concurrently--trying to streamline
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
var sem chan int

func main() {
	url := "https://golang.org"
	fmt.Println("Processing...")
	kr := FindImages1(url)
	//wg.Wait()
	fmt.Println(kr)

}
func FindImages1(url string) string { //[]string
	//result := make([]string, 0)
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
	sem := make(chan int, 10)
	defer close(sem)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {

			for _, img := range n.Attr {
				if img.Key == "src" {
					images := img.Val
					//wg.Add(1)
					//sem := make(chan int, 10)
					//defer close(sem)
					//sem <- 1
					DownloadImages1(images) //to download image as soon as it is available
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	//fmt.Println(result)
	//images := result
	//pk := DownloadImages1(images)
	//return pk
	return "done"

}
func DownloadImages1(images string) {

	wg.Add(1)
	sem := make(chan int, 15)
	defer close(sem)
	//for _, images := range images {
	sem <- 1
	go func(images string) {
		defer wg.Done()
		//defer close(sem)

		tokens := strings.Split(images, "/")
		//fmt.Println(tokens)

		imageName := tokens[len(tokens)-1]
		fmt.Println("Found", images, "as", imageName)

		//chk := tokens[0]
		u, err := url.Parse(images)
		if err != nil {
			panic(err)
		}
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
		<-sem

	}(images)

	wg.Wait()
	//close(sem)
	//pk := "Done"
	//return pk
	//return "done"
}
