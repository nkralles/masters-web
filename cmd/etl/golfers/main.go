package main

import (
	"context"
	"encoding/csv"
	"github.com/nkralles/masters-web/internal/cfg"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"github.com/nkralles/masters-web/internal/persistence/pgdriver"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var driver persistence.MastersStorage
	var err error
	driver, err = pgdriver.Connect(cfg.Config.DatabaseURL)
	if err != nil {
		logger.Internal.Fatalf("storage driver 	init failed: %v", err)
	}
	persistence.SetDefaultDriver(driver)

	incsv := os.Args[1]

	if len(incsv) == 0 {
		logger.Internal.Fatal("missing input file")
	}

	if path := filepath.Ext(incsv); path != ".csv" {
		logger.Internal.Fatal("expected input file to be a csv file")
	}

	csvFile, err := os.Open(incsv)
	if err != nil {
		logger.Internal.Fatal(err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			logger.Internal.Error(err)
		}
	}(csvFile)
	rdr := csv.NewReader(csvFile)
	headers, err := rdr.Read()
	if err != nil {
		logger.Internal.Fatal(err)
	}
	logger.Internal.Debugf("csv headers %v", headers)
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Internal.Fatal(err)
		}
		pid, err := strconv.Atoi(rec[0])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		rank, err := strconv.Atoi(rec[1])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		golfer := persistence.Golfer{
			PlayerId:    pid,
			Rank:        rank,
			FirstName:   strings.TrimSpace(rec[6]),
			LastName:    strings.TrimSpace(rec[7]),
			CountryCode: rec[4],
		}
		err = driver.AddGolfer(context.Background(), &golfer)
		if err != nil {
			logger.Internal.Fatal(err)
		}
	}
}
