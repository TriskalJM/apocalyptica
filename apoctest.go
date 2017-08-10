package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"game"
)

func main() {
	files, err := ioutil.ReadDir("./games")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found the following games:")		
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println()

    var supportedGames []*game.Game
	for _, file := range files {
		g := game.NewGame(file.Name(), "./games/" + file.Name() + "/basic_" + strings.ToLower(file.Name()) + ".json")
		supportedGames = append(supportedGames, g)
	}
	
	for _, g := range supportedGames {
		fmt.Println(g.Name)
		fmt.Println()
		
		fmt.Println("Basic Moves:")
		for _, m := range g.BasicMoves {
			fmt.Printf("%s\n", m.Full)
		}
		fmt.Println()
	}
}
