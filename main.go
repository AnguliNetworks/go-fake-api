package main

import (
	"strings"
	"io/ioutil"
	"net/http"
	"fmt"
	"log"
)

var formats = [2]*Format{
	{Name: "json", FileSuffix: "json", MIMEType: "application/json"},
	{Name: "xml", FileSuffix: "xml", MIMEType: "application/xml"},
}
var errorMessage = "{\"error\":\"%s\"}"
var errorFormat = formats[0]
var path = "endpoints/${url}.${format}"

type Format struct {
	Name       string
	FileSuffix string
	MIMEType   string
}

type Endpoint struct {
	Content []byte
	Format  *Format
}

func load(url string) (*Endpoint, error) {

	var content []byte
	var err error
	format := errorFormat

	for _, element := range formats {

		filename := strings.Replace(path, "${url}", url, -1)
		filename = strings.Replace(filename, "${format}", element.FileSuffix, -1)

		content, err = ioutil.ReadFile(filename)

		if err == nil {
			format = element
			break
		}
	}

	return &Endpoint{Content: content, Format: format}, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	endpoint, err := load(r.URL.Path)

	w.Header().Set("Content-Type", endpoint.Format.MIMEType)

	if err != nil {
		fmt.Fprintf(w, errorMessage, err)
	} else {
		fmt.Fprintf(w, "%s", endpoint.Content)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
