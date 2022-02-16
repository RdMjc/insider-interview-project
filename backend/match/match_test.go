package match

import (
	"fmt"
	"interview/backend/team"
	"testing"
)

func TestPlayPowerfulTeamWinsOnAverageInLargeNumber(t *testing.T) {
	team1 := team.Team{
		Name:  "team1",
		Power: 100,
	}

	team2 := team.Team{
		Name:  "team2",
		Power: 10,
	}

	m := Match{
		HomeTeam:  team1,
		AwayTeam:  team2,
		HomeGoals: 0,
		AwayGoals: 0,
	}

	nDraw := 0
	nHomeWin := 0
	nAwayWin := 0

	for i := 0; i < 1000; i++ {
		m = Play(m)
		if m.HomeGoals > m.AwayGoals {
			nHomeWin += 1
		} else if m.AwayGoals > m.HomeGoals {
			nAwayWin += 1
		} else {
			nDraw += 1
		}
	}

	fmt.Println("Home win", nHomeWin)
	fmt.Println("Away win", nAwayWin)
	fmt.Println("Draw", nDraw)

	// test if homewin is greater than both draws and awaywins
	if nHomeWin < nDraw || nHomeWin < nAwayWin {
		t.Fatalf("Powerful Home team did not win enough")
	}

	fmt.Println(m)
}

func TestPlaySimilarTeamsDontMakeHugeDifferencesInWinsOrLosses(t *testing.T) {
	team1 := team.Team{
		Name:  "team1",
		Power: 50,
	}

	team2 := team.Team{
		Name:  "team2",
		Power: 50,
	}

	m := Match{
		HomeTeam:  team1,
		AwayTeam:  team2,
		HomeGoals: 0,
		AwayGoals: 0,
	}

	nDraw := 0
	nHomeWin := 0
	nAwayWin := 0

	for i := 0; i < 1000; i++ {
		m = Play(m)
		if m.HomeGoals > m.AwayGoals {
			nHomeWin += 1
		} else if m.AwayGoals > m.HomeGoals {
			nAwayWin += 1
		} else {
			nDraw += 1
		}
	}

	fmt.Println("Home win", nHomeWin)
	fmt.Println("Away win", nAwayWin)
	fmt.Println("Draw", nDraw)

	// test home wins are not 4 times of draws or away wins
	if nHomeWin > 4*nDraw || nHomeWin > 4*nAwayWin {
		t.Fatal("Too much home wins")
	}

	// test away wins are not 4 times of draws or home wins
	if nAwayWin > 4*nDraw || nAwayWin > 4*nHomeWin {
		t.Fatal("Too much away wins")
	}

	fmt.Println(m)
}
