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
    
    fs := http.FileServer(http.Dir("/app/www"))
    http.Handle("/", fs)
    
    log.Fatal(http.ListenAndServe(":" + port, nil))
    
}
