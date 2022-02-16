package fixture

import (
	"fmt"
	"interview/backend/match"
)

type Week struct {
	Matches []match.Match
}

func reverseWeek(week Week) Week {
	var reversedWeek Week
	for _, m := range week.Matches {
		var reversedMatch = match.Match{
			HomeTeam:  m.AwayTeam,
			AwayTeam:  m.HomeTeam,
			HomeGoals: 0,
			AwayGoals: 0,
		}

		reversedWeek.Matches = append(reversedWeek.Matches, reversedMatch)
	}

	return reversedWeek
}

func printWeek(week Week) {
	for _, weekMatch := range week.Matches {
		fmt.Printf("%v %d:%d %v\n",
			weekMatch.HomeTeam.Name,
			weekMatch.HomeGoals,
			weekMatch.AwayGoals,
			weekMatch.AwayTeam.Name)
	}
}
