package main

import (
	"github.com/mespinosaz/tournament/tournament"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Missing argmuent\n")
		return
	}

	t := tournament.NewTournament(len(os.Args))

	for _,element := range os.Args[1:] {
		t.Add(element)
	}

	fmt.Printf("\nTournament winners is ... %s!!!!\n\n", t.Solve())
}