package main

import (
  "bufio"
  "fmt"
  "net"
  "strconv"
  "strings"
  "time"
)

// starts a new game
func startNewGame(conn net.Conn, r *bufio.Reader) {
  var username string
  board := makeBoard()
  started := false

  for (!started) {
    message, _ := r.ReadString('\n')
    data := strings.Split(strings.TrimSuffix(message, "\r\n"), " ")
    if ("START" == data[0]) {
      username = data[1]
      started = true
      fmt.Println("Starting New Game\n")
    } else {
      conn.Write([]byte("Invalid Command\n"))
    }
  }
  ticker := time.NewTicker(1 * time.Second)
  input := make(chan string)
  go checkShot(input, &board, r, conn, username)

  for range ticker.C {
    if !board.won {
      gameLoop(&board, r, conn, username, input)
    } else {
      ticker.Stop()
      conn.Write([]byte("Good Game! Let's Play Again!\n"))
      startNewGame(conn, r)
    }
  }
}

// takes archer's shot and adds it to the channel
func checkShot(input chan string, b *Board, r *bufio.Reader, conn net.Conn, username string) {
  message, _ := r.ReadString('\n')
  input <- message
}

// validates a shot against the zombie's current location
func checkResponse(message string, b *Board, conn net.Conn) (r bool) {
  data := strings.Split(strings.TrimSuffix(message, "\r\n"), " ")
  if len(data) == 3 {
    x, err1 := strconv.Atoi(data[1])
    y, err2 := strconv.Atoi(data[2])
    if (data[0] == "SHOOT" && err1 == nil && err2 == nil) {
      // if x == b.zombie.x && y == b.zombie.y {
      if x == 0 && y == 0 {
          b.won = true
      }
      return true
    }
  }
  conn.Write([]byte("INVALID INPUT\n"))
  return false
}

// loops through updating the zombie's location and checking for user's guess
func gameLoop(b *Board, r *bufio.Reader, conn net.Conn, username string, input chan string) {
    moveZombie(b)
    select {
        case i := <-input:
            checkResponse(i, b, conn)
            if b.won {
              conn.Write([]byte(boomOutput(username)))
            } else {
              go checkShot(input, b, r, conn, username)
            }
        case <-time.After(2 * time.Second):
            conn.Write([]byte(zombieOutput(b)))
    }
}
