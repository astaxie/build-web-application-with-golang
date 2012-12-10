# 《Go Web 编程》
因为自己对Web开发比较感兴趣，所以最近抽空在写一本开源的书籍《Go Web编程》《Build Web Application with Golang》。写这本书不表示我能力很强，而是我愿意分享，和大家一起分享Go写Web应用的一些东西。

- 对于从PHP/Python/Ruby转过来的同学了解Go怎么写Web应用开发的

- 对于从C/C++转过来的同学了解Web到底是怎么运行起来的

我一直认为知识是用来分享的，让更多的人分享自己拥有的一切知识这个才是人生最大的快乐。

这本书目前我放在Github上，我现在基本每天晚上抽空会写一些，时间有限、能力有限，所以希望更多的朋友参与到这个开源项目中来。

**参加了51CTO博客大赛，希望你能够投出宝贵的一票：[http://blog.51cto.com/contest2012/6177767](http://blog.51cto.com/contest2012/6177767)**

## 撰写方法
### 文件命名
每个章节建立一个md文件，如第11章的第3节，则建立**11.3.md**。
### 代码文件
代码文件置于src目录之下。每小节代码按目录存放。如第11章的第3节的代码保存于**src/11.3/**目录下。在正文中按需要添加代码。

## 格式规范
### 正文
请参看已有章节的规范，要注意的是，每个章节在底部都需要有一个links节，包含“目录”，“上一节”和“下一节”的链接。
### 代码
代码要**`go fmt`**后提交。注释文件注明其所属章节。

## 如何编译
`build.go`依赖markdown的一个解析包，所以第一步先

	go get github.com/russross/blackfriday

这样读者就可以把相应的Markdown文件编译成html文件，执行`go build build.go`，执行生成的文件，就会在底目录下生成相应的html文件

## 交流
欢迎大家加入QQ群：259316004 《Go Web编程》专用交流群

大家有问题还可以上德问上一起交流学习：http://www.dewen.org/topic/165

## 致谢
首先要感谢Golang-China的QQ群102319854，里面的每一个人都很热心，同时要特别感谢几个人

 - [四月份平民](https://plus.google.com/110445767383269817959) (review代码)
 - [Hong Ruiqi](https://github.com/hongruiqi) (review代码)
 - [BianJiang](https://github.com/border) (编写go开发工具Vim和Emacs的设置)
 - [Oling Cat](https://github.com/OlingCat)(review代码)
 - [Wenlei Wu](mailto:spadesacn@gmail.com)(提供一些图片展示)

## 授权许可
除特别声明外，本书中的内容使用[CC BY-SA 3.0 License](http://creativecommons.org/licenses/by-sa/3.0/)（创作共用 署名-相同方式共享3.0许可协议）授权，代码遵循[BSD 3-Clause License](<https://github.com/astaxie/build-web-application-with-golang/blob/master/LICENSE.md>)（3项条款的BSD许可协议）。

## 开始阅读
[开始阅读](<https://github.com/astaxie/build-web-application-with-golang/blob/master/preface.md>)


[![githalytics.com alpha](https://cruel-carlota.pagodabox.com/44c98c9d398b8319b6e87edcd3e34144 "githalytics.com")](http://githalytics.com/astaxie/build-web-application-with-golang)
