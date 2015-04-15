package ohmygorm

import (
	"testing"

	"github.com/jacobstr/confer"
)

type Example struct {
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
}
