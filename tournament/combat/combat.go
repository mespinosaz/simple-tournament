package combat

import (
	"math/rand"
	"time"
)

type Combat interface {
	Solve() interface{}
}

type combat struct {
	p1 interface{}
	p2 interface{}
}

func NewCombat(p1, p2 interface{}) Combat {
	return &combat{p1: p1, p2: p2}
}

func (c *combat) Solve() interface{} {
	if c.p1 == nil {
		return c.p2
	}

	if c.p2 == nil {
		return c.p1
	}

	return c.DecideWinner()
}

func (c *combat) DecideWinner() interface{} {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(2);

	if n == 0 {
		return c.p1
	}
	return c.p2
}
