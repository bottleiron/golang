package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write(toByte("go 웹 서비스 테스트"))
	})

	http.ListenAndServe(":5000", nil)
}

func toByte(text string) []byte {

	toByte := []byte(text)
	return toByte
}
