package nba

type Stats struct {
	ActivePlayers []ActivePlayer `json:"activePlayers"`
	HomeTeam      StatsTeam      `json:"hTeam"`
	LeadChanges   string         `json:"leadChanges"`
	TimesTied     string         `json:"timesTied"`
	VisitingTeam  StatsTeam      `json:"vTeam"`
}

type StatsTeam struct {
	BiggestLead        string       `json:"biggestLead"`
	FastBreakPoints    string       `json:"fastBreakPoints"`
	Leaders            Leader       `json:"leaders"`
	LongestRun         string       `json:"longestRun"`
	PointsInPaint      string       `json:"pointsInPaint"`
	PointsOffTurnovers string       `json:"pointsOffTurnovers"`
	SecondChancePoints string       `json:"secondChancePoints"`
	Totals             ActivePlayer `json:"totals"`
}
