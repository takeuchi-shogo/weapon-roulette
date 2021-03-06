package infrastrcture

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Host       string
	UserName   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		UserName: c.DB.Production.UserName,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func newDB(d *DB) *DB {

	db, err := gorm.Open("mysql", d.UserName+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	d.Connection = db
	return d
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
