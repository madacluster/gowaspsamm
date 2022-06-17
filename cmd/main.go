package main

import (
	"encoding/json"
	"fmt"

	"github.com/madacluster/gowaspsamm/pkg/owaspsamm"
	"github.com/pterm/pterm"
)

const owasp_path = "./owasp-core/model"

func main() {
	db, err := owaspsamm.NewDB(owasp_path)
	if err != nil {
		fmt.Println(err)
	}

	for _, q := range db.Questions {
		db.Responses[q.ID] = &owaspsamm.Answer{
			Question: q.ID,
			Value:    3,
		}
	}
	radar := db.GetRadarChartPractices()
	bolB, err := json.Marshal(radar)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bolB))
	positiveBars := pterm.Bars{}
	for key, p := range radar {
		positiveBars = append(positiveBars, pterm.Bar{
			Label: key,
			Value: p,
		})
	}

	pterm.Info.Println("Chart example with positive only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(positiveBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(positiveBars).Render()
}
