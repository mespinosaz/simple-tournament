package main

import (
	"github.com/mespinosaz/tournament/tournament"
	"fmt"
	"os"
)

func main() {
	t := tournament.NewTournament()

	for _,element := range os.Args[1:] {
		t.Add(element)
	}

	fmt.Print(t.Solve())
}