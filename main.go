package main

import (
	"fmt"
	"gogogo/go01"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/sayHi", SayHi)
	http.HandleFunc("/jiujiu", go01.JiujiuBiao)
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatal("ListenAndServe", er)
	}
}

//localhost:8080
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

//localhost:8080/sayHi?name=xiaoming
func SayHi(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprint(w, "Hi!", name)
}
