package main

import (
	//"fmt"
	//"github.com/fzzy/radix/redis"
	"html/template"
	"log"
	"net/http"
)

//var db *redis.Client

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTemplate := template.Must(template.ParseFiles("home.html"))
	homeTemplate.Execute(c, nil)
}

func staticHandler(c http.ResponseWriter, req *http.Request) {
	http.ServeFile(c, req, req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/", staticHandler)

	/*conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("redis:", err)
	}

	db = conn*/

	log.Println("Server started!")

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
