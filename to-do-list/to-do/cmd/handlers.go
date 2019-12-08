package main

import (
  "fmt"
  "strings"
  "net/http"
  "strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
  l, err := app.list.Latest()
  if err != nil {
    app.serverError(w, err)
    return
  }
  app.render(w, r, "home.page.tmpl", &templateData{
    Items: l,
  })
}

func (app *application) newItemForm(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "create.page.tmpl", &templateData{})
}

func (app *application) newItem(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	item := r.PostForm.Get("item")
  if strings.TrimSpace(item) == "" {
    fmt.Println("Error! empty item")
    return
  }

  err = app.list.Insert(item)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) deleteItem(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":id"))
  if err != nil || id < 1 {
    app.notFound(w)
    return
  }
  ok := app.list.Delete(id)
  if ok != nil {
    app.notFound(w)
    return
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
