// Package loader provides different Thumbgo loaders
package loader

import (
    "io/ioutil"
    "net/http"
)

func HttpLoader(url string) ([]byte, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    buf, err := ioutil.ReadAll(res.Body)
    return buf, err
}

