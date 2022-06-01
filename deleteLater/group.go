package main

type Group struct {
	name  string
	rules []*RuleGroup
}

func NewGroup(name string, rules []*RuleGroup) *Group {
	return &Group{
		name:  name,
		rules: rules,
	}
}

//without goRoutines
//average time taken -> len(g.rules)*100 milliseconds
func (g *Group) evaluate() (answer []string) {
	for i := 0; i < len(g.rules); i++ {
		//answer = append(answer, evaluate(g.rules[i]))
	}
	return
}

//using concurrency
//average time taken -> 100 milliseconds since all of them are running concurrently!
func (g *Group) evaluateConcurrently() (answer []string) {
	c := make(chan string)
	for i := 0; i < len(g.rules); i++ {
		//go func(index int) { c <- evaluate(g.rules[index]) }(i)
	}
	for i := 0; i < len(g.rules); i++ {
		result := <-c
		answer = append(answer, result)
	}
	return
}
