## 随机算子
### 目标
建立一个tcp服务器，telnet该服务器，返回一个真随机算子 (利用gorouine中的知识)
### 练习
randomNum.go
```go
// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	} // if
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		} else {
			fmt.Println("Someone is connected!")
		}
		go doServerStuff(conn)
	} // for
} // main

func randomNo() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	// fmt.Println(time.Now().UnixNano())
	// fmt.Printf("Time is: %v\n", time.Now().UnixNano())
	r1 := rand.New(s1)
	randomNumber := r1.Intn(1000)

	// fmt.Printf("The random number has type: %T\n", r1.Intn(1000))
	fmt.Printf("The random number is: %v\n", randomNumber)
	return randomNumber
} // randomNo

func doServerStuff(conn net.Conn) {
	fmt.Fprintf(conn, "Welcome to the random number generator, getting an int number within 1000!\n")
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		} // if
		inputSting := strings.Trim(string(buf[:len]), "\r\n")
		fmt.Printf("Received data: %v\n", inputSting)
		fmt.Fprintf(conn, "Random number = %v\n", randomNo())
	} // for
} // doServerStuff
```
运行
![Screen Shot 2020-07-22 at 23.15.39.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595431947223-7ca54ec9-28ca-4efa-85f8-55dff49d51c1.png#align=left&display=inline&height=900&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-22%20at%2023.15.39.png&originHeight=900&originWidth=1440&size=134393&status=done&style=none&width=1440)
### 改进

- 改进1: 上一个版本使用的math/random并不能得到一个真正的随机数，虽然使用时间作为seed能生成一个比较随机的数字，这次使用crypto/rand来生成随机数
- 改进2: 上一个版本在运用随机数时，没有使用到goroutine这一版使用了goroutine和channel，具体使用如下
   - 在doServerStuff中增加int类型的randomChannel用于接收随机数
   - 使用go randomNo(randomChannel)，并发执行randomNo来获取随机数
   - randomNo使用crypto/rand不断生成随机数并将其放入randomChannel中
   - 使用go randHandler(randomCh)来提取出生成的随机数，否则channel会被阻塞，后续随机数不会被生成
   - 用time.Sleep(10 * 1e9)暂停10秒再进行提取，这样在doServerStuff的fmt.Fprintf(conn, "Random number = %v\n", <-randomChannel)中就可以将randomChannel中的值提取出来
   - 暂停10s也可以减慢随机数生成速度(因为接收者不接收，发送方就会一直阻塞)，避免程序卡顿
```go
// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)


func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	} // if
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		} else {
			fmt.Println("Someone is connected!")
		}
		go doServerStuff(conn)
	} // for
} // main

func randomNo(randomCh chan int) {
	for {
		// use crypto/rand to generate a true random number
		n, _ := rand.Int(rand.Reader, big.NewInt(1000))
		randomNumber := int(n.Int64())
		// fmt.Printf("The random number is: %v\n", randomNumber)
		go randHandler(randomCh)
		randomCh <- randomNumber
	} // for
} // randomNo

func randHandler(randomCh chan int) {
	// wait for 10 seconds before receive so that the doServerStuff can get the
	// random number generated
	time.Sleep(10 * 1e9) // sleep for 10 seconds
	x := <-randomCh
	fmt.Printf("The random number is: %v\n", x)
} // randHandler

func doServerStuff(conn net.Conn) {
	randomChannel := make(chan int)
	go randomNo(randomChannel)
	// go randomNo()
	fmt.Fprintf(conn, "Welcome to the random number generator, getting an int number within 1000!\n")
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		} // if
		inputSting := strings.Trim(string(buf[:len]), "\r\n")
		fmt.Printf("Received data: %v\n", inputSting)
		// randomNo := randHander(randomChannel)
		// randomNumber := randomNo()
		fmt.Fprintf(conn, "Random number = %v\n", <-randomChannel)
	} // for
} // doServerStuff
```
运行
![Screen Shot 2020-07-22 at 23.10.50.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595432569122-5b3945aa-e46a-4e41-86d0-ade0ba037f10.png#align=left&display=inline&height=900&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-22%20at%2023.10.50.png&originHeight=900&originWidth=1440&size=139775&status=done&style=none&width=1440)

---

## 参考资料
Rogn
[https://www.cnblogs.com/lfri/p/11769254.html](https://www.cnblogs.com/lfri/p/11769254.html)
8.2. 示例: 并发的Clock服务
[https://books.studygolang.com/gopl-zh/ch8/ch8-02.html](https://books.studygolang.com/gopl-zh/ch8/ch8-02.html)
Go by Example: Random Numbers
[https://gobyexample.com/random-numbers](https://gobyexample.com/random-numbers)
Golang随机数生成
[https://www.cnblogs.com/jukaiit/p/10785433.html](https://www.cnblogs.com/jukaiit/p/10785433.html)
golang 生成随机数或者字符
[https://zhuanlan.zhihu.com/p/94684495](https://zhuanlan.zhihu.com/p/94684495)
Convert between int, int64 and string
[https://yourbasic.org/golang/convert-int-to-string/](https://yourbasic.org/golang/convert-int-to-string/)


