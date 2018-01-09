package test8

import (
	"net/http"
	"testing"
	"sync"
	"bytes"
	"fmt"
)

var addr = "127.0.0.1:55000"

func createServer(check func(r *http.Request)) *http.Server {
	srv := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			check(r)
		}),
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
	return srv
}

func TestContentLength(t *testing.T) {
	srv := createServer(func(r *http.Request){
		fmt.Printf("%+v\n", r.Header)
		if r.Header.Get("Content-Length") == "" {
			t.Fatal("Content-Length is empty")
		}
	})
	defer srv.Close()
	req, err := http.NewRequest("GET", "http://"+addr, bytes.NewReader([]byte("aaa")))

	if err != nil {
		t.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
}