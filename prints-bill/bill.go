package main

import (
	"fmt"
)

type Player interface {
	name() string
	volumeCreditsFor(audience int) (volumeCredits float64)
	AmountFor(audience int) (amount float64)
}

type Play struct {
	Name string
	Kind string
}

type Plays map[string]Player

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Rate struct {
	Play     Player
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

func playeFor(plays Plays, perf Performance) Player {
	return plays[perf.PlayID]
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

func statement(invoice Invoice, plays Plays) string {
	var rate = []Rate{}
	for _, perf := range invoice.Performances {
		play := playeFor(plays, perf)
		amount := play.AmountFor(perf.Audience)
		audience := perf.Audience
		r := Rate{
			Play:     play,
			Amount:   amount,
			Credit:   play.volumeCreditsFor(perf.Audience),
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
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.Play.name(), r.Amount/100, r.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
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
		"hamlet":  Tragedy{"Hamlet"},
		"as-like": Comedy{"As You Like It"},
		"othello": Tragedy{"Othello"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
