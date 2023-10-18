package main

import (
	"fmt"
	"net/http"

	"github.com/otium/queue"
)

func main() {
	q := queue.NewQueue(func(val interface{}) {
		fmt.Printf("Hi there, I received %d!\n", val)
	}, 20)

	go q.Wait()

	http.HandleFunc("/w", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 200; i++ {
			q.Push(i)
		}
	}) // http://127.0.0.1:8080/Go
	http.ListenAndServe(":8080", nil)
}
