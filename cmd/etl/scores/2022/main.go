package main

import (
	"encoding/json"
	"github.com/nkralles/masters-web/internal/cfg"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"github.com/nkralles/masters-web/internal/persistence/pgdriver"
	"io/ioutil"
	"net/http"
)

const SCORES_URL = "https://www.masters.com/en_US/scores/feeds/2022/scores.json"

func main() {
	var driver persistence.MastersStorage
	var err error
	driver, err = pgdriver.Connect(cfg.Config.DatabaseURL)
	if err != nil {
		logger.Internal.Fatalf("storage driver 	init failed: %v", err)
	}
	persistence.SetDefaultDriver(driver)
	c := http.Client{}
	req, err := http.NewRequest(http.MethodGet, SCORES_URL, nil)
	if err != nil {
		logger.Internal.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")

	resp, err := c.Do(req)
	if err != nil {
		logger.Internal.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logger.Internal.Fatalf("%d:%s %s", resp.StatusCode, resp.Status, string(b))
	}
	defer resp.Body.Close()

	var data mastersScores

	err = json.Unmarshal(b, &data)
	if err != nil {
		logger.Internal.Fatal(err, string(b))
	}
	logger.Internal.Debugf("%+v\n", data)

}

type mastersScores struct {
	Data struct {
		CurrentRound string `json:"currentRound"`
		StatusRound  string `json:"statusRound"`
		Player       []struct {
			Id           string  `json:"id,omitempty"`
			DisplayName  string  `json:"display_name,omitempty"`
			DisplayName2 string  `json:"display_name2,omitempty"`
			FirstName    string  `json:"first_name,omitempty"`
			LastName     string  `json:"last_name,omitempty"`
			CountryName  string  `json:"countryName,omitempty"`
			CountryCode  string  `json:"countryCode,omitempty"`
			Teetime      *string `json:"teetime,omitempty"`
			Topar        *string `json:"topar,omitempty"`
			Round1       *Round  `json:"round1,omitempty"`
			Round2       *Round  `json:"round2,omitempty"`
			Round3       *Round  `json:"round3,omitempty"`
			Round4       *Round  `json:"round4,omitempty"`
		} `json:"player"`
	} `json:"data"`
}

type Round struct {
	RoundStatus *string   `json:"roundStatus,omitempty"`
	Scores      []*string `json:"scores,omitempty"`
	Total       *int      `json:"total,omitempty"`
}
