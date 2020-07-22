// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"strings"
)


func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
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
	}
}

func randomNo(randomCh chan int) {
	for {
		n, _ := rand.Int(rand.Reader, big.NewInt(1000))
		randomNumber := int(n.Int64())
		// fmt.Printf("The random number is: %v\n", randomNumber)
		go randHandler(randomCh)
		randomCh <- randomNumber
	}
}

func randHandler(randomCh chan int) {
	x := <-randomCh
	fmt.Printf("The random number is: %v\n", x)
}

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
		}
		inputSting := strings.Trim(string(buf[:len]), "\r\n")
		fmt.Printf("Received data: %v\n", inputSting)
		// randomNo := randHander(randomChannel)
		// randomNumber := randomNo()
		fmt.Fprintf(conn, "Random number = %v\n", <-randomChannel)
	}
}
