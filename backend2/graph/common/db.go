package common

import (
		"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/customTypes"

		"github.com/jinzhu/gorm"
		_ "github.com/lib/pq"
)

func InitDb() (*gorm.DB, error) {
    var err error
		db, err := gorm.Open("postgres", "host=postgres user=postgres port=5432 password=postgres dbname=postgres sslmode=disable")
    db.LogMode(true)

    if err != nil {
        return nil, err
    }
    db.AutoMigrate(&customTypes.Post{})
    return db, nil
}
