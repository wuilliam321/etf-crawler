package crawler

import (
	"encoding/json"
	"fmt"
	"log"
)

func ToString(format string, s string) string {
	out := parse(s)
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		out.Ticker,
		out.Issuer,
		out.Quality,
		out.PER,
		out.ExpRatio,
		out.DistYield,
		out.Yield10Y,
		out.Yield5Y,
		out.Yield1Y,
		out.Ytd,
		out.NHoldings,
		out.AUM,
		out.WeightedAvgMarketCap,
		out.DailyDollarVolume,
		out.Top10,
		out.IndexTracked,
		out.Segment,
		out.TopAllocation,
		out.TopAllocationPerc,
	)
}

func parse(jsonText string) Output {
	var parsed Response

	err := json.Unmarshal([]byte(jsonText), &parsed)
	if err != nil {
		log.Fatal("fail unmarshal", err.Error())
	}

	ticker := parsed.Data.Results.Ticker

	var issuer string
	tags := parsed.Data.Results.Tags
	for _, t := range tags {
		if t.Type == "issuer" {
			issuer = t.Label
		}
	}

	var quality string
	var segment string
	headers := parsed.Data.Results.Header.Groups
	for _, g := range headers {
		for _, f := range g.Fields {
			if f.Name == "overallRatingScore" {
				quality = f.Value.(string)
			}
			if f.Name == "segment" {
				segment = f.Value.(string)
			}
		}
	}

	var per string
	var exp_ratio string
	var yield10y string
	var yield5y string
	var yield1y string
	var ytd string
	var dist_yield string
	var nholdings string
	var aum string
	var weightedAvgMarketCap string
	var dailyDollarVolume string
	var top10 float64
	var topAllocation string
	var topAllocationPerc string
	var indexTracked string
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
					if f.Name == "avgDailyDollarVolValue" {
						dailyDollarVolume = f.Value.(string)
					}
					if f.Name == "weightedAvgMarketCap" {
						weightedAvgMarketCap = f.Value.(string)
					}
					if f.Name == "underlyingIndex" {
						indexTracked = f.RawValue.(string)
					}
				}
				if f.Type != nil {
					if *f.Type == "pricePerformance" {
						yield10y = *f.Perf10YrAnnualized
						yield5y = *f.Perf5YrAnnualized
						yield1y = *f.Perf1Yr
						ytd = *f.PerfYtd
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
					if *f.Type == "sectorBreakdown" {
						if f.Data != nil {
							d := f.Data
							if len(d) > 0 {
								field := d[0]
								if v, ok := field["name"].(string); ok {
									topAllocation = v
								}
								if v, ok := field["weight"].(string); ok {
									topAllocationPerc = v
								}
							}
						}
					}
				}
			}
		}
	}

	return Output{
		Ticker:               ticker,
		Quality:              quality,
		PER:                  per,
		ExpRatio:             exp_ratio,
		Issuer:               issuer,
		Yield10Y:             yield10y,
		Yield5Y:              yield5y,
		Yield1Y:              yield1y,
		Ytd:                  ytd,
		DistYield:            dist_yield,
		NHoldings:            nholdings,
		AUM:                  aum,
		WeightedAvgMarketCap: weightedAvgMarketCap,
		DailyDollarVolume:    dailyDollarVolume,
		Top10:                fmt.Sprintf("%.2f", top10*100) + "%",
		Segment:              segment,
		TopAllocation:        topAllocation,
		TopAllocationPerc:    topAllocationPerc,
		IndexTracked:         indexTracked,
	}
}
