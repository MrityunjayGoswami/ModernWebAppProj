package main

import (
	"encoding/json"
	"log"
)

type Person struct{
	FirstNmae string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main(){
	myJson:= `
[
	{
		"first_name": "Clark",
		"last_name": "Kent"
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne"
	}
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson),&unmarshalled)
	if err != nil{
		log.Println("Erroe unmarshling Json",err)
	}
	log.Printf("Unmarshalled: %v", unmarshalled)
}