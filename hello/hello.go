package main 

import (
    "net/http"
    "fmt"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func main(){
    http.HandleFunc("/", HomePage)

    fmt.Println("Server is running...")
    fmt.Println("Press Ctrl+C to exit")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

