package team

import (
	"fmt"
)

type Team struct {
	Name  string `json:"name"`
	Power int    `json:"power"`
}

func GetTeams() []Team {
	// TeamList := []Team{
	// 	{"ManchesterCity", 80},
	// 	{"Chelsea", 75},
	// 	{"Liverpool", 70},
	// 	{"Manchester United", 65},
	// }

	TeamList := []Team{
		{"Manchester City", 80},
		{"Salzburg", 35},
		{"Lyon", 50},
		{"FC Barcelona", 70},
	}

	return TeamList
}

func Print(t Team) {
	fmt.Printf("Team %v with power %d", t.Name, t.Power)
}
