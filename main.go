package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println(("Hello World"));

	//builtin app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w,"Hello World");
		if err!=nil{
			//write the error to the console
			fmt.Println((err));
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	//start a web server that listens to the requests
	_=http.ListenAndServe(":8080",nil);
}
