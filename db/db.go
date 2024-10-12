package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tursodatabase/go-libsql"
)

func GetDbConnection() (*sql.DB, *libsql.Connector, string) {
	dbName := "local.db"
	primaryUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	dir, err := os.MkdirTemp(".db", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory for database", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, libsql.WithAuthToken(authToken), libsql.WithSyncInterval(time.Minute))
	if err != nil {
		fmt.Println("Error creating connector", err)
		os.Exit(1)
	}

	db := sql.OpenDB(connector)

	return db, connector, dir
}
