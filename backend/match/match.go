package match

import (
	"fmt"
	"interview/backend/team"
)

type Match struct {
	HomeTeam  team.Team
	AwayTeam  team.Team
	HomeGoals int
	AwayGoals int
}

func Print(match Match) {
	fmt.Printf("%v %d-%d %v \n", match.HomeTeam.Name, match.HomeGoals, match.AwayGoals, match.AwayTeam.Name)
}
