package league

import (
	"fmt"
	"interview/backend/fixture"
	"interview/backend/team"
)

type League struct {
	fix        fixture.Fixture
	activeWeek int
}

type TeamStatistics struct {
	games          int
	points         int
	win            int
	draw           int
	loss           int
	goalDifference int
}

func GetLeagueTable(l League) map[team.Team]TeamStatistics {
	var leagueTable = make(map[team.Team]TeamStatistics)

	var defaultStatistics = TeamStatistics{
		games:          0,
		points:         0,
		win:            0,
		draw:           0,
		loss:           0,
		goalDifference: 0,
	}

	fmt.Println(leagueTable)

	for _, t := range team.GetTeams() {
		leagueTable[t] = defaultStatistics
	}

	// go through all weeks played in the fixture
	// make team statistics
	for i := 0; i < l.activeWeek; i++ {
		// go through matches in that week
		for _, m := range l.fix.Weeks[i].Matches {

			// Home Team
			if entry, ok := leagueTable[m.HomeTeam]; ok {
				// games statistic
				entry.games += 1
				// win - draw - loss statistic
				if m.HomeGoals > m.AwayGoals {
					entry.win += 1
					entry.points += 3
				} else if m.HomeGoals == m.AwayGoals {
					entry.draw += 1
					entry.points += 1
				} else {
					entry.loss += 1
				}
				entry.goalDifference += m.HomeGoals - m.AwayGoals

				leagueTable[m.HomeTeam] = entry
			}

			// Away team
			if entry, ok := leagueTable[m.AwayTeam]; ok {
				// games statistic
				entry.games += 1
				// win - draw - loss statistic
				if m.HomeGoals > m.AwayGoals {
					entry.loss += 1
				} else if m.HomeGoals == m.AwayGoals {
					entry.draw += 1
					entry.points += 1
				} else {
					entry.win += 1
					entry.points += 3
				}
				entry.goalDifference += m.AwayGoals - m.HomeGoals

				leagueTable[m.AwayTeam] = entry
			}
		}
	}
	fmt.Println(leagueTable)

	return nil
}
