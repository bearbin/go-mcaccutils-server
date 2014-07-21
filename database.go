package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
)

func initDatabase() *gorp.DbMap {
	// Open the database.
	databasePath := filepath.Join(config.DataLocation, "mcaccutils-server.db")
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	// Create a gorp mapping.
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	// Add tables to the mapping. TODO: Set unique data that is unique.
	dbmap.AddTableWithName(Player{}, "players").SetKeys(false, "UUID")
	dbmap.AddTableWithName(NameRecord{}, "names").SetKeys(false, "UUID", "Username")
	dbmap.AddTableWithName(BanRecord{}, "bans").SetKeys(true, "ID")
	// Create the tables if they don't exist already.
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
	return dbmap
}
