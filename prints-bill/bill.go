package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Kind Kind
}

type Plays map[string]Play

type Kind string

const Comedy Kind = "comedy"
const Tragedy Kind = "tragedy"

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func kind(play Play) Kind {
	return play.Kind
}

func name(play Play) string {
	return play.Name
}

func playeFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func volumeCreditsFor(plays Plays, perf Performance) (volumeCredits float64) {
	volumeCredits += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if kind(playeFor(plays, perf)) == Comedy {
		volumeCredits += math.Floor(float64(perf.Audience / 5))
	}
	return
}

func statement(invoice Invoice, plays Plays) string {
	volumeCredits := 0.0

	totalAmount := totalAmount(plays, invoice)
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(plays, perf)
	}
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", name(play), AmountFor(plays, perf)/100, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", volumeCredits)
	return result
}

func totalAmount(plays Plays, invoice Invoice) (totalAmount float64) {
	for _, perf := range invoice.Performances {
		totalAmount += AmountFor(plays, perf)
	}
	return totalAmount
}

func AmountFor(plays Plays, perf Performance) (amount float64) {
	switch kind(playeFor(plays, perf)) {
	case Tragedy:
		amount = 40000
		if perf.Audience > 30 {
			amount += 1000 * (float64(perf.Audience - 30))
		}
	case Comedy:
		amount = 30000
		if perf.Audience > 20 {
			amount += 10000 + 500*(float64(perf.Audience-20))
		}
		amount += 300 * float64(perf.Audience)
	default:
		panic(fmt.Errorf("unknow type: %s", kind(playeFor(plays, perf))))
	}
	return amount
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := Plays{
		"hamlet":  {"Hamlet", "tragedy"},
		"as-like": {"As You Like It", "comedy"},
		"othello": {"Othello", "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
