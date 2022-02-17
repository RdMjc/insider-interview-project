package match

import (
	"fmt"
	"interview/backend/team"
	"math/rand"
	"time"
)

type Match struct {
	HomeTeam  team.Team `json:"homeTeam"`
	AwayTeam  team.Team `json:"AwayTeam"`
	HomeGoals int       `json:"homeGoals"`
	AwayGoals int       `json:"awayGoals"`
}

func Print(match Match) {
	fmt.Printf("%v %d-%d %v \n", match.HomeTeam.Name, match.HomeGoals, match.AwayGoals, match.AwayTeam.Name)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Play(m Match) Match {
	/*
		%50 probability of no goal
		%50 probability distributed according to teams' strength
	*/
	probabilityNoGoal := float64(50)
	probabilityHomeGoal := float64((float64(m.HomeTeam.Power) / float64(m.HomeTeam.Power+m.AwayTeam.Power)) * 100 / 2)
	// probabilityAwayGoal := float64(100 - float64(probabilityNoGoal) - float64(probabilityHomeGoal))

	// initial scores
	homeScore := 0
	awayScore := 0

	// split 90 minutes into equal 10min intervals
	// simulate every 10min
	for i := 0; i < 9; i++ {
		randEvent := float64(rand.Intn(100))

		// no goal
		if randEvent < probabilityNoGoal {
		} else if (probabilityNoGoal <= randEvent) && (randEvent < probabilityNoGoal+probabilityHomeGoal) {
			// home goal
			homeScore += 1
		} else {
			// away goal
			awayScore += 1
		}
	}

	// assign scores to match
	m.HomeGoals = homeScore
	m.AwayGoals = awayScore

	return m
}
