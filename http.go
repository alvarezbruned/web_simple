package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm host %s with port %s \n", hostname, port)
 	fmt.Fprintf(w, "I'm host %s with port %s \n", hostname, port)
    })

    http.HandleFunc("/", serveTemplate)
    log.Fatal(http.ListenAndServe(":" + port, nil))
    
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("www", "index.htm")
  fp := path.Join("www", r.URL.Path)

  tmpl, _ := www.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}
