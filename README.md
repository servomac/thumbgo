#thumbgo

A small on demand image resizing service written in Go.

## Installation

### Build

0. Install libvips and [bimg](https://github.com/h2non/bimg)

  ```
  curl -s https://raw.githubusercontent.com/lovell/sharp/master/preinstall.sh | sudo bash -
  go get -u gopkg.in/h2non/bimg.v0
  ```

1. Download and install thumbgo

  ```
  go get github.com/servomac/thumbgo
  cd $GOPATH/src/github.com/servomac/thumbgo
  go install
  ```

3. Run it

  ```
  thumbgo -c config.json.sample
  ```

4. Visit [localhost:8000/320x213/static.rappler.com/images/manny-pacquiao-training-20150222.jpg](http://localhost:8000/320x213/static.rappler.com/images/manny-pacquiao-training-20150222.jpg)
