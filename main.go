package main

import (
  "fmt"
  "net/http"
  "path/filepath"

  "carabarrera/internal/server"
  "carabarrera/web/pages"

  "github.com/a-h/templ"
  "github.com/go-chi/chi/v5"
)

var count = 0

func main() {
  page := pages.Index()
  resume := pages.Resume()
  myWork := pages.MyWork()

  r := chi.NewRouter()
  r.Get("/", templ.Handler(page).ServeHTTP)
  r.Get("/resume", templ.Handler(resume).ServeHTTP)
  r.Get("/my_work", templ.Handler(myWork).ServeHTTP)

  workDir := "./web/static"
  filesDir := http.Dir(filepath.Join(workDir, "/"))
  server.FileServer(r, "/public", filesDir)
  http.ListenAndServe(":3000", r)

  fmt.Println("Listening on :3000")
}
