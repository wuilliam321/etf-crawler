package crawler_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wuilliam321/etf-crawler/pkg/crawler"
)

var jsonText = `{
    "code": 200,
    "status": "Ok",
    "copyright": "© 2024 ETF.com",
    "attributionText": "Data provided by ETF.com. © 2024 ETF.com",
    "attributionHTML": "<a href=\"http://www.etf.com\">Data provided by ETF.com. © 2024 ETF.com</a>",
    "message": "Here are the results",
    "data": {
        "count": 1,
        "results": {
            "_id": "65fa0c06ac8b0c8606d62e72",
            "templateType": "liteDesktopVersion",
            "templateName": "equityFunds",
            "ticker": "VOO",
            "tags": [
                {
                    "tagId": 80054,
                    "label": "Large Cap",
                    "url_part": "large-cap-etfs",
                    "type": "editorial"
                },
                {
                    "tagId": 80031,
                    "label": "Equity",
                    "url_part": "equity-etfs",
                    "type": "editorial"
                },
                {
                    "tagId": 80086,
                    "label": "U.S.",
                    "url_part": "us-etfs",
                    "type": "editorial"
                },
                {
                    "tagId": 80241,
                    "label": "Vanguard",
                    "url_part": "vanguard-etfs",
                    "type": "issuer"
                },
                {
                    "tagId": 83322,
                    "label": "S&P 500",
                    "url_part": "sp-500-etfs",
                    "type": "popularIndex"
                },
                {
                    "tagId": 80691,
                    "label": "Size and Style",
                    "url_part": "size-and-style-etfs",
                    "type": "category"
                },
                {
                    "tagId": 80489,
                    "label": "Broad-based",
                    "url_part": "broad-based-etfs",
                    "type": "niche"
                },
                {
                    "tagId": 83569,
                    "label": "Vanilla",
                    "url_part": "vanilla-etfs",
                    "type": "strategy"
                },
                {
                    "tagId": 80591,
                    "label": "North America",
                    "url_part": "north-america-etfs",
                    "type": "region"
                }
            ],
            "header": {
                "groups": [
                    {
                        "name": "generalData",
                        "label": null,
                        "type": "general_data",
                        "renderType": "general_data",
                        "order": 1,
                        "footer": null,
                        "fields": [
                            {
                                "tooltipDescription": "The symbol for a fund as traded on the exchage.",
                                "tooltip": "Ticker",
                                "value": "VOO",
                                "rawValue": "VOO",
                                "label": null,
                                "shortLabel": "Ticker",
                                "name": "ticker",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "ETF fund name.",
                                "tooltip": "Fund Name",
                                "value": "Vanguard 500 Index Fund",
                                "rawValue": "Vanguard 500 Index Fund",
                                "label": "Name",
                                "shortLabel": null,
                                "name": "fund",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The type of securities to which the ETF provides exposure.",
                                "tooltip": "Asset Class",
                                "value": "Equity",
                                "rawValue": "Equity",
                                "label": null,
                                "shortLabel": null,
                                "name": "assetClass",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The first level of sorting in the ETF Classification System. Categories vary by asset class.",
                                "tooltip": "Category",
                                "value": "Size and Style",
                                "rawValue": "Size and Style",
                                "label": null,
                                "shortLabel": null,
                                "name": "category",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The second level of categorization in the ETF Classification System. Essentially, a subcategory of &quot;Category&quot;.",
                                "tooltip": "Focus",
                                "value": "Large Cap",
                                "rawValue": "Large Cap",
                                "label": null,
                                "shortLabel": null,
                                "name": "focus",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The third level of categorization in the ETF Classification System. Essentially, a subcategory of &quot;Focus&quot;.",
                                "tooltip": "Niche",
                                "value": "Broad-based",
                                "rawValue": "Broad-based",
                                "label": null,
                                "shortLabel": null,
                                "name": "niche",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The third level of geographical categorization in the ETF Classification System.",
                                "tooltip": "Region",
                                "value": "North America",
                                "rawValue": "North America",
                                "label": null,
                                "shortLabel": null,
                                "name": "region",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The location of the market to which the ETF provides exposure.",
                                "tooltip": "Geography",
                                "value": "U.S.",
                                "rawValue": "U.S.",
                                "label": null,
                                "shortLabel": null,
                                "name": "geography",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The precise market to which the fund provides exposure. ETF.com assigns each fund to a segment according to a rules-based process.",
                                "tooltip": "Segment",
                                "value": "Equity: U.S.  -  Large Cap",
                                "rawValue": "Equity: U.S.  -  Large Cap",
                                "label": "ETF.com segment",
                                "shortLabel": null,
                                "name": "segment",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": 261,
                                "rawValue": 261,
                                "label": null,
                                "shortLabel": null,
                                "name": "segmentId",
                                "displayType": null
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "Yes",
                                "rawValue": true,
                                "label": null,
                                "shortLabel": null,
                                "name": "hasSegmentReport",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "FactSet Research Systems' overall rating grade for the ETF. The letter grade is calculated as the average between the efficiency and tradability scores. The number corresponds to the ETF fit score.",
                                "tooltip": "Overall Rating",
                                "value": "A",
                                "rawValue": "A",
                                "label": null,
                                "shortLabel": null,
                                "name": "letterGrade",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The best ETF choice for an average investor in a market segment.",
                                "tooltip": "Analyst Pick",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "analystPick",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "An ETF which has a unique or alternative investment approach into a market segment.",
                                "tooltip": "Opportunity List",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "opportunitiesList",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The ETF.com Efficiency Score describes how well a fund delivers on its promises to investors in the areas of costs, tracking, and risks. Some of the factors in the composite score include expense ratios, goodness of fit to benchmark, tracking difference, and a breadth of risk measurements from structural risks to tax risks and fund closure risk.",
                                "tooltip": "Efficiency",
                                "value": "99",
                                "rawValue": 99.034539,
                                "label": "E",
                                "shortLabel": null,
                                "name": "efficiencyScore",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The ETF.com Tradability Score accounts for the expense and uncertainty that an investor might encounter when buying or selling a fund in the open market. It accounts for on-screen liquidity at retail levels as well as block liquidity at institutional levels, and considers both fundlevel metrics and those of the underlying holdings.",
                                "tooltip": "Tradability",
                                "value": "100",
                                "rawValue": 99.871089,
                                "label": "T",
                                "shortLabel": null,
                                "name": "tradabilityScore",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The ETF.com Fit Score measures how well a fund's exposures match those of a neutral benchmark chosen to be most representative of the ETF.com segment.",
                                "tooltip": "Fit",
                                "value": "95",
                                "rawValue": 94.767861,
                                "label": "F",
                                "shortLabel": null,
                                "name": "fitScore",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "79",
                                "rawValue": 79.40701411692308,
                                "label": null,
                                "shortLabel": null,
                                "name": "segmentAvgEfficiency",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "79",
                                "rawValue": 79.1994028128655,
                                "label": null,
                                "shortLabel": null,
                                "name": "segmentAvgTradability",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "60",
                                "rawValue": 60.44776934594594,
                                "label": null,
                                "shortLabel": null,
                                "name": "segmentAvgFit",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "03/14/24",
                                "rawValue": "2024-03-14T00:00:00.000Z",
                                "label": "as of 03/14/24",
                                "shortLabel": null,
                                "name": "overallRatingDate",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "An ETF with an MSCI ESG Fund Quality Score ranked in the top 20% against the scores of other funds within the same peer group.",
                                "tooltip": "MSCI Badge",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "msciEsgHasBadge",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "genericRpt",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "Yes",
                                "rawValue": true,
                                "label": null,
                                "shortLabel": null,
                                "name": "msciEsgReady",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "The ratio of dividends paid by the fund over past 12 months divided by fund's share price.",
                                "tooltip": "Dividend Yield",
                                "value": "1.41%",
                                "rawValue": 0.014127,
                                "label": "Div Yield",
                                "shortLabel": null,
                                "name": "dividendYield",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": 1,
                                "rawValue": 1,
                                "label": null,
                                "shortLabel": null,
                                "name": "state",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "inverse",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "leveraged",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "A 95",
                                "rawValue": "A 94.767861",
                                "label": null,
                                "shortLabel": null,
                                "name": "overallRatingScore",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "SPY,IVV,SCHX,IWB,VV",
                                "rawValue": "SPY,IVV,SCHX,IWB,VV",
                                "label": null,
                                "shortLabel": null,
                                "name": "moreComparisons",
                                "displayType": null
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "0.00%",
                                "rawValue": 0.000047,
                                "label": "Average Spread (%)",
                                "shortLabel": null,
                                "name": "medianSpreadPct45Day",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "Data provided by FactSet.",
                                "label": null,
                                "shortLabel": null,
                                "name": "factSetAttributionText1",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "Unless otherwise stated, data provided by FactSet.",
                                "label": null,
                                "shortLabel": null,
                                "name": "factSetAttributionText2",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "NYSEArca",
                                "rawValue": "NYSEArca",
                                "label": null,
                                "shortLabel": null,
                                "name": "primaryExchange",
                                "displayType": "text"
                            }
                        ]
                    },
                    {
                        "name": "fundInfo",
                        "label": null,
                        "column": "left",
                        "type": "general_data",
                        "renderType": "general_data",
                        "order": 1,
                        "footer": null,
                        "fields": [
                            {
                                "tooltipDescription": "The symbol for a fund as traded on the exchage.",
                                "tooltip": "Ticker",
                                "value": "VOO",
                                "rawValue": "VOO",
                                "label": null,
                                "shortLabel": "Ticker",
                                "name": "ticker",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "ETF fund name.",
                                "tooltip": "Fund Name",
                                "value": "Vanguard 500 Index Fund",
                                "rawValue": "Vanguard 500 Index Fund",
                                "label": "Name",
                                "shortLabel": null,
                                "name": "fund",
                                "displayType": "text"
                            }
                        ]
                    },
                    {
                        "name": "fundBadges",
                        "label": null,
                        "column": "full",
                        "type": "general_data",
                        "renderType": "general_data",
                        "order": 1,
                        "footer": null,
                        "fields": [
                            {
                                "tooltipDescription": "The best ETF choice for an average investor in a market segment.",
                                "tooltip": "Analyst Pick",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "analystPick",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "An ETF which has a unique or alternative investment approach into a market segment.",
                                "tooltip": "Opportunity List",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "opportunitiesList",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": null,
                                "tooltip": null,
                                "value": "A 95",
                                "rawValue": "A 94.767861",
                                "label": null,
                                "shortLabel": null,
                                "name": "overallRatingScore",
                                "displayType": "text"
                            },
                            {
                                "tooltipDescription": "An ETF with an MSCI ESG Fund Quality Score ranked in the top 20% against the scores of other funds within the same peer group.",
                                "tooltip": "MSCI Badge",
                                "value": "No",
                                "rawValue": false,
                                "label": null,
                                "shortLabel": null,
                                "name": "msciEsgHasBadge",
                                "displayType": "text"
                            }
                        ]
                    },
                    {
                        "name": "lastTradeHeader",
                        "label": null,
                        "column": "full",
                        "type": "endpoint_to_query",
                        "renderType": "table",
                        "order": 2,
                        "footer": null,
                        "endPointToQueryData": {
                            "serviceName": "webSockets",
                            "event": "intradayQuotes"
                        }
                    }
                ]
            },
            "body": {
                "groups": [
                    {
                        "name": "overview",
                        "label": "Overview",
                        "order": 1,
                        "footer": null,
                        "groups": [
                            {
                                "name": "fundDescription",
                                "label": "What is VOO?",
                                "column": "left",
                                "type": "description",
                                "renderType": "description",
                                "order": 1,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "The fund is passively managed to hold large-cap US stocks selected by an S&amp;P Committee.",
                                        "rawValue": "The fund is passively managed to hold large-cap US stocks selected by an S&amp;P Committee.",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "fundDescription",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "totalReturnPrice",
                                "label": null,
                                "column": "left",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "react-price-tr-chart",
                                "jsNames": [
                                    "priceAndTRChart.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "containerId": "priceTRChartAppContainer",
                                "order": 2,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "relatedArticles",
                                "label": "Recent News",
                                "column": "left",
                                "type": "description",
                                "renderType": "description",
                                "order": 4,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "overviewLbAd1",
                                "label": null,
                                "column": "left",
                                "type": "ad728x90",
                                "renderType": "ad728x90",
                                "order": 5,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "fundAnalyticsInsight",
                                "label": "VOO Factset Analytics Insight",
                                "column": "left",
                                "type": "description",
                                "renderType": "description",
                                "order": 6,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "<p>The fund provides comprehensive exposure to the US large-cap equity space by tracking a well-known index, the S&amp;P 500. Like all S&amp;P 500 funds, the fund defines large-caps as the S&amp;P Committee sees it, which means it includes a fairly large allocation to firms that may be considered mid-caps. Stocks are market-cap-weighted, with a consideration on sector balance by comparing the weight of each GICS (Global Industry Classification Standard) sector in the underlying index to its weight in the relevant market capitalization range of the S&amp;P Total Market Index (parent index). The index is rebalanced on a quarterly basis.</p>",
                                        "rawValue": "<p>The fund provides comprehensive exposure to the US large-cap equity space by tracking a well-known index, the S&amp;P 500. Like all S&amp;P 500 funds, the fund defines large-caps as the S&amp;P Committee sees it, which means it includes a fairly large allocation to firms that may be considered mid-caps. Stocks are market-cap-weighted, with a consideration on sector balance by comparing the weight of each GICS (Global Industry Classification Standard) sector in the underlying index to its weight in the relevant market capitalization range of the S&amp;P Total Market Index (parent index). The index is rebalanced on a quarterly basis.</p>",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "overviewInsight",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "msciEsgAnalyticsInsight",
                                "label": "VOO MSCI ESG Analytics Insight",
                                "column": "left",
                                "type": "description",
                                "renderType": "description",
                                "order": 7,
                                "lock": true,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "Vanguard 500 Index Fund has an MSCI ESG Fund Rating of A based on a score of 6.59 out of 10. The MSCI ESG Fund Rating measures the resiliency of portfolios to long-term risks and opportunities arising from environmental, social, and governance factors. ESG Fund Ratings range from best (AAA) to worst (CCC). Highly rated funds consist of companies that tend to show strong and/or improving management of financially relevant environmental, social and governance issues. These companies may be more resilient to disruptions arising from ESG events.<br/><br/>The fund’s Peer Rank reflects the ranking of a fund’s MSCI ESG Fund Quality Score against the scores of other funds within the same peer group, as defined by the Thomson Reuters Lipper Global Classification. Vanguard 500 Index Fund ranks in the 56th percentile within its peer group and in the 45th percentile within the global universe of all funds covered by MSCI ESG Fund Ratings.",
                                        "rawValue": "Vanguard 500 Index Fund has an MSCI ESG Fund Rating of A based on a score of 6.59 out of 10. The MSCI ESG Fund Rating measures the resiliency of portfolios to long-term risks and opportunities arising from environmental, social, and governance factors. ESG Fund Ratings range from best (AAA) to worst (CCC). Highly rated funds consist of companies that tend to show strong and/or improving management of financially relevant environmental, social and governance issues. These companies may be more resilient to disruptions arising from ESG events.<br/><br/>The fund’s Peer Rank reflects the ranking of a fund’s MSCI ESG Fund Quality Score against the scores of other funds within the same peer group, as defined by the Thomson Reuters Lipper Global Classification. Vanguard 500 Index Fund ranks in the 56th percentile within its peer group and in the 45th percentile within the global universe of all funds covered by MSCI ESG Fund Ratings.",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "msciEsgInsight",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The MSCI ESG Fund Rating measures the resiliency of portfolios to long-term risks and opportunities arising from environmental, social, and governance factors on a AAA (leader) to CCC (laggard) scale.",
                                        "tooltip": "MSCI ESG Rating",
                                        "value": "A",
                                        "rawValue": "A",
                                        "label": "MSCI ESG Rating",
                                        "shortLabel": null,
                                        "name": "msciFundEsgRating",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "msciFaCSInsight",
                                "label": "VOO MSCI F<span class='lowercase'>a</span>CS and Factor Box",
                                "column": "left",
                                "type": "description",
                                "renderType": "description",
                                "order": 8,
                                "lock": true,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "<p>MSCI FaCS is a standard method for evaluating and reporting the Factor characteristics of equity portfolios including ETFs. The Factor Box includes 6 Factors that MSCI has identified that historically provided a return premium. On the vertical axis, the Factor Groups, are displayed and the horizontal axis displays the Factor exposure, overweight, underweight or neutral.</p>",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "msciFaCSInsight",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "overviewAd1",
                                "label": null,
                                "column": "right",
                                "type": "ad300x250",
                                "renderType": "ad300x250",
                                "order": 1,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "fundSummaryData",
                                "label": "VOO Summary Data",
                                "column": "right",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 2,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "A company that produces and manages ETFs.",
                                        "tooltip": "Issuer",
                                        "value": "<a href='/channels/vanguard-etfs'>Vanguard</a>",
                                        "rawValue": "Vanguard",
                                        "label": "Issuer",
                                        "shortLabel": null,
                                        "name": "issuer",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The first date of a fund's operations, as documented by the issuer.",
                                        "tooltip": "Inception Date",
                                        "value": "09/07/10",
                                        "rawValue": "2010-09-07T00:00:00.000Z",
                                        "label": "Inception Date",
                                        "shortLabel": null,
                                        "name": "inceptionDate",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The net total annual fee a fund holder pays to the issuer.",
                                        "tooltip": "Expense Ratio",
                                        "value": "0.03%",
                                        "rawValue": 3,
                                        "label": "Expense Ratio",
                                        "shortLabel": null,
                                        "name": "expenseRatio",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The market value of total assets that a fund has accumulated and now manages on behalf of investors.",
                                        "tooltip": "Assets Under Management",
                                        "value": "$424.79B",
                                        "rawValue": 424793506000,
                                        "label": "Assets Under Management",
                                        "shortLabel": null,
                                        "name": "aum",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "This is the benchmark an ETF is designed to track or replicate.",
                                        "tooltip": "Index Tracked",
                                        "value": "<a href='/channels/sp-500-etfs'>S&P 500</a>",
                                        "rawValue": "S&P 500",
                                        "label": "Index Tracked",
                                        "shortLabel": null,
                                        "name": "underlyingIndex",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "This is the index that we have chosen as the best-in-class gauge for each segment's broad market.",
                                        "tooltip": "Segment Benchmark",
                                        "value": "MSCI USA Large Cap Index",
                                        "rawValue": "MSCI USA Large Cap Index",
                                        "label": "Segment Benchmark",
                                        "shortLabel": null,
                                        "name": "segmentBenchmark",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The organizational structure of the fund or ETN.",
                                        "tooltip": "Legal Structure",
                                        "value": "Open-Ended Fund",
                                        "rawValue": "Open-Ended Fund",
                                        "label": "Legal Structure",
                                        "shortLabel": null,
                                        "name": "legalStructure",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "https://investor.vanguard.com/investment-products/etfs/profile/voo",
                                        "rawValue": "https://investor.vanguard.com/investment-products/etfs/profile/voo",
                                        "label": "Fund Home Page",
                                        "shortLabel": null,
                                        "name": "fundHomepage",
                                        "displayType": "link"
                                    }
                                ]
                            },
                            {
                                "name": "overviewAd2",
                                "label": null,
                                "column": "right",
                                "type": "ad300x250",
                                "renderType": "ad300x250",
                                "order": 3,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "peers",
                                "label": "Peers",
                                "column": "right",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 4,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "This is the index that we have chosen as the best-in-class gauge for each segment's broad market.",
                                        "tooltip": "Segment Benchmark",
                                        "value": "MSCI USA Large Cap Index",
                                        "rawValue": "MSCI USA Large Cap Index",
                                        "label": "Segment Benchmark",
                                        "shortLabel": null,
                                        "name": "segmentBenchmark",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "ETFs from within the same segment or closely related segments with similar investment objectives or market exposures.",
                                        "tooltip": "Competing ETFs",
                                        "value": "SPY, IVV, SCHX, IWB, VV",
                                        "rawValue": "SPY, IVV, SCHX, IWB, VV",
                                        "label": "Competing ETFs",
                                        "shortLabel": null,
                                        "name": "competingETFs",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "SPY,IVV,SCHX,IWB,VV",
                                        "rawValue": "SPY,IVV,SCHX,IWB,VV",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "moreComparisons",
                                        "displayType": null
                                    }
                                ]
                            },
                            {
                                "name": "overviewAd3",
                                "label": null,
                                "column": "right",
                                "type": "ad300x250",
                                "renderType": "ad300x250",
                                "order": 5,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "fundFlows",
                                "label": "Fund Flows",
                                "column": "right",
                                "type": "general_data",
                                "renderType": "fund_flows",
                                "order": 6,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "Aggregate 5 Days Net Flows in USD Billions.",
                                        "tooltip": "5 Days",
                                        "value": "$3.27B",
                                        "rawValue": 3274.7757699999997,
                                        "label": "5 Days",
                                        "shortLabel": "5 Days",
                                        "name": "netFlows5Days",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Aggregate 30 Days Net Flows in USD Billions.",
                                        "tooltip": "30 Days",
                                        "value": "$12.62B",
                                        "rawValue": 12615.5538024,
                                        "label": "30 Days",
                                        "shortLabel": "30 Days",
                                        "name": "netFlows30Days",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Aggregate 90 Days Net Flows in USD Billions.",
                                        "tooltip": "90 Days",
                                        "value": "$29.35B",
                                        "rawValue": 29346.2018779,
                                        "label": "90 Days",
                                        "shortLabel": "90 Days",
                                        "name": "netFlows90Days",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "relatedETFChannels",
                                "label": "Related ETFs Lists & Channels",
                                "column": "right",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 7,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "<a href='/channels/large-cap-etfs'>Large Cap</a>, <a href='/channels/equity-etfs'>Equity</a>, <a href='/channels/us-etfs'>U.S.</a>, <a href='/channels/sp-500-etfs'>S&P 500</a>, <a href='/channels/size-and-style-etfs'>Size and Style</a>, <a href='/channels/broad-based-etfs'>Broad-based</a>, <a href='/channels/vanilla-etfs'>Vanilla</a>, <a href='/channels/north-america-etfs'>North America</a>",
                                        "rawValue": "<a href='/channels/large-cap-etfs'>Large Cap</a>, <a href='/channels/equity-etfs'>Equity</a>, <a href='/channels/us-etfs'>U.S.</a>, <a href='/channels/sp-500-etfs'>S&P 500</a>, <a href='/channels/size-and-style-etfs'>Size and Style</a>, <a href='/channels/broad-based-etfs'>Broad-based</a>, <a href='/channels/vanilla-etfs'>Vanilla</a>, <a href='/channels/north-america-etfs'>North America</a>",
                                        "label": "Related ETF Channels",
                                        "shortLabel": null,
                                        "name": "relatedChannels",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundPortfolioData",
                                "label": "VOO Portfolio Data",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 1,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "Describes the average of each stock's market cap (share price x shares outstanding) scaled by its weight in the portfolio.",
                                        "tooltip": "Weighted Average Market Cap",
                                        "value": "$742.84B",
                                        "rawValue": 742843826000,
                                        "label": "Weighted Average Market Cap",
                                        "shortLabel": "Average Market Cap",
                                        "name": "weightedAvgMarketCap",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Weighted average ratio of prices of a fund’s stocks to trailing earnings of underlying stocks.",
                                        "tooltip": "Price / Earnings Ratio",
                                        "value": "23.90",
                                        "rawValue": 23.895711,
                                        "label": "Price / Earnings Ratio",
                                        "shortLabel": "P/E",
                                        "name": "pe",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Weighted average ratio of prices of a fund’s stocks to the book value of underlying equity.",
                                        "tooltip": "Price / Book Ratio",
                                        "value": "4.46",
                                        "rawValue": 4.456846,
                                        "label": "Price / Book Ratio",
                                        "shortLabel": "P/B",
                                        "name": "pb",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The ratio of distributions paid by the fund over the past 12 months, divided by the fund’s NAV.",
                                        "tooltip": "Distribution Yield",
                                        "value": "1.34%",
                                        "rawValue": 0.013441,
                                        "label": "Distribution Yield",
                                        "shortLabel": null,
                                        "name": "distributionYield",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The date on which a security's price excludes an upcoming dividend.",
                                        "tooltip": "Next Ex-Dividend Date",
                                        "value": "N/A",
                                        "rawValue": null,
                                        "label": "Next Ex-Dividend Date",
                                        "shortLabel": null,
                                        "name": "exDivDate",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The number of securities held in the fund as of ETF.com’s analysis date, based on issuer portfolios or the creation basket. If an ETF holds other ETFs, we count every constituent, looking through the ETF wrapper.",
                                        "tooltip": "Number of Holdings",
                                        "value": "505",
                                        "rawValue": 505,
                                        "label": "Number of Holdings",
                                        "shortLabel": null,
                                        "name": "numberOfHoldings",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundIndexData",
                                "label": "VOO Index Data",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 2,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "This is the benchmark an ETF is designed to track or replicate.",
                                        "tooltip": "Index Tracked",
                                        "value": "<a href='/channels/sp-500-etfs'>S&P 500</a>",
                                        "rawValue": "S&P 500",
                                        "label": "Index Tracked",
                                        "shortLabel": null,
                                        "name": "underlyingIndex",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "A set of rules that the underlying index provider follows to weight its constituent securities.",
                                        "tooltip": "Index Weighting Methodology",
                                        "value": "Market Cap",
                                        "rawValue": "Market Cap",
                                        "label": "Index Weighting Methodology",
                                        "shortLabel": "Weighting",
                                        "name": "weightingScheme",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "A set of rules that the underlying index provider follows to select its constituent securities.",
                                        "tooltip": "Index Selection Methodology",
                                        "value": "Committee",
                                        "rawValue": "Committee",
                                        "label": "Index Selection Methodology",
                                        "shortLabel": "Selection",
                                        "name": "selectionCriteria",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "This is the index that we have chosen as the best-in-class gauge for each segment's broad market.",
                                        "tooltip": "Segment Benchmark",
                                        "value": "MSCI USA Large Cap Index",
                                        "rawValue": "MSCI USA Large Cap Index",
                                        "label": "Segment Benchmark",
                                        "shortLabel": null,
                                        "name": "segmentBenchmark",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "overviewAd4",
                                "label": null,
                                "column": "full",
                                "type": "ad300x250",
                                "renderType": "ad300x250",
                                "order": 3,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "allHoldings",
                                "label": null,
                                "column": "full",
                                "type": "endpoint_to_query",
                                "renderType": "modal_table",
                                "order": 6,
                                "footer": null,
                                "endPointToQueryData": {
                                    "serviceName": "API",
                                    "verb": "GET",
                                    "url": "/private/fund/VOO/holdings?type=securities&formatValues=true&",
                                    "body": null,
                                    "headers": {
                                        "x-limit": 10000
                                    }
                                }
                            },
                            {
                                "name": "fundPerformance",
                                "label": "Performance Dispersion",
                                "column": "left",
                                "type": "table",
                                "renderType": "table",
                                "order": 3,
                                "footer": "All returns over 1 year are annualized. All returns are total returns unless otherwise stated",
                                "fields": [
                                    {
                                        "ticker": "PERFORMANCE <span class='asOf'>[as of 03/18/24]</span>",
                                        "perf1Mo": "1 MONTH",
                                        "perf3Mo": "3 MONTHS",
                                        "perfYtd": "YTD",
                                        "perf1Yr": "1 YEAR",
                                        "perf3YrAnnualized": "3 YEARS",
                                        "perf5YrAnnualized": "5 YEARS",
                                        "perf10YrAnnualized": "10 YEARS",
                                        "type": "header"
                                    },
                                    {
                                        "ticker": "VOO",
                                        "perf1Mo": "3.03%",
                                        "perf3Mo": "9.04%",
                                        "perfYtd": "8.28%",
                                        "perf1Yr": "33.53%",
                                        "perf3YrAnnualized": "11.23%",
                                        "perf5YrAnnualized": "14.55%",
                                        "perf10YrAnnualized": "12.70%",
                                        "type": "pricePerformance"
                                    },
                                    {
                                        "ticker": "VOO (NAV)",
                                        "perf1Mo": "1.88%",
                                        "perf3Mo": "8.83%",
                                        "perfYtd": "7.62%",
                                        "perf1Yr": "33.55%",
                                        "perf3YrAnnualized": "10.48%",
                                        "perf5YrAnnualized": "14.49%",
                                        "perf10YrAnnualized": "12.81%",
                                        "type": "navPerformance"
                                    },
                                    {
                                        "ticker": "<a href='/channels/sp-500-etfs'>S&P 500</a>",
                                        "perf1Mo": "1.88%",
                                        "perf3Mo": "8.83%",
                                        "perfYtd": "7.63%",
                                        "perf1Yr": "33.59%",
                                        "perf3YrAnnualized": "10.53%",
                                        "perf5YrAnnualized": "14.53%",
                                        "perf10YrAnnualized": "12.85%",
                                        "type": "indexPerformance"
                                    },
                                    {
                                        "ticker": "MSCI USA Large Cap Index",
                                        "perf1Mo": "1.59%",
                                        "perf3Mo": "8.99%",
                                        "perfYtd": "7.84%",
                                        "perf1Yr": "35.48%",
                                        "perf3YrAnnualized": "10.61%",
                                        "perf5YrAnnualized": "14.98%",
                                        "perf10YrAnnualized": "13.21%",
                                        "type": "segmentBenchmark"
                                    }
                                ]
                            },
                            {
                                "name": "msciFaCSChart",
                                "label": null,
                                "column": "left",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "msci",
                                "jsNames": [
                                    "msciFacs.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "cssNames": [
                                    "main.chunk.css"
                                ],
                                "containerId": "msciFactorBoxApp",
                                "order": 9,
                                "lock": true,
                                "footer": null,
                                "fields": [
                                    {
                                        "name": "msciFacsValue",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": -0.00279,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Value",
                                        "shortDesc": "Relatively Inexpensive Stocks",
                                        "description": "This factor captures excess returns to stocks that have low prices relative to their fundamental value.",
                                        "order": 0
                                    },
                                    {
                                        "name": "msciFacsSize",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": -0.29877,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Low Size",
                                        "shortDesc": "Smaller Companies",
                                        "description": "This factor captures excess returns of smaller firms (by market capitalization) relative to their larger counterparts.",
                                        "order": 1
                                    },
                                    {
                                        "name": "msciFacsMomentum",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": 0.07423,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Momentum",
                                        "shortDesc": "Rising Stocks",
                                        "description": "This factor reflects excess returns to stocks with stronger past performance.",
                                        "order": 2
                                    },
                                    {
                                        "name": "msciFacsQuality",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": 0.12342,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Quality",
                                        "shortDesc": "Sound Balance Sheet Stocks",
                                        "description": "This factor captures excess returns to stocks that are characterized by low debt, stable earnings growth, and other “quality” metrics.",
                                        "order": 3
                                    },
                                    {
                                        "name": "msciFacsYild",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": 0.02776,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Yield",
                                        "shortDesc": "Cash Flow Paid Out",
                                        "description": "This factor captures excess returns to stocks that have higher-than-average dividend yields.",
                                        "order": 4
                                    },
                                    {
                                        "name": "msciFacsVolatil",
                                        "value": "Neutral Weight Factor",
                                        "rawValue": -0.21372,
                                        "similarETFs": [
                                            "570100",
                                            "AAA",
                                            "AAAU",
                                            "AADR",
                                            "AAIT"
                                        ],
                                        "label": "Low Volatility",
                                        "shortDesc": "Lower Risk Stocks",
                                        "description": "This factor captures excess returns to stocks with lower than average volatility, beta, and/or idiosyncratic risk.",
                                        "order": 5
                                    }
                                ]
                            },
                            {
                                "name": "top10Countries",
                                "label": null,
                                "column": "full",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "react-donut-chart",
                                "jsNames": [
                                    "donutChart.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "cssNames": [
                                    "main.chunk.css"
                                ],
                                "containerId": "top10CountriesAppContainer",
                                "order": 4,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "countryBreakdown",
                                        "data": [
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "United States",
                                                "weight": 1,
                                                "benchmarkWeight": 0.99894966
                                            },
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "Canada",
                                                "weight": null,
                                                "benchmarkWeight": 0.00105032
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "top10Sectors",
                                "label": null,
                                "column": "full",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "react-donut-chart",
                                "jsNames": [
                                    "donutChart.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "cssNames": [
                                    "main.chunk.css"
                                ],
                                "containerId": "top10SectorsAppContainer",
                                "order": 5,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "sectorBreakdown",
                                        "data": [
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Technology Services",
                                                "weight": 0.20089252076290498,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Electronic Technology",
                                                "weight": 0.17912308505147378,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Finance",
                                                "weight": 0.1244008551773917,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Health Technology",
                                                "weight": 0.09459692903928167,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Retail Trade",
                                                "weight": 0.07810246071777315,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Consumer Non-Durables",
                                                "weight": 0.046087536519607085,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Producer Manufacturing",
                                                "weight": 0.03658774583309064,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Consumer Services",
                                                "weight": 0.03451999930167543,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Energy Minerals",
                                                "weight": 0.03188042504642332,
                                                "benchmarkWeight": null
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Commercial Services",
                                                "weight": 0.029018605018009395,
                                                "benchmarkWeight": null
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "top10Holdings",
                                "label": null,
                                "column": "full",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "react-donut-chart",
                                "jsNames": [
                                    "donutChart.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "cssNames": [
                                    "main.chunk.css"
                                ],
                                "containerId": "top10HoldingsAppContainer",
                                "order": 7,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "securities",
                                        "data": [
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Apple Inc.",
                                                "weight": 0.07017778847905504,
                                                "benchmarkWeight": null,
                                                "stockId": "MH33D6-R",
                                                "shares": 356991624,
                                                "market_value": 68731597369,
                                                "symbol": "AAPL",
                                                "url": "/stock/AAPL"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Microsoft Corporation",
                                                "weight": 0.06968211713739676,
                                                "benchmarkWeight": null,
                                                "stockId": "P8R3C2-R",
                                                "shares": 181486386,
                                                "market_value": 68246140591,
                                                "symbol": "MSFT",
                                                "url": "/stock/MSFT"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Amazon.com, Inc.",
                                                "weight": 0.03445021535058239,
                                                "benchmarkWeight": null,
                                                "stockId": "MCNYYL-R",
                                                "shares": 222063192,
                                                "market_value": 33740281392,
                                                "symbol": "AMZN",
                                                "url": "/stock/AMZN"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "NVIDIA Corporation",
                                                "weight": 0.03049731464380184,
                                                "benchmarkWeight": null,
                                                "stockId": "K7TPSX-R",
                                                "shares": 60314285,
                                                "market_value": 29868840218,
                                                "symbol": "NVDA",
                                                "url": "/stock/NVDA"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Alphabet Inc. Class A",
                                                "weight": 0.020611213934645968,
                                                "benchmarkWeight": null,
                                                "stockId": "HTM0LK-R",
                                                "shares": 144509039,
                                                "market_value": 20186467658,
                                                "symbol": "GOOGL",
                                                "url": "/stock/GOOGL"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Meta Platforms Inc Class A",
                                                "weight": 0.01958830104321579,
                                                "benchmarkWeight": null,
                                                "stockId": "QLGSL2-R",
                                                "shares": 54200007,
                                                "market_value": 19184634478,
                                                "symbol": "META",
                                                "url": "/stock/META"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Alphabet Inc. Class C",
                                                "weight": 0.01750126434504762,
                                                "benchmarkWeight": null,
                                                "stockId": "WFJYTJ-R",
                                                "shares": 121624966,
                                                "market_value": 17140606458,
                                                "symbol": "GOOG",
                                                "url": "/stock/GOOG"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Tesla, Inc.",
                                                "weight": 0.017134126006526248,
                                                "benchmarkWeight": null,
                                                "stockId": "Q2YN1N-R",
                                                "shares": 67534747,
                                                "market_value": 16781033935,
                                                "symbol": "TSLA",
                                                "url": "/stock/TSLA"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Berkshire Hathaway Inc. Class B",
                                                "weight": 0.01617923914987008,
                                                "benchmarkWeight": null,
                                                "stockId": "DBNXVB-R",
                                                "shares": 44428377,
                                                "market_value": 15845824941,
                                                "symbol": "BRK.B",
                                                "url": "/stock/BRK.B"
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "JPMorgan Chase & Co.",
                                                "weight": 0.012260770832703658,
                                                "benchmarkWeight": null,
                                                "stockId": "NNKD2Y-R",
                                                "shares": 70594395,
                                                "market_value": 12008106590,
                                                "symbol": "JPM",
                                                "url": "/stock/JPM"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "efficiency",
                        "label": "Efficiency",
                        "order": 2,
                        "footer": null,
                        "groups": [
                            {
                                "name": "fundPortfolioManData",
                                "label": "VOO Portfolio Management",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 1,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "The net total annual fee a fund holder pays to the issuer.",
                                        "tooltip": "Expense Ratio",
                                        "value": "0.03%",
                                        "rawValue": 3,
                                        "label": "Expense Ratio",
                                        "shortLabel": null,
                                        "name": "expenseRatio",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Compares returns of the fund's NAV to its underlying index for a daily series of overlapping 12 month periods. The median is the middle value of the results.",
                                        "tooltip": "Median Tracking Difference (12 Mo)",
                                        "value": "-0.05%",
                                        "rawValue": -0.000451,
                                        "label": "Median Tracking Difference (12 Mo)",
                                        "shortLabel": "Median",
                                        "name": "medianTrackingDifference12Mo",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Largest deviation in a positive direction of a fund's returns vs. its underlying index over the past 12 months.",
                                        "tooltip": "Max. Upside Deviation (12 Mo)",
                                        "value": "0.00%",
                                        "rawValue": 0.000027,
                                        "label": "Max. Upside Deviation (12 Mo)",
                                        "shortLabel": "Max. Upside",
                                        "name": "maxUpsideDeviation",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Largest deviation in a negative direction of a fund's returns vs. its underlying index over the past 12 months.",
                                        "tooltip": "Max. Downside Deviation (12 Mo)",
                                        "value": "-0.06%",
                                        "rawValue": -0.000586,
                                        "label": "Max. Downside Deviation (12 Mo)",
                                        "shortLabel": "Max. Downside",
                                        "name": "maxDownsideDeviation12Mo",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundTaxExposuresData",
                                "label": "VOO Tax Exposures",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 2,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "The maximum long-term and short-term U.S. tax rates applicable to a realized capital gain.",
                                        "tooltip": "Max LT/ST Capital Gains Rate",
                                        "value": "20.00% / 39.60%",
                                        "rawValue": "20 / 39.6",
                                        "label": "Max LT/ST Capital Gains Rate",
                                        "shortLabel": "Max LT/ST Cap Gains Rate Rule",
                                        "name": "maxLTSTCapitalGainsRate",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The average capital gains paid out to shareholders in the past 36 months, measured as a percent of net asset value (NAV) at the time.",
                                        "tooltip": "Capital Gains Distributions (3 Year)",
                                        "value": "--",
                                        "rawValue": 0,
                                        "label": "Capital Gains Distributions (3 Year)",
                                        "shortLabel": null,
                                        "name": "avgCapGainsPayout3Yr",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "IRS treatment of the majority of the fund's distributions.",
                                        "tooltip": "Tax on Distributions",
                                        "value": "Qualified dividends",
                                        "rawValue": "Qualified dividends",
                                        "label": "Tax on Distributions",
                                        "shortLabel": null,
                                        "name": "taxOnDist",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "ETFs that are structured as commodities pools and classified as limited partnerships by the IRS will issue K-1 forms to holders.",
                                        "tooltip": "Distributes K1",
                                        "value": "No",
                                        "rawValue": false,
                                        "label": "Distributes K1",
                                        "shortLabel": null,
                                        "name": "distributesK1",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundStructureData",
                                "label": "VOO Fund Structure",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 3,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "The organizational structure of the fund or ETN.",
                                        "tooltip": "Legal Structure",
                                        "value": "Open-Ended Fund",
                                        "rawValue": "Open-Ended Fund",
                                        "label": "Legal Structure",
                                        "shortLabel": null,
                                        "name": "legalStructure",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "An indicator of whether a fund uses over-the-counter derivatives such as swaps or forwards to achieve its objectives.",
                                        "tooltip": "OTC Derivative Use",
                                        "value": "No",
                                        "rawValue": false,
                                        "label": "OTC Derivative Use",
                                        "shortLabel": null,
                                        "name": "derivatives",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "An indicator of whether or not a fund actively lends its portfolio holdings.",
                                        "tooltip": "Securities Lending Active",
                                        "value": "Yes",
                                        "rawValue": 1,
                                        "label": "Securities Lending Active",
                                        "shortLabel": null,
                                        "name": "secLendingActive",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The division of securities lending revenue between the fund and the issuer.",
                                        "tooltip": "Securities Lending Split (Fund/Issuer)",
                                        "value": "-- / --",
                                        "rawValue": " / ",
                                        "label": "Securities Lending Split (Fund/Issuer)",
                                        "shortLabel": null,
                                        "name": "secLendingSplitFundIssuer",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The institution responsible for providing the return of the stated index and whose credit is the sole backing of the ETN.",
                                        "tooltip": "ETN Counterparty",
                                        "value": "N/A",
                                        "rawValue": "N/A",
                                        "label": "ETN Counterparty",
                                        "shortLabel": null,
                                        "name": "etnCreditCounterparty",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The risk of default by the ETN counterparty.",
                                        "tooltip": "ETN Counterparty Risk",
                                        "value": "N/A",
                                        "rawValue": "N/A",
                                        "label": "ETN Counterparty Risk",
                                        "shortLabel": null,
                                        "name": "etnCreditCounterpartyRisk",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The likelihood that an issuer will shut down a fund for business or regulatory reasons.",
                                        "tooltip": "Fund Closure Risk",
                                        "value": "Low",
                                        "rawValue": "Low",
                                        "label": "Fund Closure Risk",
                                        "shortLabel": null,
                                        "name": "fundClosureRisk",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The frequency of an issuer's disclosure of all fund holdings.",
                                        "tooltip": "Portfolio Disclosure",
                                        "value": "Monthly",
                                        "rawValue": "Monthly",
                                        "label": "Portfolio Disclosure",
                                        "shortLabel": null,
                                        "name": "portfolioDisclosure",
                                        "displayType": "text"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "tradability",
                        "label": "Tradability",
                        "order": 3,
                        "footer": null,
                        "groups": [
                            {
                                "name": "quotesAndBookViewer",
                                "label": null,
                                "column": "left",
                                "type": "app",
                                "renderType": "react",
                                "s3Folder": "react-book-viewer-app",
                                "jsNames": [
                                    "book_viewer_app.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "containerId": "bookViewerContainer",
                                "order": 1,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "fundTradabilityChart",
                                "label": null,
                                "column": "left",
                                "type": "chart",
                                "renderType": "react",
                                "s3Folder": "react-tradability-chart",
                                "jsNames": [
                                    "tradabilityChart.js",
                                    "2.chunk.js",
                                    "main.chunk.js"
                                ],
                                "containerId": "tradChartAppContainer",
                                "order": 2,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "tradabilityAd",
                                "label": null,
                                "column": "right",
                                "type": "ad300x250",
                                "renderType": "ad300x250",
                                "order": 1,
                                "footer": null,
                                "fields": []
                            },
                            {
                                "name": "fundAnalyticsBLData",
                                "label": "VOO Factset Analytics Block Liquidity",
                                "column": "right",
                                "type": "block_liquidity",
                                "renderType": "block_liquidity",
                                "order": 2,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "03/19/24",
                                        "rawValue": "2024-03-19T00:00:00.000Z",
                                        "label": null,
                                        "shortLabel": null,
                                        "name": "asOf",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "An estimate of liquidity for the underlying baskets of securities, scaled 1 to 5.",
                                        "tooltip": "ETF.com Implied Liquidity",
                                        "value": 5,
                                        "rawValue": 5,
                                        "label": "ETF.com Implied Liquidity",
                                        "shortLabel": null,
                                        "name": "impliedLiquidity",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundTradabilityData",
                                "label": "VOO Tradability",
                                "column": "right",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 3,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "It is the daily number of shares traded, averaged over the past 45 trading days.",
                                        "tooltip": "Avg. Daily Share Volume",
                                        "value": "4,742,489",
                                        "rawValue": 4742489.422222,
                                        "label": "Avg. Daily Share Volume",
                                        "shortLabel": null,
                                        "name": "avgDailyShareVolume",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "It is the daily dollar value of shares traded, averaged over the past 45 trading days.",
                                        "tooltip": "Average Daily $ Volume",
                                        "value": "$2.17B",
                                        "rawValue": 2166957102.971111,
                                        "label": "Average Daily $ Volume",
                                        "shortLabel": null,
                                        "name": "avgDailyDollarVolValue",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Median number of shares traded over the past 45 trading days.",
                                        "tooltip": "Median Daily Share Volume",
                                        "value": "4,502,089",
                                        "rawValue": 4502089,
                                        "label": "Median Daily Share Volume",
                                        "shortLabel": null,
                                        "name": "medianDailyShareVol45Day",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "Median dollar value of shares traded over the past 45 trading days.",
                                        "tooltip": "Median Daily Volume ($)",
                                        "value": "$2.05B",
                                        "rawValue": 2053548850.86,
                                        "label": "Median Daily Volume ($)",
                                        "shortLabel": null,
                                        "name": "medianDailyDollarVol45Day",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The difference between the highest price a market participant is willing to pay to buy an ETF and the lowest price at which a market participant is willing to sell an ETF, averaged over the past 45 days, as a percent.",
                                        "tooltip": "Average Spread (%)",
                                        "value": "0.00%",
                                        "rawValue": 0.000047,
                                        "label": "Average Spread (%)",
                                        "shortLabel": null,
                                        "name": "medianSpreadPct45Day",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The difference between the highest and lowest posted prices for an ETF, averaged over the past 45 days, in dollars.",
                                        "tooltip": "Average Spread ($)",
                                        "value": "$0.02",
                                        "rawValue": 0.021532,
                                        "label": "Average Spread ($)",
                                        "shortLabel": null,
                                        "name": "medianSpreadUsd45Day",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The middle value in the ranked set of all premium/discount values over a maximum 12-month period.",
                                        "tooltip": "Median Premium / Discount (12 Mo)",
                                        "value": "0.01%",
                                        "rawValue": 0.0001,
                                        "label": "Median Premium / Discount (12 Mo)",
                                        "shortLabel": null,
                                        "name": "medianPremiumDiscount12Mo",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The greatest amount that the market price exceeded (premium) and fell below (discount) its fair value/net asset value (NAV) over a maximum 12-month period.",
                                        "tooltip": "Max. Premium / Discount (12 Mo)",
                                        "value": "0.08% / -0.15%",
                                        "rawValue": "0.000845 / -0.001518",
                                        "label": "Max. Premium / Discount (12 Mo)",
                                        "shortLabel": null,
                                        "name": "maxPremiumDiscount",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "This flags whether there is currently a systemic issue that has restricted the ability to create or redeem shares of the fund. This may be imposed by the fund's issuer, or by external circumstances.",
                                        "tooltip": "Impediment to Creations",
                                        "value": "None",
                                        "rawValue": "None",
                                        "label": "Impediment to Creations",
                                        "shortLabel": null,
                                        "name": "impedimentsToCreation",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percent of time that the underlying securities of an ETF are open to trading while US exchanges are open.",
                                        "tooltip": "Market Hours Overlap",
                                        "value": "100.00%",
                                        "rawValue": 100,
                                        "label": "Market Hours Overlap",
                                        "shortLabel": null,
                                        "name": "mhoValue",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The smallest block of ETF shares that an Authorized Participant can either create or redeem at net asset value (NAV) with the issuer in exchange for the underlying shares of the fund.",
                                        "tooltip": "Creation Unit Size (Shares)",
                                        "value": "25,000",
                                        "rawValue": 25000,
                                        "label": "Creation Unit Size (Shares)",
                                        "shortLabel": null,
                                        "name": "apTransactionSize",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The median 45 day share volume divided by the creation unit size of the fund. The higher the number, the more likely that liquidity providers will trade the fund in size, or in odd lots.",
                                        "tooltip": "Creation Unit/Day (45 Day Average)",
                                        "value": "180.08",
                                        "rawValue": 180.08356,
                                        "label": "Creation Unit/Day (45 Day Average)",
                                        "shortLabel": "Creation Unit/Day",
                                        "name": "creationUnits45DayMedian",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The standard fee to create or redeem 1 creation unit of an ETF as a percentage of the dollar value of 1 creation unit.",
                                        "tooltip": "Creation Unit Cost (%)",
                                        "value": "0.01%",
                                        "rawValue": 0.00005,
                                        "label": "Creation Unit Cost (%)",
                                        "shortLabel": "Creation Cost Per Unit (%)",
                                        "name": "creationUnitCostPct",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percentage of the median daily volume in underlying securities represented by one creation unit.",
                                        "tooltip": "Underlying Volume / Unit",
                                        "value": "0.01%",
                                        "rawValue": 0.00006,
                                        "label": "Underlying Volume / Unit",
                                        "shortLabel": "Underlying Volume/Unit %",
                                        "name": "underlyingVolPerUnitPctValue",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The total number of net outstanding options contracts for an ETF.",
                                        "tooltip": "Open Interest on ETF Options",
                                        "value": "51,573",
                                        "rawValue": 51573,
                                        "label": "Open Interest on ETF Options",
                                        "shortLabel": null,
                                        "name": "openInterest",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The total market value of the assets that an ETF holds less fund expenses.",
                                        "tooltip": "Net Asset Value (Yesterday)",
                                        "value": "$469.90",
                                        "rawValue": 469.9,
                                        "label": "Net Asset Value (Yesterday)",
                                        "shortLabel": null,
                                        "name": "lastNavValue",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "An estimate of liquidity for the underlying baskets of securities, scaled 1 to 5.",
                                        "tooltip": "ETF.com Implied Liquidity",
                                        "value": 5,
                                        "rawValue": 5,
                                        "label": "ETF.com Implied Liquidity",
                                        "shortLabel": null,
                                        "name": "impliedLiquidity",
                                        "displayType": "text"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "fit",
                        "label": "Fit",
                        "order": 4,
                        "footer": null,
                        "groups": [
                            {
                                "name": "allSectorIndustryBreakdown",
                                "label": "ALL SECTORS FOR VOO",
                                "column": "full",
                                "showTitle": true,
                                "type": "endpoint_to_query",
                                "renderType": "general_data_type2",
                                "order": 1,
                                "footer": null,
                                "endPointToQueryData": {
                                    "serviceName": "API",
                                    "verb": "GET",
                                    "url": "/private/fund/VOO/holdings?type=sectorBreakdown&formatValues=true&",
                                    "body": null,
                                    "headers": {
                                        "x-limit": 10000
                                    }
                                }
                            },
                            {
                                "name": "allHoldings",
                                "label": "ALL HOLDINGS FROM VOO",
                                "column": "full",
                                "showTitle": false,
                                "type": "endpoint_to_query",
                                "renderType": "general_data_type1",
                                "order": 3,
                                "footer": null,
                                "endPointToQueryData": {
                                    "serviceName": "API",
                                    "verb": "GET",
                                    "url": "/private/fund/VOO/holdings?type=securities&formatValues=true&",
                                    "body": null,
                                    "headers": {
                                        "x-limit": 10000
                                    }
                                }
                            },
                            {
                                "name": "fundPerformanceStatsData",
                                "label": "VOO Performance Statistics",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 5,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "The degree to which the fund and its segment benchmark move up and down in unison.",
                                        "tooltip": "Goodness of Fit (R2)",
                                        "value": "0.99",
                                        "rawValue": 0.994461,
                                        "label": "Goodness of Fit (R2)",
                                        "shortLabel": "Goodness of Fit R<sup>2</sup>",
                                        "name": "goodnessOfFit",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The sensitivity of the returns of the fund to the movement of the ETF.com segment benchmark. Beta of 1.0 means magnitude of fund returns equals that of IU benchmark returns.",
                                        "tooltip": "Beta",
                                        "value": "0.98",
                                        "rawValue": 0.981446,
                                        "label": "Beta",
                                        "shortLabel": null,
                                        "name": "benchmarkBeta12Mo",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The comparison of a fund's return to our benchmark's for days when the benchmark is up. Ideally down beta is less than up beta while beta of 1.0 means they're equal.",
                                        "tooltip": "Up Beta",
                                        "value": "0.98",
                                        "rawValue": 0.97831,
                                        "label": "Up Beta",
                                        "shortLabel": null,
                                        "name": "benchmarkBeta12MoUp",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The comparison of a fund's return to our benchmark's for days when the benchmark is down. Ideally down beta is less than up beta while beta of 1.0 means they're equal.",
                                        "tooltip": "Down Beta",
                                        "value": "0.98",
                                        "rawValue": 0.976929,
                                        "label": "Down Beta",
                                        "shortLabel": null,
                                        "name": "benchmarkBeta12MoDown",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "A measure of the variability between the fund's returns and the ETF.com segment benchmark returns on days when the fund underperforms the benchmark.",
                                        "tooltip": "Downside Standard Deviation",
                                        "value": "0.06%",
                                        "rawValue": 0.000561,
                                        "label": "Downside Standard Deviation",
                                        "shortLabel": null,
                                        "name": "downsideStdDeviation",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "This is the index that we have chosen as the best-in-class gauge for each segment's broad market.",
                                        "tooltip": "Segment Benchmark",
                                        "value": "MSCI USA Large Cap Index",
                                        "rawValue": "MSCI USA Large Cap Index",
                                        "label": "Segment Benchmark",
                                        "shortLabel": null,
                                        "name": "segmentBenchmark",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundMSCIESGRatings",
                                "label": "VOO MSCI ESG Ratings",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 9,
                                "lock": true,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": "The MSCI ESG Fund Rating measures the resiliency of portfolios to long-term risks and opportunities arising from environmental, social, and governance factors on a AAA (leader) to CCC (laggard) scale.",
                                        "tooltip": "MSCI ESG Rating",
                                        "value": "A",
                                        "rawValue": "A",
                                        "label": "MSCI ESG Rating",
                                        "shortLabel": null,
                                        "name": "msciFundEsgRating",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The weighted average of the underlying holding’s Overall ESG Scores. The score aims to measure the ability of the underlying holdings to manage key medium- to long-term ESG risks and opportunities. The factor is provided on a 0‐10 score, 10 being the highest possible score.",
                                        "tooltip": "MSCI ESG Quality Score",
                                        "value": "6.59 / 10",
                                        "rawValue": 6.5889,
                                        "label": "MSCI ESG Quality Score",
                                        "shortLabel": null,
                                        "name": "msciEsgQualityScore",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percentage of funds in a fund’s peer group with an ESG Score equal to, or lower than, the fund’s ESG Score.",
                                        "tooltip": "Peer Group Percentile Rank",
                                        "value": "44.54",
                                        "rawValue": 44.5376,
                                        "label": "Peer Group Percentile Rank",
                                        "shortLabel": null,
                                        "name": "msciEsgQualityScorePctlPeer",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percentage of funds, covered by MSCI, with a score lower than, or equal to, a fund’s ESG Quality Score.",
                                        "tooltip": "Global Percentile Rank",
                                        "value": "56.00",
                                        "rawValue": 55.9961,
                                        "label": "Global Percentile Rank",
                                        "shortLabel": null,
                                        "name": "msciEsgQualityScorePctlGlobal",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percentage of portfolio's market value exposed to companies flagged for one or more standard SRI exclusion factors (alcohol, civilian firearms, gambling, weapons, cluster bombs, landmines, nuclear power, GMOs, and tobacco).",
                                        "tooltip": "SRI Screening Criteria Exposure",
                                        "value": "9.65%",
                                        "rawValue": 9.65,
                                        "label": "SRI Screening Criteria Exposure",
                                        "shortLabel": null,
                                        "name": "msciSriExclusionCriteriaPct",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "The percentage of portfolio's market value exposed to companies addressing core environmental and social challenges.",
                                        "tooltip": "Exposure to Sustainable Impact Solutions",
                                        "value": "--",
                                        "rawValue": null,
                                        "label": "Exposure to Sustainable Impact Solutions",
                                        "shortLabel": null,
                                        "name": "msciSustainableImpactPct",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": "A measure of the fund's exposure to carbon intensive companies. The figure is the sum of security weight (normalized for corporate positions only) multiplied by the security Carbon Intensity.",
                                        "tooltip": "Weighted Average Carbon Intensity (t CO2e/$M Sales)",
                                        "value": "116.86",
                                        "rawValue": 116.86,
                                        "label": "Weighted Average Carbon Intensity (t CO2e/$M Sales)",
                                        "shortLabel": null,
                                        "name": "msciWeightedAvgCarbonInten",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundSharedHoldings",
                                "label": "VOO Benchmark Comparison Holdings",
                                "column": "full",
                                "type": "general_data",
                                "renderType": "general_data",
                                "order": 10,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "505",
                                        "rawValue": 505,
                                        "label": "VOO Number of Holdings",
                                        "shortLabel": null,
                                        "name": "numberOfHoldings",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": 278,
                                        "rawValue": 278,
                                        "label": "Benchmark Constituents",
                                        "shortLabel": null,
                                        "name": "numberOfConstituentsBenchmark",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "262",
                                        "rawValue": 262,
                                        "label": "Shared Holdings",
                                        "shortLabel": null,
                                        "name": "holdingsOverlap",
                                        "displayType": "text"
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "89.18%",
                                        "rawValue": 0.891829,
                                        "label": "Shared Holdings Weight",
                                        "shortLabel": null,
                                        "name": "holdingsOverlapPct",
                                        "displayType": "text"
                                    }
                                ]
                            },
                            {
                                "name": "fundBenchmarkComparison",
                                "label": "VOO Benchmark Comparison Summary",
                                "column": "full",
                                "gridColumns": "full",
                                "type": "comparison_table",
                                "renderType": "comparison_table",
                                "order": 11,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "505|278",
                                        "rawValue": "505|278",
                                        "label": "Number of Holdings",
                                        "shortLabel": null,
                                        "name": "numberOfHoldingsFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "$742.84B|$812.22B",
                                        "rawValue": "742843826000|812215819600",
                                        "label": "Weighted Average Market Cap",
                                        "shortLabel": null,
                                        "name": "weightedAvgMrktCapFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "23.90|24.72",
                                        "rawValue": "23.895711|24.724676",
                                        "label": "Price / Earnings Ratio",
                                        "shortLabel": null,
                                        "name": "peFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "4.46|4.89",
                                        "rawValue": "4.456846|4.88621",
                                        "label": "Price / Book Ratio",
                                        "shortLabel": null,
                                        "name": "pbFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "1.41%|1.36%",
                                        "rawValue": "0.014127|0.013581",
                                        "label": "Dividend Yield",
                                        "shortLabel": null,
                                        "name": "dividendYieldFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "Low|Medium",
                                        "rawValue": "Low|Medium",
                                        "label": "Concentration",
                                        "shortLabel": null,
                                        "name": "concentrationFundBenchmark",
                                        "displayType": null
                                    }
                                ]
                            },
                            {
                                "name": "fundMarketCapSize",
                                "label": "VOO Benchmark Comparison Market Cap Size",
                                "column": "full",
                                "gridColumns": "full",
                                "type": "comparison_table",
                                "renderType": "comparison_table",
                                "order": 12,
                                "footer": null,
                                "fields": [
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "98.19%|100.00%",
                                        "rawValue": "0.981866|1",
                                        "label": "Large (>12.9B)",
                                        "shortLabel": null,
                                        "name": "largeCapAllocationFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "1.56%|0.00%",
                                        "rawValue": "0.015635|0",
                                        "label": "Mid (>2.7B)",
                                        "shortLabel": null,
                                        "name": "midCapAllocationFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "0.00%|0.00%",
                                        "rawValue": "0|0",
                                        "label": "Small (>600M)",
                                        "shortLabel": null,
                                        "name": "smallCapAllocationFundBenchmark",
                                        "displayType": null
                                    },
                                    {
                                        "tooltipDescription": null,
                                        "tooltip": null,
                                        "value": "0.00%|0.00%",
                                        "rawValue": "0|0",
                                        "label": "Micro (<600M)",
                                        "shortLabel": null,
                                        "name": "microCapAllocationFundBenchmark",
                                        "displayType": null
                                    }
                                ]
                            },
                            {
                                "name": "topRegions",
                                "label": "VOO Regions",
                                "column": "full",
                                "gridColumns": "full",
                                "showTitle": true,
                                "type": "general_data_type2",
                                "renderType": "general_data_type2",
                                "order": 7,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "regionBreakdown",
                                        "data": [
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "North America",
                                                "weight": "100.00%",
                                                "benchmarkWeight": "100.00%",
                                                "rawWeight": 1
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "economicDevelopment",
                                "label": "VOO Economic Development",
                                "column": "full",
                                "gridColumns": "full",
                                "showTitle": true,
                                "type": "general_data_type2",
                                "renderType": "general_data_type2",
                                "order": 8,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "economicDevelopment",
                                        "data": [
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "Developed Countries",
                                                "weight": "100.00%",
                                                "benchmarkWeight": "100.00%",
                                                "rawWeight": 1
                                            },
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "Emerging Countries",
                                                "weight": "0.00%",
                                                "benchmarkWeight": "0.00%",
                                                "rawWeight": 0
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "sectorIndustryBreakdown",
                                "label": "VOO Sector/Industry Breakdown",
                                "column": "full",
                                "gridColumns": "1",
                                "showTitle": true,
                                "type": "general_data_type2",
                                "renderType": "general_data_type2",
                                "order": 2,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "sectorBreakdown",
                                        "data": [
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Technology Services",
                                                "weight": "20.09%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.20089252076290498
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Electronic Technology",
                                                "weight": "17.91%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.17912308505147378
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Finance",
                                                "weight": "12.44%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.1244008551773917
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Health Technology",
                                                "weight": "9.46%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.09459692903928167
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Retail Trade",
                                                "weight": "7.81%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.07810246071777315
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Consumer Non-Durables",
                                                "weight": "4.61%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.046087536519607085
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Producer Manufacturing",
                                                "weight": "3.66%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.03658774583309064
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Consumer Services",
                                                "weight": "3.45%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.03451999930167543
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Energy Minerals",
                                                "weight": "3.19%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.03188042504642332
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Commercial Services",
                                                "weight": "2.90%",
                                                "benchmarkWeight": "--",
                                                "rawWeight": 0.029018605018009395
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "topCountries",
                                "label": "VOO Countries",
                                "column": "full",
                                "gridColumns": "full",
                                "showTitle": true,
                                "type": "general_data_type2",
                                "renderType": "general_data_type2",
                                "order": 6,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "countryBreakdown",
                                        "data": [
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "United States",
                                                "weight": "100.00%",
                                                "benchmarkWeight": "99.89%",
                                                "rawWeight": 1
                                            },
                                            {
                                                "asOf": "2023-12-31T00:00:00.000Z",
                                                "name": "Canada",
                                                "weight": "--",
                                                "benchmarkWeight": "0.11%",
                                                "rawWeight": "--"
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "name": "top10Holdings",
                                "label": "VOO Top 10 Holdings",
                                "column": "full",
                                "gridColumns": "1",
                                "showTitle": false,
                                "type": "general_data_type1",
                                "renderType": "general_data_type1",
                                "order": 4,
                                "footer": null,
                                "fields": [
                                    {
                                        "ticker": "VOO",
                                        "type": "securities",
                                        "data": [
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Apple Inc.",
                                                "weight": "7.02%",
                                                "benchmarkWeight": "--",
                                                "shares": 356991624,
                                                "market_value": 68731597369,
                                                "symbol": "AAPL",
                                                "url": "/stock/AAPL",
                                                "rawWeight": 0.07017778847905504
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Microsoft Corporation",
                                                "weight": "6.97%",
                                                "benchmarkWeight": "--",
                                                "shares": 181486386,
                                                "market_value": 68246140591,
                                                "symbol": "MSFT",
                                                "url": "/stock/MSFT",
                                                "rawWeight": 0.06968211713739676
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Amazon.com, Inc.",
                                                "weight": "3.45%",
                                                "benchmarkWeight": "--",
                                                "shares": 222063192,
                                                "market_value": 33740281392,
                                                "symbol": "AMZN",
                                                "url": "/stock/AMZN",
                                                "rawWeight": 0.03445021535058239
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "NVIDIA Corporation",
                                                "weight": "3.05%",
                                                "benchmarkWeight": "--",
                                                "shares": 60314285,
                                                "market_value": 29868840218,
                                                "symbol": "NVDA",
                                                "url": "/stock/NVDA",
                                                "rawWeight": 0.03049731464380184
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Alphabet Inc. Class A",
                                                "weight": "2.06%",
                                                "benchmarkWeight": "--",
                                                "shares": 144509039,
                                                "market_value": 20186467658,
                                                "symbol": "GOOGL",
                                                "url": "/stock/GOOGL",
                                                "rawWeight": 0.020611213934645968
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Meta Platforms Inc Class A",
                                                "weight": "1.96%",
                                                "benchmarkWeight": "--",
                                                "shares": 54200007,
                                                "market_value": 19184634478,
                                                "symbol": "META",
                                                "url": "/stock/META",
                                                "rawWeight": 0.01958830104321579
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Alphabet Inc. Class C",
                                                "weight": "1.75%",
                                                "benchmarkWeight": "--",
                                                "shares": 121624966,
                                                "market_value": 17140606458,
                                                "symbol": "GOOG",
                                                "url": "/stock/GOOG",
                                                "rawWeight": 0.01750126434504762
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Tesla, Inc.",
                                                "weight": "1.71%",
                                                "benchmarkWeight": "--",
                                                "shares": 67534747,
                                                "market_value": 16781033935,
                                                "symbol": "TSLA",
                                                "url": "/stock/TSLA",
                                                "rawWeight": 0.017134126006526248
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "Berkshire Hathaway Inc. Class B",
                                                "weight": "1.62%",
                                                "benchmarkWeight": "--",
                                                "shares": 44428377,
                                                "market_value": 15845824941,
                                                "symbol": "BRK.B",
                                                "url": "/stock/BRK.B",
                                                "rawWeight": 0.01617923914987008
                                            },
                                            {
                                                "asOf": "2023-12-29T00:00:00.000Z",
                                                "name": "JPMorgan Chase & Co.",
                                                "weight": "1.23%",
                                                "benchmarkWeight": "--",
                                                "shares": 70594395,
                                                "market_value": 12008106590,
                                                "symbol": "JPM",
                                                "url": "/stock/JPM",
                                                "rawWeight": 0.012260770832703658
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        }
    }
}`

func TestCaltOutput(t *testing.T) {
	actual := crawler.ToString("calc", jsonText)
	require.Equal(t, "VOO\tVanguard\tA 95\t23.90\t0.03%\t1.34%\t12.70%\t14.55%\t33.53%\t8.28%\t505\t$424.79B\t$742.84B\t$2.17B\t30.81%\tS&P 500\tEquity: U.S.  -  Large Cap\tTechnology Services\t20.09%", actual)
}
