#/bin/bash

VERSION=$1
if [ "$VERSION" == "" ]; then
  echo "usage: $0 <version>"
  exit 1
fi

glide install
for GOOS in darwin linux; do
  for GOARCH in 386 amd64; do
    echo ">> Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -o bin/groot-$VERSION-$GOOS-$GOARCH
  done
done
