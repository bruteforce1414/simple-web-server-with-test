package main

import (
	"fmt"
	"github.com/bruteforce1414/study-go/testChannels/say"
	"net/http"


)


func main() {

	var elementPath string;

	elementPath="/saygoodbye"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/sayhello"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/saynothing"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/sayeverything"
	http.HandleFunc(elementPath, say.Say)

	fmt.Println("Server is listening...")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}