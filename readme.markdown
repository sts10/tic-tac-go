# Tic Tac Go

I'm writing a command line tic-tac-toe game as I try to learn Golang!

It's not going _great_... right now it's not compiling! This is basically the compile error I'm getting as of writing this:

```text
# command-line-arguments
./game.go:19:17: player declared and not used
./game.go:22:17: player declared and not used
```

## Some Notes on My Plan

To check for a winner, I'm using a second array called `sums` that adds up each of the possible wins in the game of tic-tac-toe. (Fun fact: I used this idea (and drew the sketch below) back in 2013 as part of my admission test to The Flatiron School.)

![sums explained](img/map.jpg)

## Where I've Been Learning Go From

[This great YouTube video](https://www.youtube.com/watch?v=CF9S4QZuV30&feature=youtu.be) as well as ["A Tour of Go"](https://tour.golang.org).

If you have other resources-- especially for those coming from dynamic type languages like Ruby (oh, sweet, sweet Ruby)-- that you think would help me (or others), feel free to file an issue or [hit me up on Mastodon](https://octodon.social/@schlink)!
