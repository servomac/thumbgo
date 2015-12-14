#thumbgo

A small on demand image resizing service written in Go.

## Installation

### Build with Go tools

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

### or build the docker image

```
git clone https://github.com/servomac/thumbgo.git
cd thumbgo
docker build -t thumbgo .
docker run -d \
    --name thumbgo \
    --restart always \
    -v `pwd`/config.json.sample:/etc/thumbgo/config.json \
    -p 8000:8000 \
    thumbgo
```
