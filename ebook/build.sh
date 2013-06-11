#!/bin/sh

rm -f *.html *~

export GOPATH=`pwd`

#go get -u github.com/russross/blackfriday
go get -u github.com/fairlyblank/md2min

go run build.go

