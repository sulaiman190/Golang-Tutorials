package main

import (
    "fmt"
    "log"
    "net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
         http.ServeFile(w, r, "forms.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        name := r.FormValue("name")
        address := r.FormValue("address")
        subject := r.FormValue("subject")
        fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Address = %s\n", address)
        fmt.Fprintf(w, "subject = %s\n", subject)
    default:
        fmt.Fprintf(w, "GET and POST methods are supported.")
    }
}

func main() {
    http.HandleFunc("/", hello)

    fmt.Printf("Starting server 8080...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
