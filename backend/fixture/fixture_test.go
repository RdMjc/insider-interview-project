package fixture

import (
	"interview/backend/team"
	"testing"
)

func TestCreateOneTourFixture(t *testing.T) {
	teamList := []team.Team{
		{"Manchester City", 80},
		{"Chelsea", 75},
		{"Liverpool", 75},
		{"Manchester United", 65},
	}

	resultFixture, _ := createOneTourFixture(teamList)

	//fmt.Println(resultFixture)
	Print(resultFixture)
}
