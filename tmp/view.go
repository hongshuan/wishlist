package views

import (  
  "html/template"
  "net/http"
)

func NewView(layout string, files ...string) *View {  
  files = append(files, layoutFiles()...)
  t, err := template.ParseFiles(files...)
  if err != nil {
    panic(err)
  }

  return &View{
    Template: t,
    Layout:   layout,
  }
}

type View struct {  
  Template *template.Template
  Layout   string
}

type ViewData struct {  
  Flashes map[string]string
  Data    interface{}
}

var flashRotator int = 0

func flashes() map[string]string {  
  flashRotator = flashRotator + 1
  if flashRotator%3 == 0 {
    return map[string]string{
      "warning": "You are about to exceed your plan limts!",
    }
  } else {
    return map[string]string{}
  }
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {  
  vd := ViewData{
    Flashes: flashes(),
    Data:    data,
  }
  return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

/*
func (v *View) Render(w http.ResponseWriter, data interface{}) error {  
  return v.Template.ExecuteTemplate(w, v.Layout, data)
}
*/

/*
package main

import (  
  "net/http"

  "calhoun.io/views"
)

var index *views.View  
var contact *views.View

func main() {  
  index = views.NewView("bootstrap", "views/index.gohtml")
  contact = views.NewView("bootstrap", "views/contact.gohtml")

  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/contact", contactHandler)
  http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {  
  index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {  
  contact.Render(w, nil)
}
*/

{{define "bootstrap"}}
<!DOCTYPE html>  
<html lang="en">  
  <head>
    <title>Calhoun.io</title>
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  </head>

  <body>

    <div class="container-fluid">
      <!-- Our content will go here. -->
      {{template "flashes" .}}
      {{template "yield" .}}
    </div>

    <!-- jquery & Bootstrap JS -->
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  </body>
</html>  
{{end}}

{{define "flashes"}}
  {{range $key, $value := .Flashes}}
    <div class="alert alert-{{$key}}">
      {{$value}}
    </div>
  {{end}}
{{end}}

{{define "yield"}}
  <h1>Filler header</h1>
  <p>Filler paragraph</p>
{{end}}
