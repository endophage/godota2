package dota2

type MatchDetail struct {
	Players               []PlayerDetail `json:"players"`
	RadiantWin            bool           `json:"radiant_win"`
	Duration              int            `json:"duration"`
	StartTime             int            `json:"start_time"`
	MatchID               int            `json:"match_id"`
	MatchSeqNum           int            `json:"match_seq_num"`
	TowerStatusRadiant    int            `json:"tower_status_radiant"`
	TowerStatusDire       int            `json:"tower_status_dire"`
	BarracksStatusRadiant int            `json:"barracks_status_radiant"`
	BarracksStatusDire    int            `json:"barracks_status_dire"`
	Cluster               int            `json:"cluster"`
	FirstBloodTime        int            `json:"first_blood_time"`
	LobbyType             int            `json:"lobby_type"`
	HumanPlayers          int            `json:"human_players"`
	LeagueID              int            `json:"league_id"`
	PositiveVotes         int            `json:"positive_votes"`
	NegativeVotes         int            `json:"negative_votes"`
	GameMode              int            `json:"game_mode"`
	Engine                int            `json:"engine"`
}

type PlayerDetail struct {
	AccountID    int `json:"account_id"`
	PlayerSlot   int `json:"player_slot"`
	HeroID       int `json:"hero_id"`
	Item0        int `json:"item_0"`
	Item1        int `json:"item_1"`
	Item2        int `json:"item_2"`
	Item3        int `json:"item_3"`
	Item4        int `json:"item_4"`
	Item5        int `json:"item_5"`
	Kills        int `json:"kills"`
	Deaths       int `json:"deaths"`
	Assists      int `json:"assists"`
	LeaverStatus int `json:"leaver_status"`
	Gold         int `json:"gold"`
	LastHits     int `json:"last_hits"`
	Denies       int `json:"denies"`
	GoldPerMin   int `json:"gold_per_min"`
	XPPerMin     int `json:"xp_per_min"`
	GoldSpent    int `json:"gold_spend"`
	HeroDamage   int `json:"hero_damage"`
	TowerDamage  int `json:"tower_damage"`
	HeroHealing  int `json:"hero_healing"`
	Level        int `json:"level"`
}

type Matches struct {
	Status           int     `json:"status"`
	NumResults       int     `json:"num_results"`
	TotalResults     int     `json:"total_results"`
	ResultsRemaining int     `json:"results_remaining"`
	Matches          []Match `json:"matches"`
}

type Match struct {
	MatchID       int      `json:"match_id"`
	MatchSeqNum   int      `json:"match_seq_num"`
	StartTime     int      `json:"start_time"`
	LobbyType     int      `json:"lobby_type"`
	RadiantTeamID int      `json:"radiant_team_id"`
	DireTeamID    int      `json:"dire_team_id"`
	Players       []Player `json:"players"`
}

type Player struct {
	AccountID  int `json:"account_id"`
	PlayerSlot int `json:"player_slot"`
	HeroID     int `json:"hero_id"`
}

type Api interface {
	MatchDetail(id int) (MatchDetail, error)
	Matches(from int) (Matches, error)
}
