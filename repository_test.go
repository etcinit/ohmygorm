package ohmygorm

import (
	"testing"

	"github.com/facebookgo/inject"
	"github.com/jacobstr/confer"
	"github.com/stretchr/testify/assert"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func Test_Exists(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	var migrator MigrationsService
	var connections ConnectionsService
	var repository RepositoryService

	var g inject.Graph
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &migrator},
		&inject.Object{Value: &connections},
		&inject.Object{Value: &repository},
	)
	g.Populate()

	migrator.Run([]interface{}{&User{}})

	db, _ := connections.Make()

	bobby := User{FirstName: "Bobby", LastName: "Tables"}

	assert.False(t, repository.Exists(&User{}, 1))

	db.Create(&bobby)

	assert.True(t, repository.Exists(&User{}, 1))
}

func Test_Find(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	var migrator MigrationsService
	var connections ConnectionsService
	var repository RepositoryService

	var g inject.Graph
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &migrator},
		&inject.Object{Value: &connections},
		&inject.Object{Value: &repository},
	)
	g.Populate()

	migrator.Run([]interface{}{&User{}})

	db, _ := connections.Make()

	bobby := User{FirstName: "Bobby", LastName: "Tables"}

	var found User
	repository.Find(&found, 1)

	assert.Equal(t, 0, found.ID)

	db.Create(&bobby)

	repository.Find(&found, 1)
	assert.NotEqual(t, 0, found.ID)
}

func Test_FirstOrFail(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	var migrator MigrationsService
	var connections ConnectionsService
	var repository RepositoryService

	var g inject.Graph
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &migrator},
		&inject.Object{Value: &connections},
		&inject.Object{Value: &repository},
	)
	g.Populate()

	migrator.Run([]interface{}{&User{}})

	db, _ := connections.Make()

	bobby := User{FirstName: "Bobby", LastName: "Tables"}

	var found User
	err := repository.FirstOrFail(&found, db.Where("id = ?", 1))

	assert.NotNil(t, err)

	db.Create(&bobby)

	err = repository.FirstOrFail(&found, db.Where("id = ?", 1))
	assert.Nil(t, err)
}
