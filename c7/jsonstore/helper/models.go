package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "127.0.0.1"
	port = 5432
	user = "postgres"
	password = "123456"
	dbname = "mydb"
)

type Shipment struct {
	gorm.Model
	Packages []Package
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

// GORM creates tables with plural names.
// Use this to suppress it
func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}


func InitDB() (*gorm.DB, error){
	var err error
	//dsn := "postgres://postgres:123456@localhost/mydb?sslmode=disable"
	dsn := "host=localhost user=postgres password=123456 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Shipment{}, &Package{})
	if err != nil {
		return nil, err
	}
	return db, nil
}