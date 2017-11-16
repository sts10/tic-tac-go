# Notes on writing Tic Tac Go

## Issue #1: 

I had some trouble taking an input from the console and then converting it to an integer. My assumption that teh input would come in as a String was correct. 

```go
func askForPlay() int{
  fmt.Println("Select a move")
  var moveInt int
  fmt.Scan(&moveInt)
  return moveInt
}
```

After a bunch of Googling and false starts, I found that line `fmt.Scan(&moveInt)` which somehow did both things I wanted-- prompt the user for input _while also_ maintaining type `int` for the variable `moveInt`. Awesome-- but, you know, weird how cryptic it is.

The line `fmt.Printf("moveInt is type: %T\n", moveInt)` was a helpful debug step at one point, as it printed the type of the variable `moveInt`. However in my final code, thankfully, I don't think `moveInt` is ever _not_ of type "int".

## Issue #2 

Re-declaring a variable when I really meant to just re-assign it.

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
  currentMove := askForPlay()
  board = executePlayerMove(currentMove, player, board)
  // more code here
}
```

The code above gave me the following error:

```text
# command-line-arguments
./game.go:22:44: undefined: player
```

So I (sloppily) added `player := 0` above the if statement and ran it again. I then got this error:

```text
# command-line-arguments
./game.go:19:17: player declared and not used
./game.go:22:17: player declared and not used
```

The problem here is that the lines in the conditionals that read `player := 1` and `player := 2` _declare_ the variable `player`, as well as re-assign it. What I want to do is simply re-assign the variable. To do that, I changed the block to:

```go
for gameOver != true{
  // some code here
  player := 0
  if turnNumber % 2 == 1{
    fmt.Println("Player 1's turn")
    player = 1
  } else {
    fmt.Println("Player 2's turn")
    player = 2
  }

  currentMove := askForPlay()
  board = executePlayerMove(currentMove, player, board)
  // more code here
}
```

## Smaller Issues

When declaring the `presentBoard` function, I found that you need to not only specify the type of each input and any outputs, but, if one of the inputs or outputs is an array, you also need to tell the function how big the array will be. 

```go
func presentBoard(b [9]int) {
  for i, v := range b {
    if v == 0{
      fmt.Printf("%d", i)
      // more code here
    }
  }
}
```

At first I just wrote `func presentBoard(b []int) { ` figuring that'd be cool but the Go compiler threw me an error: `cannot use board (type [9]int as type []int in argument`. My guess is that `[]int` is actually a Slice rather than an Array.

## Other Notes

### Loops

I found it interesting that Go only has one type of loop: the `for` loop. I ended up using five such loops in my tic-tac-toe game, which I'll informally place into these three categories:

For example, when I wanted something like what is a `while` loop in other languages, I used `for gameOver != true{ /* code block */ }`

The `each` loop is `for _, value := range sums { /* code block */ }`. `range sums` tells Go we want to iterate over all of sums

The `_` is where your index would go. Since the Go compiler throws an error if you declare a variable but don't use it, we need to "kill" the index variable with `_`.

If you do want to use the index (like Ruby `each_with_index`), you'd want `for index, value := range b { /* code block */ }` 

### First Impressions

Given the languages I've played with, Go feels a lot like JavaScript. Go does have structs, but I didn't use any in my game. Rather, my code is organized into different functions, as I would do if writing JavaScript.

For example, functions are called by `funcName(parameterVariable)` as opposed to `parameterVariable.funcName` (like Ruby or a heavily object-oriented language).

Things that were _unlike_ JavaScript or Ruby did throw me a bit. For example, the shortcut syntax for declaring and assigning variables, `:=`, was strange to me. In addition to the issue I discussed above, I kept accidentally typing `=:` or also defining the variables type when I didn't need to because of how the shortcut works.

That said, compiling was nice: I didn't have to create an executable and then run it. Instead, the workflow was very similar to Ruby and other dynamic languages: I simply ran: `go run <filename>.go`. And the errors I got-- which I got a lot of, of course-- were relative clear and easy to understand.


