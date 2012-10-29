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
	#���¸�ʽ��md�ļ�
	sed -i  '/^[#]\{1,\}/s!^\([#]\{1,\}\)\([^#]\{1,\}\)!\1 \2!' $i
	sed -i  '/^[#]\{1,\}/s!  ! !' $i
	#����md�ļ��е�image src����
	sed -i  '/!\[\](images/s#images\(.*\)?raw=true#../Images\1#' $i

done
../build-web-application-with-golang >/dev/null
rm *.md
echo "�ļ��Ѿ���������ʹ��sigil����htmlĿ¼�е�html�ļ���imagesĿ¼�е�ͼƬ�ļ�������epub"
