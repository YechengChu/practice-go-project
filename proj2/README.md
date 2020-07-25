# :pencil2: Tasks
1. 建立一个tcp服务器，telnet该服务器，输入一个字符串进去，立马返回该字符串的逆序字符串 
2. 建立一个tcp服务器，telnet该服务器，输入一个表达式，可以计算结果(可以使用外部模块来解析表达式，至少支持加减乘除)
3. 建立一个tcp服务器，telnet该服务器，返回一个真随机算子 (利用gorouine中的知识)
4. 使用cgo封装C++ STL中的Set类型，支持所有主要函数（insert, find, size, erase）
5. 使用colly爬取站点 http://tumregels.github.io/Network-Programming-with-Go/ 所有页面(在本地打开页面后展开的效果和服务器上差不多)
6. 建立一个http服务器，上传一个彩色图片，返回该彩色的黑白图片

参考内容: 
- http://tumregels.github.io/Network-Programming-with-Go/
- https://chai2010.cn/advanced-go-programming-book/
- https://books.studygolang.com/gopl-zh/

## :paperclip: Task1
### Reference
- https://www.cnblogs.com/lfri/p/11769254.html
- https://www.cnblogs.com/famine/p/11973534.html
- https://books.studygolang.com/gopl-zh/ch8/ch8-02.html
### How to run?
- In one terminal run 
  ```
  $ go run reverse.go
  ```
- In another terminal run 
  ```
  $ telnet 127.0.0.1 50000
  ```
  
## :paperclip: Task2
### Reference
- code in Task1
- https://www.ctolib.com/dengsgo-math-engine.html
### How to run?
- In one terminal run
  ```
  $ go run calculator.go
  ```
- In another terminal run
  ```
  $ telnet 127.0.0.1 50000
  ```

## :paperclip: Task3
### Reference
- code in Task1
- https://gobyexample.com/random-numbers
- https://zhuanlan.zhihu.com/p/94684495
- https://yourbasic.org/golang/convert-int-to-string/
### How to run?
- In one terminal run
  ```
  $ go run randomNum.go
  OR
  $ go run randomNum2.go
  ```
- In another terminal run
  ```
  $ telnet 127.0.0.1 50000
  ```

## :paperclip: Task4
### Due to lack of knowledge of CGO programming, this task is skipped!
  
## :paperclip: Task5
### Reference
- http://go-colly.org/docs/examples/basic/
- https://www.jianshu.com/p/cda08dde65cd
- https://github.com/imthaghost/goclone/blob/master/crawler/collector.go
- https://www.jianshu.com/p/f5a6e39deaa5
- https://golang.org/pkg/net/url/#URL
- https://www.jianshu.com/p/c5fcef31707d
- https://www.cnblogs.com/bossliu/p/5043631.html
### How to run?
#### For version 1 & 2
- In the terminal run
  ```
  $ go run crawlerv1.go  
  OR
  $ go run crawlerv2.go  
  ```
- Wait until the program is finished
- Open the folder named 'downloaded' and select folders and htmls to open
#### For version 3
- In the terminal run
  ```
  $ go run crawlerv3.go  
  ```
- Wait until the program is finished
- Make sure you have [Node js](https://nodejs.org/en/download/) installed on your pc
- Also you need to install express under the Node js enviroment
  ```
  $ npm install express
  ```
- Then you can start the server
  ```
  $ node server.js
  ```
- Open a web browser and go to the link below
  ```
  localhost:3000
  ```
## :paperclip: Task6
### Reference
- https://blog.csdn.net/stpeace/article/details/82716145
- https://blog.csdn.net/BlackCarDriver/article/details/87870109
- https://riptutorial.com/go/example/31693/convert-color-image-to-grayscale
- https://blog.csdn.net/mirage003/article/details/88084303
### How to run?
- In terminal run
  ```
  $ go run server.go
  ```
- Open a web browser and go to 
  ```
  localhost:3000/upload
  ```
