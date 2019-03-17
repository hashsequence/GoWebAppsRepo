package main

import (
  "fmt"
  "net/http"
  "log"
)

func handler (writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  if err := http.ListenAndServe("localhost:8080", nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }

}
