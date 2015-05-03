package ohmygorm

import (
	"errors"
	"fmt"

	"github.com/jacobstr/confer"
	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"

	// Load database drivers for Gorm
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// ConnectionsService provides functions for connecting to database
type ConnectionsService struct {
	Config            *confer.Config `inject:""`
	CurrentConnection *gorm.DB
}

// Make creates a new Gorm connection to the configured database
func (c *ConnectionsService) Make() (*gorm.DB, error) {
	// Check if we already have a reference to a connection
	if c.CurrentConnection != nil {
		return c.CurrentConnection, nil
	}

	var db gorm.DB
	var err error

	// Setup Gorm with the specified driver
	switch c.Config.GetString("database.driver") {
	case "sqlite":
		db, err = gorm.Open("sqlite3", c.Config.GetString("database.file"))
		break
	case "mysql":
		pretty.Println(c.Config.GetString("database.host"))
		uri := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			c.Config.GetString("database.username"),
			c.Config.GetString("database.password"),
			c.Config.GetString("database.host"),
			c.Config.GetInt("database.port"),
			c.Config.GetString("database.name"),
		)

		db, err = gorm.Open("mysql", uri)
	default:
		err = errors.New("Invalid or no driver provided")
	}

	// Check if we got nay errors while setting up the connection
	if err != nil {
		return nil, err
	}

	// Keep a reference to the connection so that we can access it later
	c.CurrentConnection = &db

	return &db, nil
}
