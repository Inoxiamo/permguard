// Copyright 2024 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"database/sql"
	"embed"
	"flag"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	azconfigs "github.com/permguard/permguard/pkg/configs"
	azidb "github.com/permguard/permguard/plugin/storage/sqlite/internal/extensions/db"
)

const (
	flagPath = "filepath"
	flagUp   = "up"
	flagDown = "down"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// SQLiteStorageProvisioner is the storage provisioner for SQLite.
type SQLiteStorageProvisioner struct {
	debug    bool
	logLevel string
	logger   *zap.Logger
	filePath string
	up       bool
	down     bool
	config   *azidb.SQLiteConnectionConfig
}

// NewSQLiteStorageProvisioner creates a new SQLiteStorageProvisioner.
func NewSQLiteStorageProvisioner() (*SQLiteStorageProvisioner, error) {
	config, err := azidb.NewSQLiteConnectionConfig()
	if err != nil {
		return nil, err
	}
	return &SQLiteStorageProvisioner{
		config: config,
	}, nil
}

// AddFlags adds flags.
func (p *SQLiteStorageProvisioner) AddFlags(flagSet *flag.FlagSet) error {
	err := azconfigs.AddFlagsForCommon(flagSet)
	if err != nil {
		return err
	}
	flagSet.String(flagPath, ".", "file path to the database")
	flagSet.Bool(flagUp, false, "provision the database")
	flagSet.Bool(flagDown, false, "deprovision the database")
	err = p.config.AddFlags(flagSet)
	if err != nil {
		return err
	}
	return nil
}

// InitFromViper initializes the configuration from viper.
func (p *SQLiteStorageProvisioner) InitFromViper(v *viper.Viper) error {
	debug, logLevel, err := azconfigs.InitFromViperForCommon(v)
	if err != nil {
		return err
	}
	p.debug = debug
	p.logLevel = logLevel
	p.filePath = v.GetString(flagPath)
	p.up = v.GetBool(flagUp)
	p.down = v.GetBool(flagDown)
	err = p.config.InitFromViper(v)
	if err != nil {
		return err
	}
	p.logger, err = azconfigs.NewLogger(p.debug, p.logLevel)
	if err != nil {
		return err
	}
	return nil
}

// setup sets up the database.
func (p *SQLiteStorageProvisioner) setup() (*sql.DB, error) {
	filePath := p.filePath
	dbName := p.config.GetDBName()
    if !strings.HasSuffix(dbName, ".db") {
        dbName += ".db"
    }
    dbPath := filepath.Join(filePath, dbName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	goose.SetLogger(azidb.NewGooseLogger(p.logger))
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}
	return db, nil
}

// Up provisions the database.
func (p *SQLiteStorageProvisioner) Up() error {
	if !p.up {
		p.logger.Info("Database provisioning skipped")
		return nil
	}
	p.logger.Debug("Provisioning database")
	db, err := p.setup()
	if err != nil {
		p.logger.Error("Database provisioning failed", zap.Error(err))
		return err
	}
	defer db.Close()
	if err := goose.Up(db, "migrations"); err != nil {
		p.logger.Error("Database provisioning failed", zap.Error(err))
		return err
	}
	p.logger.Info("Database provisioned")
	return nil
}

// Down deprovisions the database.
func (p *SQLiteStorageProvisioner) Down() error {
	if !p.down {
		p.logger.Info("Database deprovisioning skipped")
		return nil
	}
	p.logger.Debug("Deprovisioning database")
	db, err := p.setup()
	if err != nil {
		p.logger.Error("Database deprovisioning failed", zap.Error(err))
		return err
	}
	defer db.Close()
	for err == nil {
		err = goose.Down(db, "migrations")
	}
	p.logger.Info("Database deprovisioned")
	return nil
}