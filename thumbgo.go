// Thumbgo is an on-demand image resizer
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
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
    nbytes, err := resp.Write(resizedImage.Body)    //improve TODO bufio?
    fmt.Printf("%s (%s) as %dx%d [%d bytes]\n", url, resizedImage.Mimetype, width, height, nbytes)
    if err != nil {
        log.Printf("ERROR while serving %s (%v)\n", url, err)
        http.Error(resp, err.Error(), http.StatusInternalServerError)
        return
    }

    resp.Header().Set("Content-Type", "image/"+resizedImage.Mimetype)
    resp.Header().Set("Content-Length", string(nbytes))
}


type Config struct {
    Addr string
    Port int
}

func (c *Config) ReadConfig(path string) (err error) {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("ERROR with config file: %v\n", err)
    }

    if err = json.Unmarshal(file, c); err != nil {
        log.Fatalf("ERROR parsing config file: %v\n", err)
    }
    return
}


func main() {
    var configFile = flag.String("c", "/etc/thumbgo/config.json", "configuration file")
    flag.Parse()

    var cfg Config
    cfg.ReadConfig(*configFile)
    log.Printf("Starting Thumbgo. Listening on %s:%d", cfg.Addr, cfg.Port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port), nil)
}
