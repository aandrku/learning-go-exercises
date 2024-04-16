package main

import "fmt"

type TeamName string
type Player string
type Players []Player
type WinsCount int

type Team struct {
	name    TeamName
	players Players
}

type League struct {
	teams []Team
	wins  map[TeamName]WinsCount
}

func (l League) MatchResult(t1 TeamName, score1 int, t2 TeamName, score2 int) {
	if score1 > score2 {
		l.wins[t1]++
	} else if score1 < score2 {
		l.wins[t2]++
	}

	swaps := 1

	for swaps != 0 {
		swaps = 0
		for i := 0; i < len(l.teams)-1; i++ {
			if l.wins[l.teams[i].name] < l.wins[l.teams[i+1].name] {
				tmp := l.teams[i]
				l.teams[i] = l.teams[i+1]
				l.teams[i+1] = tmp
				swaps++
			}
		}
	}
}

func (l League) Ranking() []TeamName {
	var teams []TeamName
	for _, v := range l.teams {
		teams = append(teams, v.name)
	}

	return teams
}

func main() {
	players1 := []Player{
		"p1", "p2", "p3", "p4",
		"p5", "p6", "p7", "p8",
	}

	players2 := []Player{
		"p1", "p2", "p3", "p4",
		"p5", "p6", "p7", "p8",
	}
		
	players3 := []Player{
		"p1", "p2", "p3", "p4",
		"p5", "p6", "p7", "p8",
	}

	team1 := Team{
		name: "Team1",
		players: players1,
	}

	team2 := Team{
		name: "Team2",
		players: players2,
	}
	
	team3 := Team{
		name: "Team3",
		players: players3,
	}

	var league League
	league.teams = []Team{team1, team2, team3}
	league.wins = map[TeamName]WinsCount{}

	league.MatchResult(team1.name, 3, team2.name, 8)
	league.MatchResult(team2.name, 3, team3.name, 8)
	league.MatchResult(team2.name, 3, team1.name, 8)
	league.MatchResult(team2.name, 3, team2.name, 8)
	league.MatchResult(team3.name, 3, team1.name, 8)
	league.MatchResult(team1.name, 3, team2.name, 8)
	league.MatchResult(team2.name, 3, team3.name, 8)
	league.MatchResult(team3.name, 3, team2.name, 8)

	fmt.Println(league.wins)

	fmt.Println(league.Ranking())
		
}
