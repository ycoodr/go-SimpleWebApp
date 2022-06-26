package main

import (
	"fmt"
	"net/http"
)


func handlefunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Simple go web app</h1>")
}


func main(){
	http.HandleFunc("/", handlefunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}