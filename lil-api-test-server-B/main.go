package main

import (
    "fmt"
    "net/http"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello from server B"}`))
}

func main() {
    s := &server{}
    http.Handle("/", s)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

