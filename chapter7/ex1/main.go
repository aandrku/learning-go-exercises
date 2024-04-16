package main

type TeamName string
type Player string
type Players []Player
type WinsCount int


type Team struct {
	name TeamName
	players Players
}

type League struct {
	teams []Team
	wins map[TeamName]WinsCount 
}

func main() {

}
