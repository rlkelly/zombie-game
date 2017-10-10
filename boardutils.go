package main

import (
  "math/rand"
  "time"
)

// Initializes a game board
func makeBoard() (l Board) {
  a := make([][]int, 30)
  for i := range a {
      a[i] = make([]int, 10)
  }
  return Board{a, Location{0, 0}, false}
}

// checks if a value is valid for the board
func between(x int, low int, high int) (a bool) {
	if (x <= high) {
		if (x >= low) {
			return true
		}
	}
	return false
}

// gets all neighbors for current zombie location
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

// moves zombie to a valid board location
func moveZombie(i *Board) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	neighbors := getNeighbors(i)
	n := r1.Intn(len(neighbors))
	i.zombie = neighbors[n]
}
