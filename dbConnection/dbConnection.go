package dbConnection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
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
		// If there is no environment variable,
		//    the value of the test environment will be inserted.
		DbDriver:   envOrDefault("DBDRIVER", "mysql"),
		DbHost:     envOrDefault("DBHOST", "127.0.0.1"),
		DbPort:     envOrDefault("DBPORT", "3306"),
		DbUserName: envOrDefault("DBUSERNAME", "root"),
		DbPassword: envOrDefault("DBPASSWORD", "passw0rd"),
		DbName:     envOrDefault("DBNAME", "techtraindb"),
	}
	fmt.Println(cfg, cfg.DbHost)

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
		logger.Errorf("Error when executing sql.Open: %s", err)
		return nil
	}
	logger.Printf("pool information : %v", pool)

	if err = pool.Ping(); err != nil {
		logger.Errorf("Error when executing pool.Ping: %s", err)
		panic("Error  pool.Ping")
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
