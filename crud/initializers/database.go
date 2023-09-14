package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	// postgres://zjjihiid:ejk67IfL6ih9ncVfMchFAN3ku6Ujd_Ft@berry.db.elephantsql.com/zjjihiid
	dsn := "user=zjjihiid password=ejk67IfL6ih9ncVfMchFAN3ku6Ujd_Ft dbname=zjjihiid host=berry.db.elephantsql.com port=5432 sslmode=disable TimeZone=UTC"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}