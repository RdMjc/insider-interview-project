package league

import (
	"fmt"
	"interview/backend/fixture"
	"interview/backend/team"
	"testing"
)

func TestGetLeagueStatistics(t *testing.T) {
	var f fixture.Fixture
	f, _ = fixture.CreateFixture(team.GetTeams())

	var l = League{
		fix:        f,
		activeWeek: 1,
	}

	GetLeagueStatistics(l)
}

func TestPlayOneWeek(t *testing.T) {
	var f fixture.Fixture
	f, _ = fixture.CreateFixture(team.GetTeams())

	var l = League{
		fix:        f,
		activeWeek: 1,
	}

	// get league stats
	leagueStats := GetLeagueStatistics(l)

	fmt.Println("Before week played")
	fmt.Println(leagueStats)

	// play the league one week
	l = PlayOneWeek(l)

	// get league stats
	leagueStats = GetLeagueStatistics(l)

	fmt.Println("After week played")
	fmt.Println(leagueStats)
}

func TestPlayLeague(t *testing.T) {
	var f fixture.Fixture
	f, _ = fixture.CreateFixture(team.GetTeams())

	var l = League{
		fix:        f,
		activeWeek: 0,
	}

	// get league stats
	leagueStats := GetLeagueStatistics(l)

	fmt.Println("LEAGUE STARTING")
	fmt.Println(leagueStats)

	// play the league one week
	l = PlayLeague(l)

	// get league stats
	leagueStats = GetLeagueStatistics(l)

	fmt.Println("LEAGUE IS OVER")
	fmt.Println(leagueStats)
}
