#!/bin/bash

SED='sed'

if [ `uname -s` == 'Darwin' ] ; then
  SED='gsed'
fi

bn="`basename $0`"
WORKDIR="$(cd $(dirname $0); pwd -P)"

#
# Default language: zh
# You can overwrite following variables in config file.
#
MSG_INSTALL_PANDOC_FIRST='请先安装pandoc，然后再次运行'
MSG_SUCCESSFULLY_GENERATED='build-web-application-with-golang.epub 已经建立'
MSG_CREATOR='M2shad0w'
MSG_DESCRIPTION='一本开源的Go Web编程书籍'
MSG_LANGUAGE='zh-CN'
MSG_TITLE='Go Web编程'
[ -e "$WORKDIR/config" ] && . "$WORKDIR/config"


TMP=`mktemp -d 2>/dev/null || mktemp -d -t "${bn}"` || exit 1
# TMP=./build
trap 'rm -rf "$TMP"' 0 1 2 3 15


cd "$TMP"

(
# [ go list github.com/a8m/mark >/dev/null 2>&1 ] || export GOPATH="$PWD"
# go get -u github.com/a8m/mark
WORKDIR="$WORKDIR" TMP="$TMP" go run "$WORKDIR/build_new.go"
)

if [ ! type -P pandoc >/dev/null 2>&1 ]; then
	echo "$MSG_INSTALL_PANDOC_FIRST"
	exit 0
fi

cat <<__METADATA__ > metadata.txt
<dc:creator>$MSG_CREATOR</dc:creator>
<dc:description>$MSG_DESCRIPTION</dc:description>
<dc:language>$MSG_LANGUAGE</dc:language>
<dc:rights>Creative Commons</dc:rights>
<dc:title>$MSG_TITLE</dc:title>
__METADATA__

mkdir -p $TMP/images
cp -r $WORKDIR/images/* $TMP/images/
ls [0-9]*.html | xargs $SED -i "s/png?raw=true/png/g"

echo "工作目录$WORKDIR, 临时目录$TMP"

pandoc --reference-links -S --toc -f html -t epub --epub-metadata=metadata.txt --epub-cover-image="$WORKDIR/images/cover.png" -o "$WORKDIR/build-web-application-with-golang.epub" `ls [0-9]*.html | sort`

echo "$MSG_SUCCESSFULLY_GENERATED"
