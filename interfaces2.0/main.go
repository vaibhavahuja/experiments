package main

import (
	"fmt"
	"github.com/vaibhavahuja/experiments/interfaces2/entities"
	"time"
)

func main() {
	rule := &entities.Rule{Name: "First"}

	ruleGroup1 := &entities.RuleGroup{
		Name:      "FirstRG",
		Meta:      "none",
		DefinedBy: "me",
	}
	ruleGroup2 := &entities.RuleGroup{
		Name:      "SecondRN",
		Meta:      "none",
		DefinedBy: "me",
	}
	ruleGroup3 := &entities.RuleGroup{
		Name:      "ThirdRG",
		Meta:      "none",
		DefinedBy: "me",
	}
	ruleGroup4 := &entities.RuleGroup{
		Name:      "FourthRG",
		Meta:      "none",
		DefinedBy: "me",
	}

	group := &entities.Group{
		RuleGroupList: []*entities.RuleGroup{ruleGroup1, ruleGroup2, ruleGroup3, ruleGroup4},
		Age:           28,
		Description:   "no description just a check",
	}

	fmt.Println("Time to evaluate one rule")
	start := time.Now()
	//result := EvaluateNormally(rule)
	result := EvaluateConcurrently(rule)
	elapsed := time.Since(start)
	fmt.Println("evaluated rules result ", result, "in time ", elapsed)

	fmt.Println("Time to evaluate one group")
	start = time.Now()
	//result = EvaluateNormally(group)
	result = EvaluateConcurrently(group)
	elapsed = time.Since(start)
	fmt.Println("evaluated group result ", result, "in time ", elapsed)

}
