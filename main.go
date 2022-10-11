package main

import (
	"fmt"
	"net/http"
	"html/template"
)
var portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request)  {
	
}
func renderTemplate(w http.ResponseWriter, tmpl string)  {
	parsedTemplate, _ := template.ParseFiles("./templates" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("error parsing template : ",err)
		return
	}
}


func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n,err := fmt.Fprintf(w,"Hello, World")
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bites written : %d",n))
	// })
	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}