package models

type PublishScoreBody struct {
	Score int64 `json:"score" binding:"required"`
}

type PublishScoreAmqp struct {
	GameId    string `json:"game_id"`
	PlayerID  string `json:"player_id"`
	Score     int64  `json:"score" binding:"required"`
	CreatedAt int64  `json:"created_at"`
}
