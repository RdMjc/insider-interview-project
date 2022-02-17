package league

import (
	"fmt"
	"interview/backend/fixture"
	"interview/backend/team"
)

type League struct {
	Fix        fixture.Fixture `json:"fixture"`
	ActiveWeek int             `json:"activeWeek"`
}

type TeamStatistics struct {
	Games          int `json:"games"`
	Points         int `json:"points"`
	Win            int `json:"wins"`
	Draw           int `json:"draws"`
	Loss           int `json:"losses"`
	GoalDifference int `json:"goalDifference"`
}

type LeagueTable struct {
	TeamList  []team.Team                  `json:"teamList"`
	TeamStats map[team.Team]TeamStatistics `json:"teamStats"`
}

func GetLeagueStatistics(l League) map[team.Team]TeamStatistics {
	var leagueTable = make(map[team.Team]TeamStatistics)

	var defaultStatistics = TeamStatistics{
		Games:          0,
		Points:         0,
		Win:            0,
		Draw:           0,
		Loss:           0,
		GoalDifference: 0,
	}

	fmt.Println(leagueTable)

	for _, t := range team.GetTeams() {
		leagueTable[t] = defaultStatistics
	}

	// go through all weeks played in the fixture
	// make team statistics
	for i := 0; i < l.ActiveWeek; i++ {
		// go through matches in that week
		for _, m := range l.Fix.Weeks[i].Matches {

			// Home Team
			if entry, ok := leagueTable[m.HomeTeam]; ok {
				// games statistic
				entry.Games += 1
				// win - draw - loss statistic
				if m.HomeGoals > m.AwayGoals {
					entry.Win += 1
					entry.Points += 3
				} else if m.HomeGoals == m.AwayGoals {
					entry.Draw += 1
					entry.Points += 1
				} else {
					entry.Loss += 1
				}
				entry.GoalDifference += m.HomeGoals - m.AwayGoals

				leagueTable[m.HomeTeam] = entry
			}

			// Away team
			if entry, ok := leagueTable[m.AwayTeam]; ok {
				// games statistic
				entry.Games += 1
				// win - draw - loss statistic
				if m.HomeGoals > m.AwayGoals {
					entry.Loss += 1
				} else if m.HomeGoals == m.AwayGoals {
					entry.Draw += 1
					entry.Points += 1
				} else {
					entry.Win += 1
					entry.Points += 3
				}
				entry.GoalDifference += m.AwayGoals - m.HomeGoals

				leagueTable[m.AwayTeam] = entry
			}
		}
	}
	fmt.Println(leagueTable)

	return leagueTable
}

func PlayOneWeek(l League) League {
	/*
		Play the given league one week and return the updated league
	*/

	// check if league is over
	if l.ActiveWeek > len(l.Fix.Weeks)-1 {
		// league is over
		// return unchanged
		return l
	} else {
		// play the active week
		fmt.Println("LEAGUE PLAYED")
		l.Fix.Weeks[l.ActiveWeek] = fixture.PlayMatches(l.Fix.Weeks[l.ActiveWeek])
		l.ActiveWeek += 1
	}

	return l
}

func PlayLeague(l League) League {
	/*
		Play the whole league starting from active week
		Return the updated league
	*/

	// play while league is not over
	for l.ActiveWeek <= len(l.Fix.Weeks)-1 {

		// play the league one week
		l = PlayOneWeek(l)
	}

	return l
}

// func GetLeagueTable(l League) ([]team.Team, map[team.Team]TeamStatistics) {
func GetLeagueTable(l League) LeagueTable {
	/*
		Returns the league table
		- First returned value is the list of teams starting from 1st position in the league
		- Second returned value is the corresponding team's statistics
	*/

	teamList := team.GetTeams()
	teamStats := GetLeagueStatistics(l)

	// APPLY A MODIFIED RADIX SORT
	// first sort according to goal difference
	var isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(teamList)-1; i++ {
			// compare points of i'th and i+1'th teams points
			if teamStats[teamList[i]].GoalDifference < teamStats[teamList[i+1]].GoalDifference {
				// swap teams in their position
				teamList[i], teamList[i+1] = teamList[i+1], teamList[i]
				isDone = false
			}
		}
	}

	// then, sort by points
	// apply bubble sort
	isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(teamList)-1; i++ {
			// compare points of i'th and i+1'th teams points
			if teamStats[teamList[i]].Points < teamStats[teamList[i+1]].Points {
				// swap teams in their position
				teamList[i], teamList[i+1] = teamList[i+1], teamList[i]
				isDone = false
			}
		}
	}

	return LeagueTable{teamList, teamStats}
}

func PrintLeagueTable(l League) {
	// get league table
	// teamsSorted, teamsStats := GetLeagueTable(l)
	leagueTable := GetLeagueTable(l)
	teamsSorted := leagueTable.TeamList
	teamsStats := leagueTable.TeamStats

	fmt.Printf("%15v \t PTS \t P \t W \t D \t L \t GD \n", "Teams")
	for _, te := range teamsSorted {
		fmt.Printf("%15v \t %d \t %d \t %d \t %d \t %d \t %d \n",
			te.Name,
			teamsStats[te].Points,
			teamsStats[te].Games,
			teamsStats[te].Win,
			teamsStats[te].Draw,
			teamsStats[te].Loss,
			teamsStats[te].GoalDifference,
		)
	}

}
