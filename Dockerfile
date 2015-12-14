FROM debian:jessie
MAINTAINER Toni Pizà <servomac@gmail.com>

ENV VIPS_MAJOR_MINOR 8.1
ENV VIPS_PATCH 1

ENV GOPATH /go
ENV BUILD_PACKAGES="ca-certificates g++ gcc git golang-go make pkg-config wget"
ENV RUN_PACKAGES="libglib2.0-0 libxml2 libvips-dev libgsf-1-dev"

WORKDIR $GOPATH
RUN apt-get update && apt-get install -y --no-install-recommends \
    $BUILD_PACKAGES \
    $RUN_PACKAGES \
 && wget http://www.vips.ecs.soton.ac.uk/supported/${VIPS_MAJOR_MINOR}/vips-${VIPS_MAJOR_MINOR}.${VIPS_PATCH}.tar.gz \
 && tar xvzf vips-${VIPS_MAJOR_MINOR}.${VIPS_PATCH}.tar.gz \
 && cd vips-${VIPS_MAJOR_MINOR}.${VIPS_PATCH} \
 && ./configure --disable-debug --disable-static --disable-introspection --disable-dependency-tracking --enable-cxx=yes --without-python --without-orc --without-fftw \
 && make \
 && make install \
 && ldconfig \
 && echo "Installed libvips $(PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig:/usr/lib/pkgconfig pkg-config --modversion vips)" \
 && cd .. \
 && rm -rf vips-* \
 && go get -u gopkg.in/h2non/bimg.v0 \
 && go get github.com/servomac/thumbgo \
 && cd src/github.com/servomac/thumbgo \
 && go build \
# && apt-get remove --purge -y $BUILD_PACKAGES $(apt-mark showauto) \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR $GOPATH/src/github.com/servomac/thumbgo

EXPOSE 8000

ENTRYPOINT ["./thumbgo"]
