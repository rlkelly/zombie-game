# zombie-game


This is a basic zombie game implementation in Golang.
To start a game run the go server, by building the package or running `go run *.go` which you can access via telnet.

From there, the command's are:
  1) `START <NAME>` to start a new game.
  2) `SHOOT <x> <y>` to shoot at the zombie

The zombie's location will be updated every two seconds.
Upon killing the zombie, you can start a new game by typing `START <NAME>` again.
