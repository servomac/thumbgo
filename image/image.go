package image

import (

    "gopkg.in/h2non/bimg.v0"
)

type ImageOptions struct {
  Width     int
  Height    int
}

type Image struct {
  Body      []byte
  Mimetype  string
}


func Resize(buf []byte, o ImageOptions) (out Image, err error) {
    imageType := bimg.DetermineImageTypeName(buf)
    if o.Width == 0 && o.Height == 0 {
        return Image{Body: buf, Mimetype: imageType}, nil
    }

    opts := bimg.Options{
      Width:   o.Width,
      Height:  o.Height,
      Crop:    false,
      Quality: 95,
    }

    buf, err = bimg.Resize(buf, opts)
    if err != nil {
        return Image{}, err
    }

    return Image{Body: buf, Mimetype: imageType}, nil
}
