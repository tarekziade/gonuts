package main

import (
    "net/http"
    "html/template"
)

type Context struct {
    Url string
}



func renderIndex(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[1:]
    c := Context{Url: title}
    t, _ := template.ParseFiles("app.tmpl")
    t.Execute(w, c)
}


func main() {
    http.HandleFunc("/", renderIndex)
    http.ListenAndServe(":8282", nil)
}

