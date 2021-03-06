/*
Copyright (C) 2018 Grathium Sofwares <grathiumsoftwears@gmail.com>
	This program comes with ABSOLUTELY NO WARRANTY
	This is a free software, and you are welcome to redistribute it under certain
	conditions.
*/

package main

import (
    "fmt"
    "flag"
    "io/ioutil"
	"net/http"
	"strings"
)

var website string
var replacer = strings.NewReplacer("/", "#", "?", "#")
func main() {
    flag.StringVar(&website, "w", "", "Usage") // get the website that the user wants to access
    flag.Parse()

    websiteEdit := "https://api.ready.mobi/api/v1/prism/proxy?prismid=1&testid=1qo0&i=0&device=0&url=" + website
    websiteHTML := string(getHTML(websiteEdit)) // get website source
    // correct file name for use
	str := website
    str = replacer.Replace(str)
    
    // create the website file
	err := ioutil.WriteFile(str + ".html", []byte(websiteHTML), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}


// get HTML function
func getHTML(w string) []byte {
	url := w
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	return html
}