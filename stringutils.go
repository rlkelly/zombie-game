package main

import (
  "fmt"
)

func boomOutput(username string) (response string) {
  return fmt.Sprintf("BOOM %s\n", username)
}

func zombieOutput(b *Board) (response string) {
  return fmt.Sprintf("WALK night-king %d %d\n", b.zombie.x, b.zombie.y)
}
