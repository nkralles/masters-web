package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nkralles/masters-web/internal/cfg"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"github.com/nkralles/masters-web/internal/persistence/pgdriver"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		logger.Internal.Error(err)
	}

	//logger.Internal.Debugf("%+v\n", data)

	for {
		now := time.Now().In(loc)
		fmt.Printf("Hour:%d Min:%d\n", now.Hour(), now.Minute())
		if now.Hour() >= 10 && now.Hour() < 21 {
			logger.Internal.Info("pulling data feed at %s...", now.String())
			req, err := http.NewRequest(http.MethodGet, SCORES_URL, nil)
			if err != nil {
				logger.Internal.Error(err)
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")

			resp, err := c.Do(req)
			if err != nil {
				logger.Internal.Error(err)
			}
			b, err := ioutil.ReadAll(resp.Body)
			if resp.StatusCode != http.StatusOK {
				logger.Internal.Error("%d:%s %s", resp.StatusCode, resp.Status, string(b))
			}
			defer resp.Body.Close()

			var data mastersScores

			err = json.Unmarshal(b, &data)
			if err != nil {
				logger.Internal.Error(err)
			}

			for _, player := range data.Data.Player {
				if player.Status != "W" && player.Status != "C" {
					go func(player playerScore) {
						//if player.Round1.RoundStatus != nil && player.Topar != nil {
						//	logger.Internal.Debugf("%s %s r1 %s", player.FirstName, player.LastName, *player.Topar)
						//	golfer, err := driver.GetGolferByFullName(context.Background(), fmt.Sprintf("%s %s", player.FirstName, player.LastName))
						//	if err != nil {
						//		logger.Internal.Errorf("failed to find %s %s\n", player.FirstName, player.LastName)
						//		return
						//	}
						//	par, err := parsePar(*player.Topar)
						//	if err == nil {
						//		err = driver.AddScore(context.Background(), golfer, 1, par)
						//		if err != nil {
						//			logger.Internal.Error(err)
						//			return
						//		}
						//	} else {
						//		if !errors.Is(err, ErrNotValidScore) {
						//			logger.Internal.Errorf("unexpected score... %s", *player.Topar)
						//			return
						//		}
						//	}
						//}

						//if player.Round2.RoundStatus != nil && player.Topar != nil {
						//	logger.Internal.Debugf("%s %s r2 %s", player.FirstName, player.LastName, *player.Topar)
						//	golfer, err := driver.GetGolferByFullName(context.Background(), fmt.Sprintf("%s %s", player.FirstName, player.LastName))
						//	if err != nil {
						//		logger.Internal.Errorf("failed to find %s %s\n", player.FirstName, player.LastName)
						//		return
						//	}
						//	par, err := parsePar(*player.Topar)
						//	if err == nil {
						//		err = driver.AddScore(context.Background(), golfer, 2, par)
						//		if err != nil {
						//			logger.Internal.Error(err)
						//			return
						//		}
						//	} else {
						//		if !errors.Is(err, ErrNotValidScore) {
						//			logger.Internal.Errorf("unexpected score... %s", *player.Topar)
						//			return
						//		}
						//	}
						//}

						//if player.Round3.RoundStatus != nil && player.Topar != nil {
						//	logger.Internal.Debugf("%s %s r3 %s", player.FirstName, player.LastName, *player.Topar)
						//	golfer, err := driver.GetGolferByFullName(context.Background(), fmt.Sprintf("%s %s", player.FirstName, player.LastName))
						//	if err != nil {
						//		logger.Internal.Errorf("failed to find %s %s\n", player.FirstName, player.LastName)
						//		return
						//	}
						//	par, err := parsePar(*player.Topar)
						//	if err == nil {
						//		err = driver.AddScore(context.Background(), golfer, 3, par)
						//		if err != nil {
						//			logger.Internal.Error(err)
						//			return
						//		}
						//	} else {
						//		if !errors.Is(err, ErrNotValidScore) {
						//			logger.Internal.Errorf("unexpected score... %s", *player.Topar)
						//			return
						//		}
						//	}
						//}
						if player.Round4.RoundStatus != nil && player.Topar != nil {
							logger.Internal.Debugf("%s %s r4 %s", player.FirstName, player.LastName, *player.Topar)
							golfer, err := driver.GetGolferByFullName(context.Background(), fmt.Sprintf("%s %s", player.FirstName, player.LastName))
							if err != nil {
								logger.Internal.Errorf("failed to find %s %s\n", player.FirstName, player.LastName)
								return
							}
							par, err := parsePar(*player.Topar)
							if err == nil {
								err = driver.AddScore(context.Background(), golfer, 4, par)
								if err != nil {
									logger.Internal.Error(err)
									return
								}
							} else {
								if !errors.Is(err, ErrNotValidScore) {
									logger.Internal.Errorf("unexpected score... %s", *player.Topar)
									return
								}
							}
						}
					}(player)
				}
			}
		}

		time.Sleep(2 * time.Minute)
	}

}

type mastersScores struct {
	Data struct {
		CurrentRound string        `json:"currentRound"`
		StatusRound  string        `json:"statusRound"`
		Player       []playerScore `json:"player"`
	} `json:"data"`
}

type playerScore struct {
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
	Active       bool    `json:"active,omitempty"`
	Status       string  `json:"status"`
}

var ErrNotValidScore = fmt.Errorf("empty score not valid")

func parsePar(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrNotValidScore
	}
	if s == "E" || s == "e" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

type Round struct {
	RoundStatus *string `json:"roundStatus,omitempty"`
	Scores      []*int  `json:"scores,omitempty"`
	Total       *int    `json:"total,omitempty"`
}
