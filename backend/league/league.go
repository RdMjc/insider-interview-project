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

func GetLeagueStatistics(l League) map[team.Team]TeamStatistics {
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

	return leagueTable
}

func PlayOneWeek(l League) League {
	/*
		Play the given league one week and return the updated league
	*/

	// check if league is over
	if l.activeWeek > len(l.fix.Weeks)-1 {
		// league is over
		// return unchanged
		return l
	} else {
		// play the active week
		fmt.Println("LEAGUE PLAYED")
		l.fix.Weeks[l.activeWeek] = fixture.PlayMatches(l.fix.Weeks[l.activeWeek])
		l.activeWeek += 1
	}

	return l
}

func PlayLeague(l League) League {
	/*
		Play the whole league starting from active week
		Return the updated league
	*/

	// play while league is not over
	for l.activeWeek <= len(l.fix.Weeks)-1 {

		// play the league one week
		l = PlayOneWeek(l)
	}

	return l
}

func GetLeagueTable(l League) ([]team.Team, map[team.Team]TeamStatistics) {
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
			if teamStats[teamList[i]].goalDifference < teamStats[teamList[i+1]].goalDifference {
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
			if teamStats[teamList[i]].points < teamStats[teamList[i+1]].points {
				// swap teams in their position
				teamList[i], teamList[i+1] = teamList[i+1], teamList[i]
				isDone = false
			}
		}
	}

	return teamList, teamStats
}

func PrintLeagueTable(l League) {
	// get league table
	teamsSorted, teamsStats := GetLeagueTable(l)

	fmt.Printf("%15v \t PTS \t P \t W \t D \t L \t GD \n", "Teams")
	for _, te := range teamsSorted {
		fmt.Printf("%15v \t %d \t %d \t %d \t %d \t %d \t %d \n",
			te.Name,
			teamsStats[te].points,
			teamsStats[te].games,
			teamsStats[te].win,
			teamsStats[te].draw,
			teamsStats[te].loss,
			teamsStats[te].goalDifference,
		)
	}

}
