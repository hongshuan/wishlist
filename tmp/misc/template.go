// This is a trivial web server that uses the text/template (and html/template) 
// package's "block" feature to implement a kind of template inheritance.
package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image/", imageHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// indexTemplate is the main site template.
// The default template includes two template blocks ("sidebar" and "content")
// that may be replaced in templates derived from this one.
var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

// Index is a data structure used to populate an indexTemplate.
type Index struct {
	Title string
	Body  string
	Links []Link
}

type Link struct {
	URL, Title string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Image gallery",
		Body:  "Welcome to the image gallery.",
	}
	for name, img := range images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// imageTemplate is a clone of indexTemplate that provides
// alternate "sidebar" and "content" templates.
var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("image.tmpl"))

// Image is a data structure used to populate an imageTemplate.
type Image struct {
	Title string
	URL   string
}

// imageHandler is an HTTP handler that serves the image pages.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := images[strings.TrimPrefix(r.URL.Path, "/image/")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// images specifies the site content: a collection of images.
var images = map[string]*Image{
	"go":     {"The Go Gopher", "https://golang.org/doc/gopher/frontpage.png"},
	"google": {"The Google Logo", "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"},
}

// index.tmpl
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}</title>
		<style>
			body { font-family: sans-serif; }
			h1 { background: #ddd; }
			#sidebar { float: right; }
		</style>
	</head>
	<body>
		<h1>{{.Title}}</h1>

		<div id="sidebar">
			{{block "sidebar" .}}
			<h2>Links</h2>
			<ul>
				{{- range .Links}}
				<li><a href="{{.URL}}">{{.Title}}</a></li>
				{{- end}}
			</ul>
			{{end}}
		</div>

		{{block "content" .}}
		<div id="content">{{.Body}}</div>
		{{end}}
	</body>
</html>

// image.tmpl
{{define "content"}}
<div id="image">
	<img src="{{.URL}}" title="{{.Title}}" alt="{{.Title}}">
</div>
{{end}}

{{define "sidebar"}}
<a href="/">Back</a>
{{end}}
