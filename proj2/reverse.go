// code adapted from https://www.cnblogs.com/lfri/p/11769254.html
package main

import (
    "fmt"
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

// code adapted from https://www.cnblogs.com/famine/p/11973534.html
func reverse(s string) string{
  o := []rune(s)
  for i,j := 0, len(o)-1 ; i < j ; i, j = i+1, j-1 {
    o[i],o[j] = o[j], o[i]
  }
  return string(o)
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
        fmt.Fprintf(conn,"Reversed input is: %v\n", reverse(inputSting))
    }
}
