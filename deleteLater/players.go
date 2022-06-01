package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Player struct {
	Name string
	Age  int
}

//func InitPlayers() string {
//	players := []Player{
//		{Name: "ronaldo", Age: 27},
//		{Name: "vaibhav", Age: 27},
//		{Name: "ahuja", Age: 27},
//		{Name: "messi", Age: 27},
//	}
//	ans, _ := json.Marshal(players)
//	return string(ans)
//}

func InitPlayers() (players []Player) {
	myPlayers, err := ioutil.ReadFile("template/myFile.json")
	if err != nil {
		fmt.Println("error while reading file", err)
		return nil
	}
	//var players []Player
	if err = json.Unmarshal(myPlayers, &players); err != nil {
		fmt.Println("error while unmarshalling players", err)
		return nil
	}
	return
}
