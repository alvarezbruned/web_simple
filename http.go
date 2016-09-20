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
    fs := http.FileServer(http.Dir("/app/www"))
    http.Handle("/", fs)
    http.HandleFunc("/", serveTemplate)
    log.Fatal(http.ListenAndServe(":" + port, nil))
    
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("/app/www", "index.htm")
  fp := path.Join("/app/www", r.URL.Path)

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}
