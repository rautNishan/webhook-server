package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var DB *sql.DB

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env")
	}

	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func (c *Config) ConnectionString() string {
	fmt.Println("This is config: ", c)
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.DbName)
}

func InitDB() error {
	config := LoadConfig()
	connStr := config.ConnectionString()

	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)

	}

	DB.SetMaxIdleConns(25)
	DB.SetMaxOpenConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	fmt.Println("Database Connected Successfully")

	return nil
}
