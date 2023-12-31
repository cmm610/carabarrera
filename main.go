package main

import (
  "carabarrera/web/pages"
  "fmt"
  "github.com/a-h/templ"
  "net/http"
)
func main() {
  page := pages.Index()
  http.Handle("/", templ.Handler(page))
  fmt.Println("Listening on :3000")

  http.ListenAndServe(":3000", nil)
}
