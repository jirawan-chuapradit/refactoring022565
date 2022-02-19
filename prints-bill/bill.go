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

func GetKind(play Play) Kind {
	return play.Kind
}

func GetName(play Play) string {
	return play.Name
}

func GetPlayer(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func statement(invoice Invoice, plays Plays) string {
	totalAmount := 0.0
	volumeCredits := 0.0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, perf := range invoice.Performances {
		play := GetPlayer(plays, perf)
		thisAmount := 0.0

		switch GetKind(play) {
		case Tragedy:
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (float64(perf.Audience - 30))
			}
		case Comedy:
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(float64(perf.Audience-20))
			}
			thisAmount += 300 * float64(perf.Audience)
			volumeCredits += math.Floor(float64(perf.Audience / 5))
		default:
			panic(fmt.Errorf("unknow type: %s", GetKind(play)))
		}

		// add volume credits
		volumeCredits += math.Max(float64(perf.Audience-30), 0)

		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", GetName(play), thisAmount/100, perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", volumeCredits)
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
		"hamlet":  {"Hamlet", "tragedy"},
		"as-like": {"As You Like It", "comedy"},
		"othello": {"Othello", "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
