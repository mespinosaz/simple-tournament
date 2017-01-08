package main

import (
	"github.com/mespinosaz/tournament/tournament"
	"fmt"
	"os"
)

func main() {
	t := tournament.NewTournament(len(os.Args))

	for _,element := range os.Args[1:] {
		t.Add(element)
	}

	fmt.Printf("\nTournament winners is ... %s!!!!\n\n", t.Solve())
}