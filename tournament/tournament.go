package tournament

import (
	"github.com/mespinosaz/tournament/stack"
	"github.com/mespinosaz/tournament/tournament/combat"
	"fmt"
)

type Tournament interface {
	Add(m string) int
	Solve() interface{}
}

type tournament struct {
	members stack.Stack
}

func NewTournament() Tournament {
	return &tournament{members: stack.NewStack()}
}

func (t *tournament) Add(m string) int {
	t.members.Push(m)

	return t.members.Size()
}

func (t *tournament) Solve() interface{} {
	if t.members.Size() <= 1 {
		return t.members.Pop()
	}

	var result stack.Stack = stack.NewStack();

	fmt.Print(t.members);

	for t.members.Size() > 0 {
		c := combat.NewCombat(t.members.Pop(), t.members.Pop())
		result.Push(c.Solve())
	}

	t.members = result

	return t.Solve()

}
