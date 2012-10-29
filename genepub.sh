#!/bin/sh
if [ ! -f build-web-application-with-golang ];then
	go build
fi

if [ ! -d html ];then
	mkdir html 
fi
cd html
cp ../*.md .
for i in *.md;do
	#重新格式化md文件
	sed -i  '/^[#]\{1,\}/s!^\([#]\{1,\}\)\([^#]\{1,\}\)!\1 \2!' $i
	sed -i  '/^[#]\{1,\}/s!  ! !' $i
	#处理md文件中的image src属性
	sed -i  '/!\[\](images/s#images\(.*\)?raw=true#../Images\1#' $i

done
../build-web-application-with-golang >/dev/null
rm *.md
echo "文件已经就绪，请使用sigil导入html目录中的html文件和images目录中的图片文件，制作epub"
