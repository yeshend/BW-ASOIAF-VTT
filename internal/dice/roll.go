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
