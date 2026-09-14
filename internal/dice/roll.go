package dice

import "math/rand"

type RollInput struct {
	Skill    int
	Shade    Shade
	Obstacle int
}

type RollResult struct {
	Dice      []int
	Successes int
	Passed    bool
}

func rollDie() int {
	return rand.Intn(6) + 1
}

func rollDice(count int) []int {
	dice := make([]int, count)

	for i := 0; i < count; i++ {
		dice[i] = rollDie()
	}

	return dice
}

func isSuccess(value int, shade Shade) bool {
	switch shade {
	case Black:
		return value >= 4
	case Gray:
		return value >= 3
	case White:
		return value >= 2
	default:
		return false
	}
}

func countSuccesses(dice []int, shade Shade) int {
	passed := 0
	for i := 0; i < len(dice); i++ {
		if isSuccess(dice[i], shade) {
			passed++
		}
	}
	return passed
}

func Roll(r RollInput) RollResult {
	result := new(RollResult)
	result.Dice = rollDice(r.Skill)
	result.Successes = countSuccesses(result.Dice, r.Shade)
	result.Passed = result.Successes >= r.Obstacle
	return *result
}
