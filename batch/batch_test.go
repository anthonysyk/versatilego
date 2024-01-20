package batch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatchProcess(t *testing.T) {
	t.Parallel()

	type PlayerHistoryScore struct {
		Player string
		Score  int
	}

	iterations := 0

	GetPlayerHistoryScore := func(players []string) []PlayerHistoryScore {
		iterations++
		scores := make([]PlayerHistoryScore, len(players))
		for i, player := range players {
			scores[i] = PlayerHistoryScore{Player: player, Score: i * 10}
		}
		return scores
	}

	tests := []struct {
		label              string
		players            []string
		batchSize          int
		expectedIterations int
		expectedResult     []PlayerHistoryScore
	}{
		{
			label:              "nominal case",
			batchSize:          3,
			players:            []string{"player1", "player2", "player3", "player4", "player5", "player6", "player7"},
			expectedResult:     []PlayerHistoryScore{{"player1", 0}, {"player2", 10}, {"player3", 20}, {"player4", 0}, {"player5", 10}, {"player6", 20}, {"player7", 0}},
			expectedIterations: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			t.Parallel()
			scores := Process(test.players, test.batchSize, GetPlayerHistoryScore)
			assert.Equal(t, test.expectedIterations, iterations)
			assert.Equal(t, test.expectedResult, scores)
			iterations = 0
		})
	}
}
