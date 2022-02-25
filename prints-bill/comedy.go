package main

import (
	"math"
)

type Comedy struct {
	Name string
}

func (play Comedy) name() string {
	return play.Name
}
func (play Comedy) volumeCreditsFor(audience int) (volumeCredits float64) {
	volumeCredits += math.Max(float64(audience-30), 0)
	volumeCredits += float64(audience / 5)
	return
}
func (play Comedy) AmountFor(audience int) (amount float64) {
	amount = 30000
	if audience > 20 {
		amount += 10000 + 500*(float64(audience-20))
	}
	amount += 300 * float64(audience)

	return amount
}
