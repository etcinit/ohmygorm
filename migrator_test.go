package ohmygorm

import (
	"testing"

	"github.com/jacobstr/confer"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	ID        int
	FirstName string
	LastName  string
}

func Test_Run(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	connections := ConnectionsService{
		Config: config,
	}

	migrator := MigrationsService{
		Connections: &connections,
	}

	migrator.Run([]interface{}{&Example{}})

	db, _ := connections.Make()

	bobby := Example{FirstName: "Bobby", LastName: "Tables"}

	assert.Equal(t, 0, bobby.ID)

	db.Create(&bobby)

	assert.NotEqual(t, 0, bobby.ID)
}
