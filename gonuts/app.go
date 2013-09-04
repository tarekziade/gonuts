package main

import (
    "net/http"
    "html/template"
)

type Context struct {
    Url string
    Hits map[string]int
}


var hits = make(map[string]int)

func renderIndex(w http.ResponseWriter, r *http.Request) {
    var url = r.URL.Path[1:]
    if url == "" {
      url = "/"
    }

    hits[url] += 1
    c := Context{Url: url, Hits: hits}
    t, _ := template.ParseFiles("app.tmpl")
    t.Execute(w, c)
}


func resetHits(w http.ResponseWriter, r *http.Request) {
    hits = make(map[string]int)
    http.Redirect(w, r, "/", http.StatusFound)
}



func main() {
    http.HandleFunc("/", renderIndex)
    http.HandleFunc("/reset", resetHits)
    http.ListenAndServe(":8282", nil)
}

