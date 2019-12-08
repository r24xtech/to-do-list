package main

import (
  "net/http"
  "github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(app.home))
  mux.Get("/list/add", http.HandlerFunc(app.newItemForm))
  mux.Post("/list/add", http.HandlerFunc(app.newItem))
  mux.Post("/list/delete/:id", http.HandlerFunc(app.deleteItem))

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Get("/static/", http.StripPrefix("/static", fileServer))

  return mux
}
