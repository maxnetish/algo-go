package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type MakeBigArrayParams struct {
	NumberOfElements int
	From             int
	To               int
}

// MakeBigArray generates slice of integers with numberOfElements elements
func MakeBigArray(params MakeBigArrayParams) []int {
	result := make([]int, params.NumberOfElements, params.NumberOfElements)
	for ind := range result {
		result[ind] = getRandomInRange(params.From, params.To)
	}
	return result
}

func getRandomInRange(from int, to int) int {
	return rand.Intn(to-from) + from
}
