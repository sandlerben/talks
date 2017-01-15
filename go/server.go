package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey!!!"))
}

func main() {
	http.HandleFunc("/", hello)
	for i := 0; i < 10; i++ {
		go http.ListenAndServe(fmt.Sprintf(":808%d", i), nil)
	}

	http.ListenAndServe(":8080", nil)
}
