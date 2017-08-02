package game

import (
	"fmt"
    "moves"
)

type Playbook struct {
    Moves moves.Move
}

type Game struct {
    Name string
    BasicMoves moves.Move
    BasicMoveList []string
    Playbooks []Playbook
}

func (g *Game) MakeList() []string {
   l := make([]string, 0)
   for _, c := range g.BasicMoves {
           l = append(l, fmt.Sprintf("**%s** (%s)", c.Full, c.Name))
   }
   return l
}

func NewGame(name, basicMovesFilename string) *Game {
    g := new(Game)
	g.Name = name
    g.BasicMoves = moves.LoadMoves(basicMovesFilename)
    g.BasicMoveList = g.MakeList()
    return g
}
