package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){
	var err error
	//postgresql://postgres:3ZXA0D0INTcvSDUo@db.pfjdkelldjqmgxtqndps.supabase.co:5432/postgres
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
}