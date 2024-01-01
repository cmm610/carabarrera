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
func main() {
  page := pages.Index()
  r := chi.NewRouter()
  r.Get("/", templ.Handler(page).ServeHTTP)

  workDir := "./web/static"
  filesDir := http.Dir(filepath.Join(workDir, "/"))
  server.FileServer(r, "/public", filesDir)
  http.ListenAndServe(":3000", r)

  fmt.Println("Listening on :3000")
}
