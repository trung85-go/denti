package storage

import (
	"errors"
	"fmt"

	"denti/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb(c *config.Config) (*gorm.DB, error) {
	if c.DB.Use == "postgres" {
		return newPostgres(c)
	} else if c.DB.Use == "mysql" {
		return newMysql(c)
	}
	return nil, errors.New("Not supported db")
}

func newPostgres(c *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Postgres.Host,
		c.DB.Postgres.Port,
		c.DB.Postgres.UserName,
		c.DB.Postgres.Password,
		c.DB.Postgres.Database)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newMysql(c *config.Config) (*gorm.DB, error) {
	msqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.Mysql.UserName,
		c.DB.Mysql.Password,
		c.DB.Mysql.Host,
		c.DB.Mysql.Port,
		c.DB.Mysql.Database)

	db, err := gorm.Open("mysql", msqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
