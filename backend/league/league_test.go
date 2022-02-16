package league

import (
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
