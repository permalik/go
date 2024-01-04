package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "net/http"
	"os"
	"path"
    "strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type App struct {
     NoteHandler *NoteHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    var head string
    head, _ = bisectPath(req.URL.String())
    if head == "note" {
        h.NoteHandler.HandleHTTP(res, req)
        return
    }
    http.Error(res, "Not Found", http.StatusNotFound)
}

type NoteHandler struct {
}

func (h *NoteHandler) HandleHTTP(res http.ResponseWriter, req *http.Request) {
    // var head string
    var tail string
    var strId string
    fmt.Println(req.URL.String())
    _, tail = bisectPath(req.URL.String())
    strId, _ = bisectPath(tail)
    id, err := strconv.Atoi(strId)
    if err != nil {
        http.Error(res, fmt.Sprintf("Invalid note id %q", tail), http.StatusBadRequest)
        return
    }
    switch req.Method {
    case "GET":
    fmt.Printf("handleGet(%d)", id)
    case "PUT":
    fmt.Printf("handlePut(%d)", id)
    default:
        http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
}
}

func bisectPath(p string) (head, tail string) {
    p = path.Clean("/" + p)

    i := strings.Index(p[1:], "/") + 1

    if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}

func main() { 
    app := &App{
        NoteHandler: new(NoteHandler),
    }

    err := http.ListenAndServe(":9000", app)
    if err != nil {
        log.Fatal("errrr", err)
    }

    // fs := http.FileServer(http.Dir("./static"))
    // http.Handle("/static/", http.StripPrefix("/static/", fs))

    // connect()
    os.Remove("./reminder.db")
    fmt.Println("db purged")

    db, err := sql.Open("sqlite", "./reminder.db")
    fmt.Println("db created")
    if err != nil {
        log.Fatal("Error creating db instance", err)
    }
    defer db.Close()

    // createTables()
    sqlStmt := `
    create table reminder (
    id integer PRIMARY KEY,
    title varchar(255),
    body text
    );
    delete from reminder;
    `

    _, err = db.Exec(sqlStmt)
    fmt.Println("table created")
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }

    // createReminder()

    tx, err := db.Begin()
    fmt.Println("transaction created")
    if err != nil {
        log.Fatal(err)
    }

    stmt, err := tx.Prepare("insert into reminder(title, body) values(?, ?)")
    fmt.Println("statement prepared")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

   //  func (r *Reminder) create() error {
   //      // persist to db
   //      return nil
   //  }

    type Reminder struct {
        Title string
        Body string
    }

    r := &Reminder{
        Title: "NewTitle",
        Body: "NuevoBodissimo",
    }

    _, err = stmt.Exec(r.Title, r.Body)
    if err != nil {
        log.Fatal(err)
    }

    err = tx.Commit()
    fmt.Println("data input")
    if err != nil {
        log.Fatal(err)
    }

    // getReminder()
    rows, err := db.Query("select id, title, body from reminder")
    fmt.Println("query created")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // logData()
    for rows.Next() {
        var id int
        var title string
        var body string
        err = rows.Scan(&id, &title, &body)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, title, body)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}

