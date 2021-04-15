package dbConnection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	DBDriver   string `json:"DBDRIVER"`
	DBHost     string `json:"DBHOST"`
	DBPort     string `json:"DBPORT"`
	DBUserName string `json:"DBUSERNAME"`
	DBPassword string `json:"DBPASSWORD"`
	DBName     string `json:"DBNAME"`
}

func (cfg *config) dbSrc() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUserName,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)
}

func loadConfig() (*config, error) {
	f, err := os.Open("./dbConnection/config.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return nil, err
	}

	defer f.Close()

	var cfg config

	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

type DbConnection struct {
	pool *sql.DB
}

var sharedInstance *DbConnection = newDbConnection()

func newDbConnection() *DbConnection {

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Error when executing loadConfig: %s", err)
		return nil
	}

	pool, err := sql.Open(cfg.DBDriver, cfg.dbSrc())
	if err != nil {
		//TODO
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
