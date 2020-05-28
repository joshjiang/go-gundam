#Notes

## Simon Code Review

### 5/20
- [x] use go project layout
- [x] Move err to init
- [x] Move global game assets to struct
- [x] Use math.Abs
- [x] Use pointer receiver method on type player called move, instead of movePlayer
- [x] ~~init should always be at the top~~
- [x] when you move to game board instance/type, remove init
- [ ] check error of *ebiten.Image.Screen.DrawImage
- [x] don't panic

nice to have:

- [ ] package static assets using pkger

## Questions for Simon

- ~~How do I make an interface for Bomb, Player, etc. as Pieces? They reuse some functions like draw, getpos, (java extends, implements)~~
- ~~What do I use an interface for?~~

## Features

 1. Fight Village
  - Touch village switches to new pokemon fight scene 
  - Fight Scene includes 
   - fight, run (raytheon wouldn't like that), negotiate (that's unamerican), etc.
   - gundam and village both have health
   - village has no fight options its a village lmao
  - Exit fight scene