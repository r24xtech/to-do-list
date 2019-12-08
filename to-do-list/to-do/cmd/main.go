// go mod init r24xtech.net/to-do
package main

import(
  "fmt"
  "database/sql"
  "log"
  "net/http"
  "html/template"
  //"r24xtech.net/to-do/model"
  "r24xtech.net/to-do/model/mysql"
  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  list *mysql.ListModel
  templateCache map[string]*template.Template
}

func main() {
  db, err := openDB("root:Ax7zMa!po3W!Lk@m2drRc9kgY8s@/to_do_list?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  templateCache, err := newTemplateCache("./ui/html")
  if err != nil {
    log.Fatal(err)
  }

  app := &application{
    list: &mysql.ListModel{DB: db},
    templateCache: templateCache,
  }

  srv := &http.Server{
    Addr: ":4000",
    Handler: app.routes(),
  }
  fmt.Println("Starting server on :4000")
  err = srv.ListenAndServe()
  log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
