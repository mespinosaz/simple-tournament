package main

import (
	"fmt"
	"os"

	"github.com/mespinosaz/tournament/tournament"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Missing argmuent\n")
		return
	}

	ch := make(chan string)

	t := tournament.NewTournament(ch)

	for _, element := range os.Args[1:] {
		t.Add(element)
	}

	go t.Solve()

	w := <-ch

	fmt.Printf("\nTournament winners is ... %s!!!!\n\n", w)
}
