package main

import "time"

func EvaluateConcurrently(listAst ASTInterface) (answer []string) {
	c := make(chan string)
	for i := 0; i < len(listAst.GetListOfNames()); i++ {
		go func(j int) { c <- evalHelper(listAst.GetListOfNames()[j]) }(i)
	}
	for i := 0; i < len(listAst.GetListOfNames()); i++ {
		answer = append(answer, <-c)
	}
	return
}

func EvaluateNormally(listAst ASTInterface) (answer []string) {
	for i := 0; i < len(listAst.GetListOfNames()); i++ {
		answer = append(answer, evalHelper(listAst.GetListOfNames()[i]))
	}
	return
}

func evalHelper(ast string) (answer string) {
	time.Sleep(100 * time.Millisecond)
	answer = "heyy reached here " + ast
	return
}
