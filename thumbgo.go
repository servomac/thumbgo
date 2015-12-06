// Thumbgo is an on-demand image resizer
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

// handlerHttpLoader loads the desired url and resizes it
func handlerHttpLoader(resp http.ResponseWriter, r *http.Request) {
    url := "http://" + r.URL.Path[len("/"):]
    reqImage, err := http.Get(url)
    if err != nil {
        log.Printf("ERROR while fetching %s (%v)\n", url, err)
        http.Error(resp, err.Error(), http.StatusInternalServerError)
        return
    }
    defer reqImage.Body.Close()

    resp.Header().Set("Content-Length", fmt.Sprint(reqImage.ContentLength))
    resp.Header().Set("Content-Type", reqImage.Header.Get("Content-Type"))

    if _, err = io.Copy(resp, reqImage.Body); err != nil {
        log.Printf("ERROR while reading %s (%v)\n", url, err)
        http.Error(resp, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    port := 8000
    http.HandleFunc("/", handlerHttpLoader)
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
