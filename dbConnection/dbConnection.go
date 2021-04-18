package dbConnection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	DbDriver   string
	DbHost     string
	DbPort     string
	DbUserName string
	DbPassword string
	DbName     string
}

func (cfg *config) dbSrc() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DbUserName,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName)
}

func loadConfig() *config {

	//Call the environment variables written in docker-compose.yaml
	cfg := config{
		DbDriver:   envOrDefault("DBDRIVER", "mysql"),
		DbHost:     envOrDefault("DBHOST", "go_db"),
		DbPort:     envOrDefault("DBPORT", "3306"),
		DbUserName: envOrDefault("DBUSERNAME", "root"),
		DbPassword: envOrDefault("DBPASSWORD", "passw0rd"),
		DbName:     envOrDefault("DBNAME", "techtraindb"),
	}

	return &cfg
}

func envOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

type DbConnection struct {
	pool *sql.DB
}

var sharedInstance *DbConnection = newDbConnection()

func newDbConnection() *DbConnection {

	cfg := loadConfig()

	pool, err := sql.Open(cfg.DbDriver, cfg.dbSrc())
	if err != nil {
		log.Fatalf("Error when executing sql.Open: %s", err)
		return nil
	}

	return &DbConnection{
		pool: pool,
	}
}

func GetInstance() *DbConnection {
	return sharedInstance
}

func (conn *DbConnection) GetConnection() *sql.DB {
	return conn.pool
}

func (conn *DbConnection) Close() {
	conn.pool.Close()
}
