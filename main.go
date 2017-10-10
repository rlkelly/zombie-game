package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "net"
  "strconv"
  "strings"
  "time"
)

type Location struct {
  x int
  y int
}

type Board struct {
  grid [][]int
  zombie Location
  won bool
}

func main() {
  fmt.Println("Launching server...")

  ln, _ := net.Listen("tcp", ":8000")
  conn, _ := ln.Accept()
  r := bufio.NewReader(conn)

  startNewGame(conn, r)
}

func startNewGame(conn net.Conn, r *bufio.Reader) {
  var username string
  board := makeBoard()
  started := false

  for (!started) {
    message, _ := r.ReadString('\n')
    if ("START" == strings.Split(message, " ")[0]) {
      username = strings.Split(message, " ")[1]
      started = true
      fmt.Println("STARTING NEW GAME\n")
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
      conn.Write([]byte("GOOD GAME!\n"))
      startNewGame(conn, r)
    }
  }
}

func checkShot(input chan string, b *Board, r *bufio.Reader, conn net.Conn, username string) {
  message, _ := r.ReadString('\n')
  input <- message
  go checkShot(input, b, r, conn, username)
}

func checkResponse(message string, b *Board, conn net.Conn) (r bool) {
  data := strings.Split(strings.TrimSuffix(message, "\n"), " ")
  if len(data) == 3 {
    x, err1 := strconv.Atoi(data[1])
    y, err2 := strconv.Atoi(data[2])
    if (data[0] == "SHOOT" && err1 == nil && err2 == nil) {
      if x == b.zombie.x && y == b.zombie.y {
          b.won = true
      }
      return true
    }
  }
  conn.Write([]byte("INVALID INPUT\n"))
  return false
}

func gameLoop(b *Board, r *bufio.Reader, conn net.Conn, username string, input chan string) {
    moveZombie(b)
    select {
        case i := <-input:
            checkResponse(i, b, conn)
            if b.won {
              conn.Write([]byte(boomOutput(username)))
            }
        case <-time.After(2000 * time.Millisecond):
            conn.Write([]byte(zombieOutput(b)))
    }
}

func boomOutput(username string) (response string) {
  return fmt.Sprintf("BOOM %s\n", username)
}

func zombieOutput(b *Board) (response string) {
  return fmt.Sprintf("WALK night-king %d %d\n", b.zombie.x, b.zombie.y)
}

func makeBoard() (l Board) {
  a := make([][]int, 30)
  for i := range a {
      a[i] = make([]int, 10)
  }
  return Board{a, Location{0, 0}, false}
}

func between(x int, low int, high int) (a bool) {
	if (x <= high) {
		if (x >= low) {
			return true
		}
	}
	return false
}

func getNeighbors(b *Board) (l []Location) {
	x := b.zombie.x
	y := b.zombie.y
	l = make([]Location, 0, 8)

	for row := x-2; row <= x+2; row += 2 {
    for col := y-2;  col <= y+2; col += 2 {
    	if (between(row, 0, 10) && between(col, 0, 30)) {
    	   if x != row && y != col {
           l = append(l, Location{row, col})
    		 }
    	}
    }
  }
	return l
}

func moveZombie(i *Board) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	neighbors := getNeighbors(i)
	n := r1.Intn(len(neighbors))
	i.zombie = neighbors[n]
}
