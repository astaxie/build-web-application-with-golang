#!/bin/sh
if [ ! -f build-web-application-with-golang ];then
	go build
fi

if [ ! -d tmp ];then
	mkdir tmp
fi
cd tmp
cp ../*.md .
for i in *.md;do
	sed -i  '/^[#]\{1,\}/s!^\([#]\{1,\}\)\([^#]\{1,\}\)!\1 \2!' $i
	sed -i  '/!\[\](images/s#images\(.*\)?raw=true#../Images\1#' $i

done
cd ../build-web-application-with-golang
rm -rf *.md
