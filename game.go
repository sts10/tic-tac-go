package main

import ("fmt")

func main(){
  gameOver := false
  board := [9]int{0,0,0,0,0,0,0,0,0}
  turnNumber := 1

  for gameOver != true{
    presentBoard(board)
    player := 0
    if turnNumber % 2 == 1{
      fmt.Println("Player 1's turn")
      player = 1
    } else {
      fmt.Println("Player 2's turn")
      player = 2
    }

    
    currentMove := askForPlay()
    // Prevent out-of-bounds moves
    if currentMove > 8 {
      currentMove = askForPlay()
    }

    /*if turnNumber > 1 {
      for v := range board {
        if v == 1 {
          fmt.Println("Illegal Move! Pick an unoccupied space.")
        } else if v == 10 {
          fmt.Println("Illegal Move! Pick an unoccupied space.")
        }
      }  
    }*/

    board = executePlayerMove(currentMove, player, board)
    
    // Press q to quit
    /*if currentMove == 113 {
      gameOver = true
    }*/

    result := checkForWin(board)
    if result > 0 {
      // Adding winner board
      presentBoard(board)

      fmt.Printf("Player %d wins!\n\n", result)
      gameOver = true
    } else {
      turnNumber++
    }

    // Cat's game
    if turnNumber == 10 {
      fmt.Printf("Cat's Game!\n\n")
      gameOver = true
    }

    // Play again
    /*if gameOver == true {
      fmt.Println("Play again? (y/n)")
      var playAgain int
      fmt.Scan(&playAgain)

      if playAgain == 121 {
        
      }
      else if playAgain == 110 {

      }
      else {
        fmt.Println("Type y to play again or n to quit")
      }
    }*/
  }

}

func askForPlay() int{
  fmt.Println("Select a move")
  var moveInt int
  fmt.Scan(&moveInt)
  
  if moveInt > 8 {
    fmt.Println("Illegal Move! Pick a number 0-8.")
    }
    fmt.Println("moveInt is", moveInt)
    return moveInt
}

func executePlayerMove(moveInt int, player int, b [9]int) [9]int {
  if player == 1{
    b[moveInt] = 1
  }else if player == 2{
    b[moveInt] = 10
  }
  return b
}

func presentBoard(b [9]int) {
  for i, v := range b {
    if v == 0{
      // empty space. Display number
      fmt.Printf("%d", i)
    } else if v == 1{
      fmt.Printf("X")
    } else if v == 10{
      fmt.Printf("O")
    }
    // And now the decorations
    if i> 0 && (i+1) % 3 == 0{
      fmt.Printf("\n")
    } else {
      fmt.Printf(" | ")
    }
  }
}

func checkForWin(b [9]int) int {
  // re-calculate sums Array
  sums := [8] int {0,0,0,0,0,0,0,0}
  for _, v := range b[0:2] { sums[7] += v }
  for _, v := range b[3:5] { sums[6] += v }
  for _, v := range b[6:8] { sums[5] += v }

  sums[0] = b[2]+b[4]+b[6]
  sums[1] = b[0]+b[3]+b[6]
  sums[2] = b[1]+b[4]+b[7]
  sums[3] = b[2]+b[5]+b[8]
  sums[4] = b[0]+b[4]+b[8]
  // Adding 3 missing win conditions
  sums[5] = b[6]+b[7]+b[8]
  sums[6] = b[3]+b[4]+b[5]
  sums[7] = b[0]+b[1]+b[2]

  for _, v := range sums {
    if v == 3{
      return 1
    } else if v == 30{
      return 2
    }
  }
  return 0
}
