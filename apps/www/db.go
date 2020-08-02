package www

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/twocucao/zelda/conf"
	"log"
)

var db *gorm.DB

func init() {

	config := conf.GetConfig()

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.MySQL.User,
			config.MySQL.Password,
			config.MySQL.Host,
			config.MySQL.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.MySQL.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
