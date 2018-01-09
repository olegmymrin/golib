package main

import (
	"io"
	"net/http"
	//"runtime"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	//runtime.GOMAXPROCS(1)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe("127.0.0.1:30700", nil)
}