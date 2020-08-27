package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Woa, Go is very neat!</h1>")
// 	fmt.Fprintf(w, "<p>Woa, Go is very neat!</p>")
// 	fmt.Fprintf(w, "<p>Woa, Go is very neat!</p>")
// 	fmt.Fprintf(w, "<p>You can %s pass %s</p>", "can", "<strong>variable</strong>")
// }

// func about_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Designed by kashif")
// }

//for xml parsing
type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

//to get string url without curly braces
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	//Unmarshal xaml
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

	// //create a handler
	// http.HandleFunc("/", index_handler)
	// http.HandleFunc("/about/", about_handler)
	// //creating a server
	// http.ListenAndServe(":8000", nil)
}
