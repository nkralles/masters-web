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
		golfer, err := driver.GetGolferByFullName(context.Background(), rec[0])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		r1 := 1000
		if len(rec[4]) > 0 {
			r1, err = strconv.Atoi(rec[4])
			if err != nil {
				logger.Internal.Fatal(err)
			}
		}
		err = driver.AddScore(context.Background(), golfer, 1, r1)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		r2 := 1000
		if len(rec[5]) > 0 {
			r2, err = strconv.Atoi(rec[5])
			if err != nil {
				logger.Internal.Fatal(err)
			}
		}
		err = driver.AddScore(context.Background(), golfer, 2, r2)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		r3 := 1000
		if len(rec[6]) > 0 {
			r3, err = strconv.Atoi(rec[6])
			if err != nil {
				logger.Internal.Fatal(err)
			}
		}
		err = driver.AddScore(context.Background(), golfer, 3, r3)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		r4 := 1000
		if len(rec[7]) > 0 {
			r4, err = strconv.Atoi(rec[7])
			if err != nil {
				logger.Internal.Fatal(err)
			}
		}
		err = driver.AddScore(context.Background(), golfer, 4, r4)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		logger.Internal.Debug(rec)
	}
}
