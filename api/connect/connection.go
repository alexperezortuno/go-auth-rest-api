package connect

import (
	"../migration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var Connection *gorm.DB
var engine_sql = os.Getenv("DB_DRIVER")
var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var username = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var data_base = os.Getenv("DB_NAME")

func Init() {
	Connection = ORM(StrConn())
	log.Println("Connection has been successfully")
}

func ORM(str_conn string) *gorm.DB {
	connection, err := gorm.Open(engine_sql, str_conn)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return connection
}

func CloseConn() {
	_ = Connection.Close()
	log.Println("Connection close")
}

func StrConn() string {
	if engine_sql == "postgres" {
		return "host=" + host + " port=" + port + " user=" + username + " dbname=" + data_base + " password=" + password + " sslmode=disable"
	}

	return ""
}

func Migrate() {
	migration.UserMigrate(Connection)
}
