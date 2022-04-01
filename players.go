package nba

import (
	"strings"
)

type Players struct {
	Internal Internal                  `json:"_internal"`
	Leagues  map[string][]ActivePlayer `json:"league"`
}

type ActivePlayer struct {
	FirstName string           `json:"firstName,omitempty"`
	LastName  string           `json:"lastName,omitempty"`
	PersonID  string           `json:"personId,omitempty"`
	TeamID    string           `json:"teamId,omitempty"`
	IsOnCourt bool             `json:"isOnCourt,omitempty"`
	Points    string           `json:"points"`
	Position  string           `json:"pos,omitempty"`
	Minutes   string           `json:"min"`
	Fgm       string           `json:"fgm"`
	Fga       string           `json:"fga"`
	Fgp       string           `json:"fgp"`
	Ftm       string           `json:"ftm"`
	Fta       string           `json:"fta"`
	Ftp       string           `json:"ftp"`
	Tpm       string           `json:"tpm"`
	Tpa       string           `json:"tpa"`
	Tpp       string           `json:"tpp"`
	OffReb    string           `json:"offReb"`
	DefReb    string           `json:"defReb"`
	TotReb    string           `json:"totReb"`
	Assists   string           `json:"assists"`
	PFouls    string           `json:"pFouls"`
	Steals    string           `json:"steals"`
	Turnovers string           `json:"turnovers"`
	Blocks    string           `json:"blocks"`
	PlusMinus string           `json:"plusMinus"`
	Dnp       string           `json:"dnp,omitempty"`
	SortKey   map[string]int64 `json:"sortKey,omitempty"`
}

type Leader struct {
	Players []Player `json:"players"`
	Value   int      `json:"value"`
}

type Player struct {
	CollegeName     string `json:"collegeName"`
	Country         string `json:"country"`
	DateOfBirthUTC  string `json:"dateOfBirthUTC"`
	FirstName       string `json:"firstName"`
	HeightFeet      string `json:"heightFeet"`
	HeightInches    string `json:"heightInches"`
	IsActive        bool   `json:"isActive"`
	Jersey          string `json:"jersey"`
	LastAffiliation string `json:"lastAffiliation"`
	LastName        string `json:"lastName"`
	NbaDebutYear    string `json:"nbaDebutYear"`
	PersonID        string `json:"personId"`
	TeamID          string `json:"teamId"`
	WeightPounds    string `json:"weightPounds"`
	WeightKilograms string `json:"weightKilograms"`
	YearsPro        string `json:"yearsPro"`
}

func (a *ActivePlayer) GetPosition() string {
	switch strings.ToUpper(a.Position) {
	case "C":
		return "Center"
	case "PG":
		return "Point Guard"
	case "PF":
		return "Power Forward"
	case "SF":
		return "Small Forward"
	case "SG":
		return "Shooting Guard"
	default:
		return "N/A"
	}
}

func GetActivePlayers(league string) map[string]ActivePlayer {
	// endpoint:
	// v1/2021/players.json
	var players Players
	mappedPlayers := make(map[string]ActivePlayer)

	httpRequest("v1/2021/players.json", &players)

	for _, player := range players.Leagues[league] {
		mappedPlayers[player.PersonID] = player
	}
	return mappedPlayers
}
