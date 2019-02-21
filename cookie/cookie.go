package main

import (
	"fmt"
	"net/http"
)

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	sessid, err := r.Cookie("SESSID")
	if err == nil {
		fmt.Fprintln(w, "SESSID value: "+sessid.Value)
	} else {
		fmt.Println(w, "No data found")
	}
}

func WriteCookie(w http.ResponseWriter, r *http.Request) {
	sessid, err := r.Cookie("SESSID")
	if err == nil {
		fmt.Fprintln(w, "SESSID value: ", sessid.Value)
	}
}

func RemoveCookie(w http.ResponseWriter, r *http.Request) {
	sessid, err := r.Cookie("SESSID")
	if err == nil {
		fmt.Fprintln(w, "SESSID value: ", sessid.Value)
	}
}

func main() {
	http.HandleFunc("/cookie/read", ReadCookie)
	http.HandleFunc("/cookie/write", WriteCookie)
	http.HandleFunc("/cookie/remove", RemoveCookie)

	fmt.Println("Server is running...")
	fmt.Println("Press Ctrl+C to exit")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
