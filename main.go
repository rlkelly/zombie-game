package main

import (
  "bufio"
  "fmt"
  "net"
)

func main() {
  fmt.Println("Launching server...")

  ln, _ := net.Listen("tcp", ":8000")
  conn, _ := ln.Accept()
  r := bufio.NewReader(conn)

  startNewGame(conn, r)
}
