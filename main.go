package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if a filename is provided as command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <format>")
		return
	}

	// Get the filename from command-line arguments
	format := os.Args[1]

	data, err := os.ReadFile("output.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	out := parse(string(data))
	if format == "calc" {
		fmt.Printf(
			"%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
			out.Ticker,
			out.Issuer,
			out.Quality,
			out.PER,
			out.ExpRatio,
			out.DistYield,
			out.Yield10Y,
			out.Yield5Y,
			out.Yield1Y,
			out.NHoldings,
			out.AUM,
			out.Top10,
		)
		return
	}

	fmt.Printf("%+v", out)
}

func parse(jsonText string) Output {
	var parsed Response

	err := json.Unmarshal([]byte(jsonText), &parsed)
	if err != nil {
		log.Fatal("fail unmarshal", err.Error())
	}

	var ticker string
	ticker = parsed.Data.Results.Ticker

	var issuer string
	tags := parsed.Data.Results.Tags
	for _, t := range tags {
		if t.Type == "issuer" {
			issuer = t.Label
		}
	}

	var quality string
	headers := parsed.Data.Results.Header.Groups
	for _, g := range headers {
		for _, f := range g.Fields {
			if f.Name == "overallRatingScore" {
				quality = f.Value.(string)
			}
		}
	}

	var per string
	var exp_ratio string
	var yield10y string
	var yield5y string
	var yield1y string
	var dist_yield string
	var nholdings string
	var aum string
	var top10 float64
	body := parsed.Data.Results.Body.Groups
	for _, g := range body {
		for _, sg := range g.Groups {
			for _, f := range sg.Fields {
				if f.Label != nil {
					// fmt.Printf("%v %s %v\n", *f.Label, f.Value, pointer("Price / Earnings Ratio"))
					if f.Name == "pe" {
						per = f.Value.(string)
					}
					if f.Name == "expenseRatio" {
						exp_ratio = f.Value.(string)
					}
					if f.Name == "distributionYield" {
						dist_yield = f.Value.(string)
					}
					if f.Name == "numberOfHoldings" {
						nholdings = f.Value.(string)
					}
					if f.Name == "aum" {
						aum = f.Value.(string)
					}
				}
				if f.Type != nil {
					if *f.Type == "pricePerformance" {
						yield10y = *f.Perf10YrAnnualized
						yield5y = *f.Perf5YrAnnualized
						yield1y = *f.Perf1Yr
					}
					if *f.Type == "securities" {
						if f.Data != nil {
							for _, d := range f.Data {
								if v, ok := d["weight"].(float64); ok {
									top10 += v
								}
							}
						}
					}
				}
			}
		}
	}

	return Output{
		Ticker:    ticker,
		Quality:   quality,
		PER:       per,
		ExpRatio:  exp_ratio,
		Issuer:    issuer,
		Yield10Y:  yield10y,
		Yield5Y:   yield5y,
		Yield1Y:   yield1y,
		DistYield: dist_yield,
		NHoldings: nholdings,
		AUM:       aum,
		Top10:     fmt.Sprintf("%.2f", top10*100) + "%",
	}
}
