package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beefsack/go-rate"
)

var rl *rate.RateLimiter

func hello(w http.ResponseWriter, req *http.Request) {

	if ok, _ := rl.Try(); ok {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Success"))
		fmt.Println("Success")
	} else {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("This query timed out. Sorry"))
		fmt.Println("Failure")
	}

}

func main() {
	rl = rate.New(26, time.Second)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":10001", nil)
}
