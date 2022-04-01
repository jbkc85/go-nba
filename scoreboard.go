package nba

import (
	"fmt"
	"time"
)

type Scoreboard struct {
	Games         []Game   `json:"games"`
	Internal      Internal `json:"_internal"`
	NumberOfGames int      `json:"numGames"`
}

func GetScoreboard() Scoreboard {
	// endpoint:
	// /prod/v1/YYYYMMDD/scoreboard.json
	loc, _ := time.LoadLocation(timezone)
	today := time.Now().In(loc)
	var scoreboard Scoreboard

	httpRequest(
		fmt.Sprintf("v1/%s/scoreboard.json", today.Format(datetimeLayout)),
		&scoreboard,
	)

	return scoreboard
}
