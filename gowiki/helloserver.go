package main

import(
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hi there, I love %s!",r.URL.Path[1:])
}

//handles all requests to web root '/' with handler
//listens on port 8080 on ANY interface
//ListenAndServe blocks until program is terminated
func main(){
  http.HandleFunc("/",handler)
  http.ListenAndServe(":8080",nil)
}