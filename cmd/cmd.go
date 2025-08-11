package cmd


import (
    "fmt"
    "net/http"
)

func Serve(port string) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
    })
    http.ListenAndServe(port, nil)
}