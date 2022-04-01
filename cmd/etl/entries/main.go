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
		var entry *persistence.Entry
		if entry, err = driver.GetEntryUser(context.Background(), rec[0]); err != nil {
			logger.Internal.Fatal(err)
		}
		score, err := strconv.Atoi(rec[9])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.SetEntryWinningScore(context.Background(), entry.Name, score)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		err = driver.DeleteAllEntriesForUser(context.Background(), entry)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer1, err := driver.GetGolferByFullName(context.Background(), rec[1])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		if cfg.Config.EnforceRank && !golfer1.Top12() {
			logger.Internal.Fatalf("entry %s has invalid golfer for top12_1 with %s %s with and rank of %d",
				entry.Name, golfer1.FirstName, golfer1.LastName, golfer1.Rank)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer1)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer2, err := driver.GetGolferByFullName(context.Background(), rec[2])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		if cfg.Config.EnforceRank && !golfer2.Top12() {
			logger.Internal.Fatalf("entry %s has invalid golfer for top12_2 with %s %s with and rank of %d",
				entry.Name, golfer2.FirstName, golfer2.LastName, golfer2.Rank)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer2)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer3, err := driver.GetGolferByFullName(context.Background(), rec[3])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		if cfg.Config.EnforceRank && !golfer3.Top12() {
			logger.Internal.Fatalf("entry %s has invalid golfer for top12_1 with %s %s with and rank of %d",
				entry.Name, golfer3.FirstName, golfer3.LastName, golfer3.Rank)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer3)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer4, err := driver.GetGolferByFullName(context.Background(), rec[4])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer4)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer5, err := driver.GetGolferByFullName(context.Background(), rec[5])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer5)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer6, err := driver.GetGolferByFullName(context.Background(), rec[6])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer6)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer7, err := driver.GetGolferByFullName(context.Background(), rec[7])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer7)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		golfer8, err := driver.GetGolferByFullName(context.Background(), rec[8])
		if err != nil {
			logger.Internal.Fatal(err)
		}
		err = driver.AddGolferEntryForUser(context.Background(), entry, golfer8)
		if err != nil {
			logger.Internal.Fatal(err)
		}

		//logger.Internal.Debugf("%+v\n%+v\n%+v\n%+v\n%+v\n%+v\n%+v\n%+v\n%+v",
		//	entry,
		//	golfer1,
		//	golfer2,
		//	golfer3,
		//	golfer4,
		//	golfer5,
		//	golfer6,
		//	golfer7,
		//	golfer8)
	}
}
