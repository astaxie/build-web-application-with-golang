#!/bin/bash
if ! which pandoc >/dev/null ;then
	echo "请先安装pandoc，然后再次运行"
	exit 0
fi

sed -i 's!https://github.com/astaxie/build-web-application-with-golang/blob/master/!!g' README.md
for i in *.md;do
	#重新格式化md文件
	sed -i  '/^[#]\{1,\}/s!^\([#]\{1,\}\)\([^#]\{1,\}\)!\1 \2!' $i #以#开头的行，在#后增加空格
	sed -i  '/^[#]\{1,\}/s!  ! !' $i  #以#开头的行, 删除多余的空格
	#sed -i  '/!\[\](images/s#images\(.*\)?raw=true#../Images\1#' $i
	#sed -i  '/!\[\](images/s#images\(.*\)?raw=true#../images\1#' $i #处理md文件中的image src属性
	sed -i '/[#]\{2,\} links/,/[ ]\{0,\}Id\$.*/d' $i #删除页面链接
done
list="`ls [0-9]*.html |sort `"
cat > metadata.txt <<EOF
<dc:creator>Astaxie</dc:creator>
<dc:description>一本开源的Go Web编程书籍</dc:description>
<dc:language>zh-CN</dc:language>
<dc:rights>Creative Commons</dc:rights>
<dc:title>Go Web编程</dc:title>
EOF

pandoc --reference-links -S --toc -f html -t epub --epub-metadata=metadata.txt --epub-cover-image=../images/cover.png \
-o ../build-web-application-with-golang.epub $list

rm -rf html
echo "build-web-application-with-golang.epub 已经建立"
