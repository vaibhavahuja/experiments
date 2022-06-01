package main

import (
	"fmt"
	"time"
)

type IRule interface {
	GetName() string
}

type RuleGroup struct {
	name        string
	description string
	age         int
	meta        interface{}
}

func (rg *RuleGroup) GetName() string {
	time.Sleep(100 * time.Millisecond)
	return "ruleGroupIsDone -> " + rg.name
}

func main() {
	rule := &Rule{name: "FirstRule"}

	ruleGroup := &RuleGroup{
		name:        "RuleGroupMate",
		description: "this is my description",
		age:         78,
		meta:        nil,
	}
	ruleGroup2 := &RuleGroup{
		name:        "RuleGroupMate",
		description: "this is my description",
		age:         78,
		meta:        nil,
	}
	ruleGroup3 := &RuleGroup{
		name:        "RuleGroupMate",
		description: "this is my description",
		age:         78,
		meta:        nil,
	}
	ruleGroup4 := &RuleGroup{
		name:        "RuleGroupMate",
		description: "this is my description",
		age:         78,
		meta:        nil,
	}

	group := NewGroup("myFirstGroup", []*RuleGroup{ruleGroup, ruleGroup2, ruleGroup3, ruleGroup4})
	fmt.Println(group)
	//start := time.Now()
	//result := group.evaluateConcurrently()
	//elapsed := time.Since(start)
	//fmt.Println(result, "in ", elapsed, " seconds")
	x := []IRule{rule, ruleGroup}
	fmt.Println(evaluate(x))
}
