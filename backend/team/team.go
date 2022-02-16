package team

import (
	"fmt"
)

type Team struct {
	Name  string
	Power int
}

func GetTeams() []Team {
	TeamList := []Team{
		{"ManchesterCity", 80},
		{"Chelsea", 75},
		{"Liverpool", 70},
		{"Manchester United", 65},
	}

	return TeamList
}

func Print(t Team) {
	fmt.Printf("Team %v with power %d", t.Name, t.Power)
}
