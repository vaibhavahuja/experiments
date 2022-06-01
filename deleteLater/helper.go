package main

func evaluate(object []IRule) (ans []string) {
	for i := 0; i < len(object); i++ {
		myName := object[i].GetName()
		ans = append(ans, myName)
	}
	return ans
}
