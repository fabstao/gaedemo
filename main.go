package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	Title string
	Body  []byte
}

func view(w http.ResponseWriter, r *http.Request) {
	pv := &page{Title: "Demo Google App Engine", Body: []byte("Portal Web Golang, if you can see this yell hooray!")}
	t, _ := template.ParseFiles("index.html")
	fmt.Println(pv.Title)
	fmt.Println(string(pv.Body))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, pv)

}

func imagenes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, r.URL.Path[1:])
	fmt.Println(r.UserAgent())
	fmt.Println("Sirviendo Contenido")
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:])
}

func init() {
	http.HandleFunc("/", view)
	http.HandleFunc("/img/", imagenes)
	//http.ListenAndServe(":8188", nil)
}
