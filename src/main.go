package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	http.HandleFunc("/", WelcomeHandler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	Name string
	Age string
}

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		t, err := template.ParseFiles("welcome.html")
		check(err)
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.Age = r.Form.Get("entered_age")
		t, err := template.ParseFiles("welcomeresp.html")
		check(err)
		t.Execute(w, myUser)
	}
}