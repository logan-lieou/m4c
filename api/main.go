package main

import (
  "net/http"
  "fmt"
  "log"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

type ProductProperties struct {
  name string
  date string
  price int
  id int
}

func HomeLink(res http.ResponseWriter, req *http.Request) {
  // send our file to the request
  res.Write([]byte("Hello world!"))
}

func ProductsHandler(res http.ResponseWriter, req *http.Request) {
  // write a response
  fmt.Fprintf(res, "products endpoint!")
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", HomeLink)
  router.HandleFunc("/api/products", ProductsHandler)
  log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}
