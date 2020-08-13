package core

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/go-pg/pg"
	_ "github.com/lib/pq" // pq
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
)

var Conn *pg.DB

type DbConf struct {
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	HasMigrations bool
	MigrationPath string
}

func ConnectDb(c *DbConf) *pg.DB {
	opts := &pg.Options{
		Addr:     fmt.Sprintf("%v:%v", c.Host, c.Port),
		Database: c.Name,
		User:     c.User,
		Password: c.Password,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if c.HasMigrations {
		runMigrations(c)
	}

	Conn = pg.Connect(opts)
	log.Info("Database connection established ...")

	return Conn
}

func runMigrations(c *DbConf) {
	log.Info("Preparing to run migrations")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Opening connection for migrations failed with message: ", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: c.MigrationPath,
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Error applying migrations, message: ", err)
	}

	log.Info("Migrations successfully applied, count=", n)
}
