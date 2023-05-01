package scoreboard

import (
	"fmt"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/text"
)

type Scoreboard interface {
	components.Component

	GetHomeTeamScore() int
	SetHomeTeamScore(score int) Scoreboard

	GetAwayTeamScore() int
	SetAwayTeamScore(score int) Scoreboard
}

type scoreboardImpl struct {
	components.Component

	homeScore int
	awayScore int

	display text.Text
}

func New(opts ...ScoreboardOpts) Scoreboard {
	display := text.New()

	result := &scoreboardImpl{
		Component: display,
		display:   display,
		homeScore: 0,
		awayScore: 0,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (s scoreboardImpl) GetHomeTeamScore() int {
	return s.homeScore
}

func (s *scoreboardImpl) SetHomeTeamScore(score int) Scoreboard {
	s.homeScore = score
	return s
}

func (s scoreboardImpl) GetAwayTeamScore() int {
	return s.awayScore
}

func (s *scoreboardImpl) SetAwayTeamScore(score int) Scoreboard {
	s.awayScore = score
	return s
}

func (s *scoreboardImpl) updateDisplay() {
	s.display.SetContents(fmt.Sprintf("Home: %v, Away: %v", s.homeScore, s.awayScore))
}
