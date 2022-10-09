package main

import (
	"errors"
	"fmt"
	"net/http"
)
var portNumber = ":8080"
func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"This is home page")
}
func About(w http.ResponseWriter, r *http.Request)  {
	sum:= AddValues(2,2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page and 2+2 is %d",sum))
}
func AddValues(x,y int) int {
	return x+y
}
func Divide(w http.ResponseWriter, r *http.Request)  {
	var x,y float32 =100.0,0.0
	f, err := divideValues(x,y)
	if err != nil{
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %F",x,y,f))
}
func divideValues(x,y float32) (float32, error) {
	if y<=0{
		err := errors.New("Cannot be divided by zero")
		return 0, err
	}
	return x/y, nil
}

func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

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