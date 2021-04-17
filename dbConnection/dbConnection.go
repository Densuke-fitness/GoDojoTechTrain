package dbConnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
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

func loadConfig() (*config, error) {

	var cfg config
	//Call the environment variables written in docker-compose.yaml
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	fmt.Println(cfg)

	return &cfg, nil
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

	pool, err := sql.Open(cfg.DbDriver, cfg.dbSrc())
	fmt.Println(pool)
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
