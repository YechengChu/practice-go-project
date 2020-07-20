// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
    "fmt"
    "net"
    "strings"
    "math/rand"
    "time"
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

func randomNo() int{
  s1 := rand.NewSource(time.Now().UnixNano())
  // fmt.Printf("Time is: %v\n", time.Now().UnixNano())
  r1 := rand.New(s1)
  randomNumber := r1.Intn(1000)

  // fmt.Printf("The random number has type: %T\n", r1.Intn(1000))
  fmt.Printf("The random number is: %v\n",randomNumber)
  return randomNumber
}

func doServerStuff(conn net.Conn) {
    fmt.Fprintf(conn,"Welcome to the random number generator, getting an int number within 1000!\n")
    for {
        buf := make([]byte, 512)
        len, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return //终止程序
        }
        inputSting := strings.Trim(string(buf[:len]), "\r\n")
        fmt.Printf("Received data: %v\n", inputSting)
        fmt.Fprintf(conn,"Random number = %v\n", randomNo())
    }
}
