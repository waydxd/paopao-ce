// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Only support mysql database
package migrate

import (
	"fmt"

	"database/sql"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/waydxd/paopao-ce/cmd"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		Username     string `yaml:"Username"`
		Password     string `yaml:"Password"`
		Host         string `yaml:"Host"`
		DBName       string `yaml:"DBName"`
		Charset      string `yaml:"Charset"`
		ParseTime    bool   `yaml:"ParseTime"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"MySQL"`
}

func init() {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database data",
		Long:  "migrate database data when paopao-ce upgrade",
		Run:   migrateRun,
	}
	cmd.Register(migrateCmd)
}

func migrateRun(_cmd *cobra.Command, _args []string) {
	// Read the config file
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the config file
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create the DSN (Data Source Name)
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&multiStatements=true",
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.DBName,
		config.MySQL.Charset,
		config.MySQL.ParseTime,
	)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer db.Close()

	// Set the connection pool parameters
	db.SetMaxIdleConns(config.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(config.MySQL.MaxOpenConns)

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("Database connection established successfully!")

	// Run migrations
	runMigrations(dsn)
}

func runMigrations(dsn string) {
	// Open the database connection for migration
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("99 error: %v", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("105 error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/migration/mysql",
		"mysql", driver)
	if err != nil {
		log.Fatalf("112 error: %v", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply.")
		} else {
			log.Fatalf("116 error: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully!")
	}
}
