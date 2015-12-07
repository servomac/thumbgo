// Thumbgo is an on-demand image resizer
package main

import (
    "fmt"
    "log"
    "net/http"
    "regexp"
    "strconv"

    "github.com/servomac/thumbgo/loader"
    "github.com/servomac/thumbgo/image"
)

var validPath = regexp.MustCompile("^/([0-9]+)x([0-9]+)/(.+)$")


// handler loads the desired url and resizes it
func handler(resp http.ResponseWriter, r *http.Request) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        log.Printf("ERROR invalid path %s\n", r.URL.Path)
        http.NotFound(resp, r)
        return
    }

    url := "http://" + m[3]
    width, _ := strconv.Atoi(m[1])
    height, _ := strconv.Atoi(m[2])
    options := image.ImageOptions{Width: width, Height: height}

    // loader
    imageBody, err := loader.HttpLoader(url)
    if err != nil {
        log.Printf("ERROR while loading %s (%v)\n", url, err)
        http.Error(resp, err.Error(), http.StatusInternalServerError)
        return
    }

    // processing
    resizedImage, err := image.Resize(imageBody, options)
    nbytes, err := resp.Write(resizedImage.Body)
    if err != nil {
        log.Printf("ERROR while serving %s (%v)\n", url, err)
        http.Error(resp, err.Error(), http.StatusInternalServerError)
        return
    }

    resp.Header().Set("Content-Length", string(nbytes))
    resp.Header().Set("Content-Type", "image")
}


func main() {
    port := 8000
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
