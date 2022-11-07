package common

import (
		// "github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
    "fmt"
    "log"
    "os"
    "time"
		"github.com/jinzhu/gorm"
		_ "github.com/lib/pq"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
)

// func InitDb() *gorm.DB {
//     // var err error
// 		db := gorm.Open("postgres", "host=postgres user=postgres port=5432 password=postgres dbname=postgres sslmode=disable")
//     db.LogMode(true)

//     // if err != nil {
//     //     return nil, err
//     // }
//     db.AutoMigrate(&customTypes.Post{})
//     return db
// }

func InitDb() *gorm.DB {

	const (
    Dialect = "postgres"
		DBUser = "postgres"
		DBPass = "postgres"
		DBProtocol = "tcp(127.0.0.1:5432)"
		DBName = "go-postgresql1"
	)
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)
	if err != nil {
		log.Println(err.Error())
}
return db
}

func initLog() logger.Interface {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: true,
		TablePrefix:   "",
	}
}