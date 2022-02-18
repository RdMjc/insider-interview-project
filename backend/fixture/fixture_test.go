package fixture

import (
	"fmt"
	"interview/backend/team"
	"testing"
)

func TestCreateOneTourFixture(t *testing.T) {
	teamList := team.GetTeams()

	fmt.Println(teamList)

	resultFixture, _ := createOneTourFixture(teamList)

	//fmt.Println(resultFixture)
	Print(resultFixture)

	// check if each team matches others once
	// matchStats := make(map[string]map[string]int)

	// // assign default values to 0
	// for i := 0; i < len(teamList); i++ {
	// 	te := teamList[i]
	// 	fmt.Println(te.Name)
	// 	// for _, te := range teamList {
	// 	matchStats[te.Name] = make(map[string]int)
	// 	for j := 0; j < len(teamList); j++ {
	// 		te2 := teamList[j]
	// 		// for _, te2 := range teamList {
	// 		matchStats[te.Name][te2.Name] = 0
	// 		// matchStats[te2.Name][te.Name] = 0
	// 	}
	// }

	// fmt.Println(matchStats)

	// // go through weeks in the fixture
	// for _, w := range resultFixture.Weeks {
	// 	for _, m := range w.Matches {
	// 		matchStats[m.HomeTeam.Name][m.AwayTeam.Name] = 1
	// 		// matchStats[m.AwayTeam.Name][m.HomeTeam.Name] += 1
	// 	}
	// }

	// for _, te := range teamList {
	// 	for _, te2 := range teamList {
	// 		if te.Name != te2.Name {
	// 			if matchStats[te.Name][te2.Name] != 1 {
	// 				t.Fatal("Teams don't match once")
	// 			}
	// 		}
	// 	}
	// }
}

func TestCreateFixture(t *testing.T) {
	teamList := []team.Team{
		{"Manchester City", 80},
		{"Chelsea", 75},
		{"Liverpool", 75},
		{"Manchester United", 65},
	}

	resultFixture, _ := CreateFixture(teamList)

	//fmt.Println(resultFixture)
	Print(resultFixture)
}
