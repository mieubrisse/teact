package scoreboard

type ScoreboardOpts func(scoreboard Scoreboard)

func WithHomeScore(score int) ScoreboardOpts {
	return func(scoreboard Scoreboard) {
		scoreboard.SetHomeTeamScore(score)
	}
}

func WithAwayScore(score int) ScoreboardOpts {
	return func(scoreboard Scoreboard) {
		scoreboard.SetAwayTeamScore(score)
	}
}
