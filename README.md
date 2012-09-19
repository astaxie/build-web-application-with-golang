#《Go Web 编程》
因为自己对Web开发比较感兴趣，所以最近抽空在写一本开源的书籍《Go Web编程》《Build Web Application with Golang》。写这本书不表示我能力很强，而是我愿意分享，和大家一起分享Go写Web应用的一些东西。

- 对于从php/python/ruby转过来的同学了解go怎么写web应用开发的

- 对于从c/c++转过来的同学了解web到底是怎么运行起来的

我一直认为知识是用来分享的，让更多的人分享自己拥有的一切知识这个才是人生最大的快乐。

这本书目前我放在github上，我现在基本每天晚上抽空会写一些，时间有限、能力有限，所以希望更多的朋友参与到这个开源项目中来。

## 撰写方法
### 文件命名
每个章节建立一个md文件，如第11章的第3节，则建立**11.3.md**。

## 格式规范
请参看已有章节的规范，要注意的是，每个章节在底部都需要有一个links节，包含“目录”，“上一节”和“下一节”的链接。

##如何编译
`build.go`依赖markdown的一个解析包，所以第一步先

	go get github.com/russross/blackfriday

这样读者就可以把相应的Markdown文件编译成html文件，执行`go build build.go`，执行生成的文件，就会在底目录下生成相应的html文件

##如何编译
目前可以把相应的Markdown编译成html文件，执行`go build build.go`，执行生成的文件，就会在底目录下生成相应的html文件。

##致谢
首先要感谢golang的QQ群102319854，里面的每一个人都很热心，同时要特别感谢几个人

 - [四月份平民](https://plus.google.com/110445767383269817959) (review代码)
 - [Hong Ruiqi](https://github.com/hongruiqi) (review代码)
 - [BianJiang](https://github.com/border) (编写go开发工具Vim和Emacs的设置)
 - [Oling Cat](https://github.com/OlingCat)(review代码)
 - [Wenlei Wu](mailto:spadesacn@gmail.com)(提供一些图片展示)

##授权许可
除特别声明外，本书中的内容使用[CC BY-SA 3.0 License](http://creativecommons.org/licenses/by-sa/3.0/)（创作共用 署名-相同方式共享3.0许可协议）授权，代码遵循[BSD 3-Clause License](<LICENSE.md>)（3项条款的BSD许可协议）。

##开始阅读
[开始阅读](https://github.com/astaxie/build-web-application-with-golang/blob/master/preface.md)
