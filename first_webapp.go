package main

import(
    "fmt"
    "net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello world,%s!", request.URL.Path[1:])
}


func main() {
    http.HandleFunc("/",handler)
    http.ListenAndServe(":8082",nil)
}
