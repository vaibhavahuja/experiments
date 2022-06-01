package main

import "time"

type Rule struct {
	name string
}

func NewRule(name string) *Rule {
	return &Rule{name: name}
}

//attempt 1 : without Go Routines
func (r *Rule) evaluate() string {
	//say it takes 100 milliseconds for evaluation of rule to complete
	time.Sleep(100 * time.Millisecond)
	return "done -> "
}

func (r *Rule) GetName() string {
	//say it takes 100 milliseconds for evaluation of rule to complete
	time.Sleep(100 * time.Millisecond)
	return "done -> " + r.name
}
