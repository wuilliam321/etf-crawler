package main

type Tag struct {
	TagID   int    `json:"tagId"`
	Label   string `json:"label"`
	URLPart string `json:"url_part"`
	Type    string `json:"type"`
}

type FieldDataList []FieldData

type FieldData struct {
	AsOf            string   `json:"asOf"`
	Name            string   `json:"name"`
	Weight          *float64 `json:"weight,omitempty"`
	BenchmarkWeight *float64 `json:"benchmarkWeight"`
	Shares          int      `json:"shares"`
	MarketValue     int      `json:"market_value"`
	Symbol          string   `json:"symbol"`
	URL             string   `json:"url"`
	RawWeight       *float64 `json:"rawWeight"`
}

type Field struct {
	TooltipDescription *string                  `json:"tooltipDescription,omitempty"`
	Tooltip            *string                  `json:"tooltip,omitempty"`
	Value              any                      `json:"value"`
	RawValue           any                      `json:"rawValue"`
	Label              *string                  `json:"label,omitempty"`
	Type               *string                  `json:"type"`
	Perf10YrAnnualized *string                  `json:"perf10YrAnnualized,omitempty"`
	Perf5YrAnnualized  *string                  `json:"perf5YrAnnualized,omitempty"`
	Perf1Yr            *string                  `json:"perf1Yr,omitempty"`
	PerfYtd            *string                  `json:"perfYtd,omitempty"`
	Data               []map[string]interface{} `json:"data,omitempty"`
	ShortLabel         string                   `json:"shortLabel"`
	Name               string                   `json:"name"`
	DisplayType        string                   `json:"displayType"`
}

type BodyGroup struct {
	Name   string      `json:"name"`
	Label  *string     `json:"label,omitempty"`
	Order  int         `json:"order"`
	Footer *string     `json:"footer,omitempty"`
	Groups []BodyGroup `json:"groups,omitempty"`
	Fields []Field     `json:"fields,omitempty"`
}

type HeaderGroup struct {
	Name       string  `json:"name"`
	Label      *string `json:"label,omitempty"`
	Order      int     `json:"order"`
	Footer     *string `json:"footer,omitempty"`
	Fields     []Field `json:"fields"`
	Type       string  `json:"type"`
	RenderType string  `json:"renderType"`
}

type Header struct {
	Groups []HeaderGroup `json:"groups"`
}

type Body struct {
	Groups []BodyGroup `json:"groups"`
}

type Results struct {
	ID           string `json:"_id"`
	TemplateType string `json:"templateType"`
	TemplateName string `json:"templateName"`
	Ticker       string `json:"ticker"`
	Tags         []Tag  `json:"tags"`
	Header       Header `json:"header"`
	Body         Body   `json:"body"`
}

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Results Results `json:"results"`
	} `json:"data"`
}

type Output struct {
	Ticker    string `json:"ticker"`
	Issuer    string `json:"issuer"`
	Quality   string `json:"quality"`
	PER       string `json:"per"`
	ExpRatio  string `json:"exp_ratio"`
	DistYield string `json:"distributionYield"`
	Yield10Y  string `json:"yield10y"`
	Yield5Y   string `json:"yield5y"`
	Yield1Y   string `json:"yield1y"`
	Ytd       string `json:"perf_td"`
	NHoldings string `json:"nholdings"`
	AUM       string `json:"aum"`
	Top10     string `json:"top10perc"`
}
