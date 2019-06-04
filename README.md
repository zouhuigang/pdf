### pdf生成和分割

源于分支:

    https://github.com/unidoc/unidoc/commits/master?after=4702ef208700251a5b35bfb882dab5113d52a53a+384
    https://github.com/unidoc/unidoc/tree/7d13be2991fb2791a735fbc9eb1c8abe64ed18f2


    下一个版本:
    https://github.com/unidoc/unidoc/tree/8ae4f6a63a9f7b17bb7a0ea9caeace3ffe45f910




    例子:
    https://github.com/unidoc/unipdf-examples/blob/5e8c12c27a131d0f74140cb9a901f68567e73988/pdf/pdf_rotate.go
    
    下一个版本
    https://github.com/unidoc/unipdf-examples/commits/master?before=ac26b5b4b3a279ead1199bdb152e440eb7474752+70

    https://github.com/unidoc/unipdf-examples/tree/5cf8aa8b497f5d2cc0b7da1a23dca8ae03890401


### 在线版本

    https://foxyutils.com/wordtopdf/
    https://www.ilovepdf.com/zh-cn
    https://github.com/sunreaver/docanalysis
    https://github.com/sajari/docconv


### 其他插件


    https://github.com/zouhuigang/converter

    1、在阿里云CentOS7.4安装calibre碰到了之前在其它版本未遇到的很多问题，安装了如下依赖包，使得ebook-convert可以生成各种文件：
    yum -y install mesa-libGL.x86_64
    yum -y install ImageMagick
    yum install openssl-devel bzip2-devel expat-devel gdbm-devel readline-devel sqlite-devel gcc gcc-c++ openssl-devel
    yum install Xcb
    yum -y install qt5*

    2、虽然生成文件了，但中文都没了，是因为Linux系统少了中文字体
    将Windows系统的C:\Windows\Fonts目录中的中文字体（比如输入微软雅黑）上传到Linux，可在
    /usr/share/fonts/下创建一个目录（比如chinese)保存各种中文字体

