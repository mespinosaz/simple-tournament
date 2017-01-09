package combat

import (
	"math/rand"
	"time"
)

// Combat of two participants
type Combat interface {
	Solve() string
}

type combat struct {
	p1 string
	p2 string
}

// NewCombat ...
func NewCombat(p1, p2 string) Combat {
	return &combat{p1: p1, p2: p2}
}

func (c *combat) Solve() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(2)

	if n == 0 {
		return c.p1
	}
	return c.p2
}
