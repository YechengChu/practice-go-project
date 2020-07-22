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
