package main

import (
	"fmt"
	"interview/backend/fixture"
	"interview/backend/league"
	"interview/backend/team"
	"net/http"

	"github.com/gin-gonic/gin"
)

var f, _ = fixture.CreateFixture(team.GetTeams())

var l = league.League{
	Fix:        f,
	ActiveWeek: 0,
}

func init() {
	l = league.PlayOneWeek(l)
}

func main() {

	router := gin.Default()
	router.GET("/league-table", getLeagueTable)
	router.GET("/matches", getMatches)
	router.POST("/matches", playMatches)

	router.Run("localhost:8080")
}

// type jsonSerializableTeamNameStats struct {
// 	name string
// 	stats league.TeamStatistics
// }

func getLeagueTable(c *gin.Context) {
	// obtain a JSON serializable league table
	// Example Format
	// {
	// 	"teams": [
	// 		"Manchester City" : {Stats},
	// 		"Chelsea" : {Stats},
	// 		"Liverpool" : {Stats},
	// 		"Manchester United" : {Stats}
	// 	]
	// }

	// get league table
	leagueTable := league.GetLeagueTable(l)

	jsonSerializableResponse := make(map[string][]map[string]league.TeamStatistics)

	for _, t := range leagueTable.TeamList {
		teamNameStats := make(map[string]league.TeamStatistics)
		teamNameStats[t.Name] = leagueTable.TeamStats[t]
		// fmt.Println(leagueTable.TeamStats[t])
		fmt.Println("******************")
		fmt.Println(teamNameStats)
		jsonSerializableResponse["teams"] = append(jsonSerializableResponse["teams"], teamNameStats)
	}

	fmt.Println("################")
	fmt.Println(jsonSerializableResponse)

	c.IndentedJSON(http.StatusOK, jsonSerializableResponse)
}

func getMatches(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, l.Fix.Weeks[l.ActiveWeek])
}

func playMatches(c *gin.Context) {

	// play league one week
	l = league.PlayOneWeek(l)

	c.IndentedJSON(http.StatusOK, "")
}
