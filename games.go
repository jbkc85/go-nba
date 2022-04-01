package nba

type Game struct {
	Clock                 string       `json:"clock"`
	EndTimeUTC            string       `json:"endTimeUTC"`
	ExtendedStatusNum     int64        `json:"extendedStatusNum"`
	GameDuration          GameDuration `json:"gameDuration"`
	GameID                string       `json:"gameId"`
	HasGameBookPDF        bool         `json:"hasGameBookPdf"`
	IsGameActivated       bool         `json:"isGameActivated"`
	IsBuzzerBeater        bool         `json:"isBuzzerBeater"`
	IsPreviewArticleAvail bool         `json:"isPreviewArticleAvail"`
	IsRecapArticleAvail   bool         `json:"isRecapArticleAvail"`
	IsStartTimeTBD        bool         `json:"isStartTimeTBD"`
	SeasonStageID         int64        `json:"seasonStageId"`
	SeasonYear            string       `json:"seasonYear"`
	StatusNum             int64        `json:"statusNum"`
	StartTimeEastern      string       `json:"startTimeEastern"`
	StartTimeUTC          string       `json:"startTimeUTC"`
	StartDateEastern      string       `json:"startDateEastern"`
	Attendance            string       `json:"attendance"`
	//Period                Period       `json:"period"`
	VTeam GameTeam `json:"vTeam"`
	HTeam GameTeam `json:"hTeam"`
}

type GameDuration struct {
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
}

type GameLinescore struct {
	Score string `json:"score"`
}

type GameStats struct {
	FastBreakPoints    string            `json:"fastBreakPoints"`
	PointsInPaint      string            `json:"pointsInPaint"`
	BiggestLead        string            `json:"biggestLead"`
	SecondChancePoints string            `json:"secondChancePoints"`
	PointsOffTurnovers string            `json:"pointsOffTurnovers"`
	LongestRun         string            `json:"longestRun"`
	Totals             ActivePlayer      `json:"totals"`
	Leaders            map[string]Leader `json:"leaders"`
}

type GamePeriod struct {
	Current       int64 `json:"current"`
	Type          int64 `json:"type"`
	MaxRegular    int64 `json:"maxRegular"`
	IsHalftime    bool  `json:"isHalftime"`
	IsEndOfPeriod bool  `json:"isEndOfPeriod"`
}

type GameTeam struct {
	Win        string          `json:"win"`
	Linescore  []GameLinescore `json:"linescore"`
	Loss       string          `json:"loss"`
	SeriesWin  string          `json:"seriesWin"`
	SeriesLoss string          `json:"seriesLoss"`
	Score      string          `json:"score"`
	TeamID     string          `json:"teamId"`
	TriCode    string          `json:"triCode"`
}
