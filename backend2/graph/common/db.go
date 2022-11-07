package common

import (
		// "github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
    "fmt"
    "log"
    // "os"
    // "time"
		"github.com/jinzhu/gorm"
		_ "github.com/lib/pq"
)



func InitDb() *gorm.DB {

	const (
		Dialect = "postgres"
		DBUser = "postgres"
		DBPass = "postgres"
		DBProtocol = "tcp(127.0.0.1:5433)"
		DBName = "go-postgresql2"
	)
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)
	if err != nil {
		log.Println(err.Error())
}
return db

}