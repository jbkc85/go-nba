package nba

import (
	"fmt"
	"time"
)

type Boxscore struct {
	Internal Internal `json:"_internal"`
	Stats    Stats    `json:"stats"`
}

func GetBoxscore(gameID string) Boxscore {
	// endpoint:
	// v1/YYYYMMDD/{{ gameid }}_boxscore.json
	loc, _ := time.LoadLocation(timezone)
	today := time.Now().In(loc)
	var boxscore Boxscore

	httpRequest(
		fmt.Sprintf("v1/%s/%s_boxscore.json", today.Format(datetimeLayout), gameID),
		&boxscore,
	)

	return boxscore
}
