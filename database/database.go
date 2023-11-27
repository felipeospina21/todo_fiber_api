package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Instance *gorm.DB
}

var DB DbInstance

func Connect() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		fmt.Println(".env file not found")
	}
	psw := myEnv["DB_PASSWORD"]

	dsn := fmt.Sprintf("user=postgres password=%v host=db.lufltxfmegywzhucpurw.supabase.co port=5432 dbname=postgres", psw)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = DbInstance{Instance: db}
}
