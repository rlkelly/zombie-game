package main


type Location struct {
  x int
  y int
}

type Board struct {
  grid [][]int
  zombie Location
  won bool
}
