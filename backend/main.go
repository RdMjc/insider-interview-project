package main

import (
	"fmt"
	"interview/backend/match"
	"interview/backend/team"
)

func main() {
	// team := teams.Team{"ManchesterCity", 80}

	teams := team.GetTeams()
	fmt.Println(teams)

	onematch := match.Match{
		teams[0],
		teams[1],
		0,
		0,
	}

	// fmt.Println(match)
	match.Print(onematch)
}
