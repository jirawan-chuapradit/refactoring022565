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

type Rate struct {
	Play     Play
	Amount   float64
	Credit   float64
	Audience int
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
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

func volumeCreditsFor(play Play, audience int) (volumeCredits float64) {
	volumeCredits += math.Max(float64(audience-30), 0)
	// add extra credit for every ten comedy attendees
	if play.Kind == Comedy {
		volumeCredits += math.Floor(float64(audience / 5))
	}
	return
}

func statement(invoice Invoice, plays Plays) string {
	var rate = []Rate{}
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		amount := AmountFor(play, perf.Audience)
		audience := perf.Audience
		r := Rate{
			Play:     play,
			Amount:   amount,
			Credit:   volumeCreditsFor(play, perf.Audience),
			Audience: audience,
		}
		rate = append(rate, r)
	}
	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rate,
		TotalAmount:        totalAmount(rate),
		TotalVolumeCredits: totalVolumeCredits(rate),
	}
	return printPlanText(bill)
}

func printPlanText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, r := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.Play.Name, r.Amount/100, r.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func totalVolumeCredits(rate []Rate) (totalVolumeCredits float64) {
	for _, r := range rate {
		totalVolumeCredits += r.Credit
	}
	return
}

func totalAmount(rate []Rate) (totalAmount float64) {
	for _, r := range rate {
		totalAmount += r.Amount
	}
	return
}
func AmountFor(play Play, audience int) (amount float64) {
	switch play.Kind {
	case Tragedy:
		amount = 40000
		if audience > 30 {
			amount += 1000 * (float64(audience - 30))
		}
	case Comedy:
		amount = 30000
		if audience > 20 {
			amount += 10000 + 500*(float64(audience-20))
		}
		amount += 300 * float64(audience)
	default:
		panic(fmt.Errorf("unknow type: %s", play.Kind))
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
