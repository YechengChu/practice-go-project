// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
    "fmt"
    "net"
    "strings"
    "github.com/dengsgo/math-engine/engine"
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

func calculate(s string) float64{
  // call top level function
  r, err := engine.ParseAndExec(s)
  if err != nil {
    fmt.Println(err)
  }
	// fmt.Printf("Type of r is: %T\n",r)
  fmt.Printf("%s = %v\n", s, r)
  return r
}

func doServerStuff(conn net.Conn) {
    for {
        buf := make([]byte, 512)
        len, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return //终止程序
        }
        inputSting := strings.Trim(string(buf[:len]), "\r\n")
        fmt.Printf("Received data: %v\n", inputSting)
        fmt.Fprintf(conn,"%s = %v\n", inputSting,calculate(inputSting))
    }
}
