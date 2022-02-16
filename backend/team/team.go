package team

import (
	"fmt"
)

type Team struct {
	Name  string
	Power int
}

var teamList = []Team{
	{"Manchester City", 80},
	{"Chelsea", 75},
	{"Liverpool", 75},
	{"Manchester United", 65},
}

func GetTeams() []Team {
	return teamList
}

func Print(team Team) {
	fmt.Printf("Team %v with power %d", team.Name, team.Power)
}
