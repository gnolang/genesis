package main

import (
	"fmt"
	"math/big"
)

type Node struct {
	Name     string
	Children []*Edge
	Parents  []*Edge
}

type Edge struct {
	From   *Node
	To     *Node
	Power  *big.Rat // max 1/2 delegation
	Frozen bool
}

func (n *Node) Delegate(to *Node, power *big.Rat) {
	// check capacity
	sum := new(big.Rat).Set(power)
	for _, child := range n.Children {
		sum.Add(sum, child.Power)
	}
	if sum.Cmp(big.NewRat(1, 2)) > 0 {
		panic("cannot delegate more than 1/2 of total power")
	}

	edge := &Edge{
		From:   n,
		To:     to,
		Power:  power,
		Frozen: false,
	}
	n.Children = append(n.Children, edge)
	to.Parents = append(to.Parents, edge)
}

func (n *Node) Power() *big.Rat {
	if len(n.Parents) == 0 {
		return big.NewRat(1, 1) // root node has full power
	}

	power := big.NewRat(0, 1)
	for _, parent := range n.Parents {
		parentPower := new(big.Rat).Mul(parent.Power, parent.From.Power())
		power.Add(power, parentPower)
	}
	return power
}

func (n *Node) String() string {
	return n.stringRec("")
}

func (e *Edge) stringRec(prefix string) string {
	output := fmt.Sprintf("%s|-%s- %s (Aggregated Power: %s)\n", prefix, e.Power.FloatString(2), e.To.Name, e.To.Power())
	output += e.To.stringRec(prefix + "    ")
	return output
}

func (n *Node) stringRec(prefix string) string {
	output := ""
	if n.Parents == nil {
		output += fmt.Sprintf("%s*%s*\n", prefix, n.Name)
	}
	for _, child := range n.Children {
		output += child.stringRec(prefix + "    ")
	}
	return output
}

func main() {
	var (
		root   = &Node{Name: "root"}
		alice  = &Node{Name: "alice"}
		bob    = &Node{Name: "bob"}
		charly = &Node{Name: "charly"}
		dave   = &Node{Name: "Dave"}
		erin   = &Node{Name: "Erin"}
		frank  = &Node{Name: "Frank"}
		grace  = &Node{Name: "Grace"}
		heidi  = &Node{Name: "Heidi"}
		ivan   = &Node{Name: "Ivan"}
	)

	// root & top level members
	root.Delegate(alice, big.NewRat(1, 1000))
	root.Delegate(bob, big.NewRat(1, 1000))
	for _, topLevel := range root.Children {
		topLevel.Power = big.NewRat(1, 1)
	}

	// DAO members
	alice.Delegate(charly, big.NewRat(1, 2))
	bob.Delegate(frank, big.NewRat(1, 2))
	charly.Delegate(dave, big.NewRat(1, 4))
	charly.Delegate(erin, big.NewRat(1, 4))
	dave.Delegate(grace, big.NewRat(1, 4))
	dave.Delegate(heidi, big.NewRat(1, 4))
	grace.Delegate(heidi, big.NewRat(1, 2))
	heidi.Delegate(ivan, big.NewRat(1, 2))

	fmt.Println(root)
}
