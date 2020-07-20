## Tasks
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

### Task1
#### Reference
- https://www.cnblogs.com/lfri/p/11769254.html
- https://www.cnblogs.com/famine/p/11973534.html
- https://books.studygolang.com/gopl-zh/ch8/ch8-02.html
#### How to run?
- In one terminal run $go run reverse.go
- In another terminal run $telnet 127.0.0.1 50000

### Task2
#### Reference
- code in Task1
- https://www.ctolib.com/dengsgo-math-engine.html
#### How to run?
- In one terminal run $go run calculator.go
- In another terminal run $telnet 127.0.0.1 50000
