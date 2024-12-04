package main

import "fmt"

type Player struct {
	Name      string
	Countries []Country
	Cards     []string //This is almost certainly temporary
}

type Country struct {
	Name       string
	Neighbor   []*Country
	Armies     int
	Continent  string
	Controller Player
}

func setup() {
	//This might be an INCREDIBLE place to use charm
	fmt.Println("How many players?")
}
func main() {

}
