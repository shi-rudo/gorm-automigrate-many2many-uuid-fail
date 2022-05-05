package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", "localhost", "gorm", "gorm", "gorm", "9920", "disable", "Etc/UTC")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Migrate the schema
	err = db.AutoMigrate(
		&User{},
		&Language{},
	)
	if err != nil {
		panic("failed to migrate")
	}
}
