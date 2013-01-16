#!/usr/bin/sh

rm -f *.html *~

export GOPATH=`pwd`

go get -u github.com/russross/blackfriday

go run build.go

