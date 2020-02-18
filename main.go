package main

import (
	cmd "github.com/astroband/astrologer/commands"
	cfg "github.com/astroband/astrologer/config"
	"github.com/astroband/astrologer/db"
	"github.com/astroband/astrologer/es"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"time"
)

func main() {
	kingpin.Version(cfg.Version)
	commandName := kingpin.Parse()

	esClient := es.Connect((*cfg.EsURL).String())

	var command cmd.Command

	switch commandName {
	case "stats":
		dbClient := db.Connect(*cfg.DatabaseURL)
		command = &cmd.StatsCommand{ES: esClient, DB: dbClient}
	case "create-index":
		config := cmd.CreateIndexCommandConfig{Force: *cfg.ForceRecreateIndexes}
		command = &cmd.CreateIndexCommand{ES: esClient, Config: config}
	case "export":
		dbClient := db.Connect(*cfg.DatabaseURL)
		config := cmd.ExportCommandConfig{
			Start:      cfg.Start,
			Count:      *cfg.Count,
			DryRun:     *cfg.DryRun,
			RetryCount: *cfg.Retries,
			BatchSize:  *cfg.BatchSize,
		}
		command = &cmd.ExportCommand{ES: esClient, DB: dbClient, Config: config}
	case "ingest":
		dbClient := db.Connect(*cfg.DatabaseURL)
		command = &cmd.IngestCommand{ES: esClient, DB: dbClient}
	case "es-stats":
		command = &cmd.EsStatsCommand{ES: esClient}
	case "fill-gaps":
		dbClient := db.Connect(*cfg.DatabaseUrl)
		config := &cmd.FillGapsCommandConfig{
			DryRun: *cfg.DryRun,
			Start:  cfg.FillGapsFrom,
			Count:  cfg.FillGapsCount,
		}
		command = &cmd.FillGapsCommand{ES: esClient, DB: dbClient, Config: config}
	}

	start := time.Now()
	command.Execute()
	elapsed := time.Since(start)

	log.Printf("Command took %s", elapsed)
}
