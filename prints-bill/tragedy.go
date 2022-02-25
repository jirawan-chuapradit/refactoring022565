package main

import "math"

type Tragedy struct {
	Name string
}

func (play Tragedy) name() string {
	return play.Name
}
func (play Tragedy) volumeCreditsFor(audience int) (volumeCredits float64) {
	volumeCredits += math.Max(float64(audience-30), 0)
	return
}
func (play Tragedy) AmountFor(audience int) (amount float64) {
	amount = 40000
	if audience > 30 {
		amount += 1000 * (float64(audience - 30))
	}
	return amount
}
