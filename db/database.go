package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Ateto/User-Login-Service/config"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	DB *sql.DB
}

func NewDB(configLocation, sqlLocation string) (*MysqlDatabase, error) {
	config.SetupEnv(configLocation)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Host,
		config.AppConfig.Database.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	fmt.Println("Connected to MySQL!")

	createTableQuery, err := os.ReadFile(sqlLocation)
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}

	_, err = db.Exec(string(createTableQuery))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
	return &MysqlDatabase{DB: db}, nil
}
