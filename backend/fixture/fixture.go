package fixture

import (
	"fmt"
	"interview/backend/match"
	"interview/backend/team"
)

type Fixture struct {
	Weeks []Week `json:"weeks"`
}

func CreateFixture(teams []team.Team) (Fixture, error) {
	/*
		Performs round-robin tournament match algorithm (Circle Method), described in
		https://en.wikipedia.org/wiki/Round-robin_tournament#Scheduling_algorithm
	*/

	// get one way fixture
	var oneWayFixture, _ = createOneTourFixture(teams)

	// append home-away teams reversed weeks on top of oneWayFixture
	// to get full fixture
	for _, week := range oneWayFixture.Weeks {

		// reverse home-away teams
		var reversedWeek = reverseWeek(week)

		oneWayFixture.Weeks = append(oneWayFixture.Weeks, reversedWeek)
	}

	return oneWayFixture, nil
}

func createOneTourFixture(teams []team.Team) (Fixture, error) {
	// Only creates one tour of the fixture
	// Duplicating and reversing home-away teams will create the full fixture
	// Performs Simple Method in https://en.wikipedia.org/wiki/Round-robin_tournament#Scheduling_algorithm

	// create empty fixture
	var newFixture Fixture

	numberOfTeams := len(teams)
	numberOfRounds := numberOfTeams - 1

	constantTeam := teams[0]
	row1 := teams[0 : numberOfTeams/2]
	row2 := teams[numberOfTeams/2 : numberOfTeams]

	row1 = removeAtIndex(row1, 0)

	// for each round
	for r := 1; r <= numberOfRounds; r++ {
		row1 = prepend(row1, constantTeam)

		// START OF A ROUND
		var matches []match.Match
		for i := 0; i < min(len(row1), len(row2)); i++ {
			var newMatch = match.Match{
				HomeTeam:  row1[i],
				AwayTeam:  row2[i],
				HomeGoals: 0,
				AwayGoals: 0,
			}

			matches = append(matches, newMatch)
		}

		var thisWeek = Week{Matches: matches}
		fmt.Println(thisWeek)

		newFixture.Weeks = append(newFixture.Weeks, thisWeek)

		// END OF A ROUND

		row1 = removeAtIndex(row1, 0)

		// add first element of row2 to row1
		row1 = prepend(row1, row2[0])

		// remove row2 first element
		row2 = removeAtIndex(row2, 0)

		// add last element of row1 to row2
		row2 = append(row2, row1[len(row1)-1])

		// remove row1 last element
		row1 = removeAtIndex(row1, len(row1)-1)

	}

	return newFixture, nil
}

func removeAtIndex(slice []team.Team, s int) []team.Team {
	/* https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang */
	return append(slice[:s], slice[s+1:]...)
}

func prepend(slice []team.Team, item team.Team) []team.Team {
	/* https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang */
	var newSlice []team.Team
	newSlice = append(newSlice, item)
	for _, e := range slice {
		newSlice = append(newSlice, e)
	}
	return newSlice
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Print(fix Fixture) {

	// go through over weeks
	for i, week := range fix.Weeks {
		fmt.Printf("----------- Week %d ----------- \n", i)
		printWeek(week)
	}
}
