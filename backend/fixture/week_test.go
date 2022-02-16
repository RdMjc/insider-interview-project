package fixture

import (
	"fmt"
	"interview/backend/team"
	"testing"
)

func TestPlayMatches(t *testing.T) {
	teamList := []team.Team{
		{"Manchester City", 80},
		{"Chelsea", 75},
		{"Liverpool", 75},
		{"Manchester United", 65},
	}

	resultFixture, _ := CreateFixture(teamList)

	firstWeek := resultFixture.Weeks[0]

	fmt.Println("BEFORE *** ", firstWeek)

	firstWeekPlayed := PlayMatches(firstWeek)

	fmt.Println("AFTER *** ", firstWeekPlayed)
}
