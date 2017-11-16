# Notes on writing Tic Tac Go

One problem: Re-declaring a variable when I really meant to just re-assign it.

```go
for gameOver != true{
  // some other code here
  if turnNumber % 2 == 1{
    fmt.Println("Player 1's turn")
    player := 1
  } else {
    fmt.Println("Player 2's turn")
    player := 2
  }
}
```

I think the code above gave me the following error:

```text
# command-line-arguments
./game.go:19:17: player declared and not used
./game.go:22:17: player declared and not used
```

The problem here is that `player := 1` and `player := 2` _declare_ the variable `player` 
