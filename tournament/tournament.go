package tournament

import (
	"github.com/mespinosaz/tournament/stack"
	"github.com/mespinosaz/tournament/tournament/combat"
	"fmt"
	"runtime"
	"sync"
	"math"
)

type Tournament interface {
	Add(m string) int
	Solve() string
}

type tournament struct {
	members stack.StringStack
	round int
}

func NewTournament(n int) Tournament {
	runtime.GOMAXPROCS(n)

	return &tournament{members: stack.NewStack(), round: 1}
}

func (t *tournament) Add(m string) int {
	t.members.Push(m)

	return t.members.Size()
}

func (t *tournament) Solve() string {
	n := t.members.Size()

	if n <= 1 {
		w,_ := t.members.Pop()
		return w
	}

	var wg sync.WaitGroup

	wg.Add(int(math.Ceil(float64(n)/2)))

	var result stack.StringStack = stack.NewStack();

	fmt.Printf("ROUND %d\n", t.round)
	fmt.Print("==========\n")

	for n > 0 {
		go func() {
			defer wg.Done()
			result.Push(t.ProcessMembersPair())
		}()
		n -= 2
	}

	wg.Wait()

	t.members = result

	t.round += 1

	fmt.Print("==========\n")

	return t.Solve()

}

func (t *tournament) ProcessMembersPair() string {
	m1,err1 := t.members.Pop()
	m2,err2 := t.members.Pop()

	if err1 != nil {
		return m2
	}

	if err2 != nil {
		return m1
	}

	c := combat.NewCombat(m1, m2)
	w := c.Solve()

	fmt.Printf("%s VS %s. Winner is %s\n", m1, m2, w)

	return w
}
